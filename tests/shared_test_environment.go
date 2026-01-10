package tests

import (
	"crypto/sha256"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

// SharedPiholeEnvironment manages a shared Pi-hole instance for non-destructive tests
type SharedPiholeEnvironment struct {
	TerraformOptions *terraform.Options
	BaseURL          string
	Password         string
	DNSPort          int
	WebPort          int
	ContainerName    string
	NetworkName      string
	Initialized      bool
	mu               sync.Mutex
}

var (
	// Global shared environment instance
	sharedEnv     *SharedPiholeEnvironment
	sharedEnvOnce sync.Once
)

// GetSharedPiholeEnvironment returns the singleton shared environment
func GetSharedPiholeEnvironment() *SharedPiholeEnvironment {
	sharedEnvOnce.Do(func() {
		// Generate unique identifier for this test session
		sessionID := fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("shared-env-%d", time.Now().Unix()))))[:8]
		
		sharedEnv = &SharedPiholeEnvironment{
			BaseURL:       fmt.Sprintf("http://localhost:%d", 30080+len(sessionID)%100),
			Password:      fmt.Sprintf("shared-test-%s", sessionID),
			DNSPort:       30353 + len(sessionID)%100,
			WebPort:       30080 + len(sessionID)%100,
			ContainerName: fmt.Sprintf("pihole-shared-test-%s", sessionID),
			NetworkName:   fmt.Sprintf("pihole-shared-net-%s", sessionID),
			Initialized:   false,
		}

		sharedEnv.TerraformOptions = terraform.WithDefaultRetryableErrors(nil, &terraform.Options{
			TerraformDir: "../terraform/modules/pihole",
			Vars: map[string]interface{}{
				"container_name":     sharedEnv.ContainerName,
				"network_name":       sharedEnv.NetworkName,
				"subnet":             fmt.Sprintf("172.%d.0.0/16", 100+len(sessionID)%150), // Wide subnet spacing for isolation
				"dns_port":           sharedEnv.DNSPort,
				"web_port":           sharedEnv.WebPort,
				"timezone":           "America/New_York",
				"web_password":       sharedEnv.Password,
				"dnsmasq_listening":  "all",
				"use_host_network":   false,
			},
		})
	})
	return sharedEnv
}

// Setup initializes the shared Pi-hole environment if not already done
func (env *SharedPiholeEnvironment) Setup(t *testing.T) error {
	env.mu.Lock()
	defer env.mu.Unlock()

	if env.Initialized {
		return nil
	}

	t.Log("Setting up shared Pi-hole test environment...")
	
	// Check if we should skip shared environment setup
	if os.Getenv("SKIP_SHARED_SETUP") == "true" {
		env.Initialized = true
		return nil
	}

	// Initialize and apply terraform
	terraform.InitAndApply(t, env.TerraformOptions)
	
	// Wait for Pi-hole to be ready
	t.Log("Waiting for shared Pi-hole to start...")
	time.Sleep(60 * time.Second) // Give shared environment more time
	
	env.Initialized = true
	t.Log("Shared Pi-hole environment ready")
	return nil
}

// Cleanup destroys the shared environment (called once at end of test suite)
func (env *SharedPiholeEnvironment) Cleanup(t *testing.T) {
	env.mu.Lock()
	defer env.mu.Unlock()

	if !env.Initialized {
		return
	}

	// Check if we should skip cleanup (useful for debugging)
	if os.Getenv("SKIP_SHARED_CLEANUP") == "true" {
		t.Log("Skipping shared environment cleanup")
		return
	}

	t.Log("Cleaning up shared Pi-hole environment...")
	terraform.Destroy(t, env.TerraformOptions)
	env.Initialized = false
}

// GetSession creates an authenticated session to the shared Pi-hole
func (env *SharedPiholeEnvironment) GetSession() (*PiholeSession, error) {
	if !env.Initialized {
		return nil, fmt.Errorf("shared environment not initialized")
	}
	
	return NewPiholeSession(env.BaseURL, env.Password)
}

// IsHealthy performs a basic health check on the shared environment
func (env *SharedPiholeEnvironment) IsHealthy(t *testing.T) bool {
	if !env.Initialized {
		return false
	}
	
	session, err := env.GetSession()
	if err != nil {
		t.Logf("Health check failed - cannot create session: %v", err)
		return false
	}
	
	err = session.TestAPIAccess()
	if err != nil {
		t.Logf("Health check failed - API access error: %v", err)
		return false
	}
	
	return true
}

// SharedTestConfig provides configuration for tests using shared environment
type SharedTestConfig struct {
	UseSharedEnvironment bool
	RequiresDestruction  bool // If true, test cannot use shared environment
	TestCategory         string // "api", "dns", "config", "destructive"
}

// CanUseSharedEnvironment determines if a test can use the shared environment
func (config SharedTestConfig) CanUseSharedEnvironment() bool {
	return config.UseSharedEnvironment && !config.RequiresDestruction
}

// GetTestEnvironment returns appropriate environment (shared or dedicated) for a test
func GetTestEnvironment(t *testing.T, config SharedTestConfig) (*terraform.Options, string, string, error) {
	if config.CanUseSharedEnvironment() {
		// Use shared environment
		env := GetSharedPiholeEnvironment()
		err := env.Setup(t)
		if err != nil {
			return nil, "", "", fmt.Errorf("failed to setup shared environment: %v", err)
		}
		
		if !env.IsHealthy(t) {
			return nil, "", "", fmt.Errorf("shared environment is not healthy")
		}
		
		return env.TerraformOptions, env.BaseURL, env.Password, nil
	}
	
	// Create dedicated environment
	return createDedicatedEnvironment(t, config)
}

// createDedicatedEnvironment creates a dedicated test environment
func createDedicatedEnvironment(t *testing.T, config SharedTestConfig) (*terraform.Options, string, string, error) {
	// Generate unique identifier based on test name + timestamp for true uniqueness
	testID := fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%s-%d", t.Name(), time.Now().UnixNano()))))[:8]
	basePort := 31000 + (len(testID) % 1000) // Use hash for port uniqueness
	
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/modules/pihole",
		Vars: map[string]interface{}{
			"container_name":     fmt.Sprintf("pihole-test-%s", testID),
			"network_name":       fmt.Sprintf("pihole-net-%s", testID),
			"subnet":             fmt.Sprintf("172.%d.0.0/16", 101+(len(testID)%150)), // Wide subnet spacing for isolation
			"dns_port":           basePort,
			"web_port":           basePort + 1,
			"timezone":           "America/New_York", 
			"web_password":       fmt.Sprintf("test-password-%s", testID),
			"dnsmasq_listening":  "all",
			"use_host_network":   false,
		},
	})
	
	baseURL := fmt.Sprintf("http://localhost:%d", basePort+1)
	password := fmt.Sprintf("test-password-%s", testID)
	
	return terraformOptions, baseURL, password, nil
}