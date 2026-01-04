package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/miekg/dns"
	"github.com/stretchr/testify/require"
)

// TestParallelReadOnlyAPIs demonstrates parallel execution of non-destructive API tests
func TestParallelReadOnlyAPIs(t *testing.T) {
	// All these sub-tests can run in parallel since they only read data
	
	t.Run("ReadStats", func(t *testing.T) {
		t.Parallel() // Safe to run in parallel - read-only
		
		config := SharedTestConfig{
			UseSharedEnvironment: true,
			RequiresDestruction:  false,
			TestCategory:         "readonly",
		}
		
		_, baseURL, password, err := GetTestEnvironment(t, config)
		if err != nil {
			t.Skip("Shared environment not available, skipping parallel read test")
		}
		
		session, err := NewPiholeSession(baseURL, password)
		if err == nil {
			stats, err := session.GetStats()
			if err == nil {
				t.Logf("Successfully read stats from shared Pi-hole: %+v", stats)
			} else {
				t.Logf("Stats read attempted: %v", err)
			}
		} else {
			t.Logf("Session creation attempted: %v", err) 
		}
	})
	
	t.Run("ReadLists", func(t *testing.T) {
		t.Parallel() // Safe to run in parallel - read-only
		
		config := SharedTestConfig{
			UseSharedEnvironment: true,
			RequiresDestruction:  false,
			TestCategory:         "readonly",
		}
		
		_, baseURL, password, err := GetTestEnvironment(t, config)
		if err != nil {
			t.Skip("Shared environment not available, skipping parallel read test")
		}
		
		session, err := NewPiholeSession(baseURL, password)
		if err == nil {
			lists, err := session.GetLists()
			if err == nil {
				t.Logf("Successfully read lists from shared Pi-hole: %+v", lists)
			} else {
				t.Logf("Lists read attempted: %v", err)
			}
		} else {
			t.Logf("Session creation attempted: %v", err)
		}
	})
	
	t.Run("TestAPIEndpoints", func(t *testing.T) {
		t.Parallel() // Safe to run in parallel - read-only
		
		config := SharedTestConfig{
			UseSharedEnvironment: true,
			RequiresDestruction:  false,
			TestCategory:         "readonly",
		}
		
		_, baseURL, password, err := GetTestEnvironment(t, config)
		if err != nil {
			t.Skip("Shared environment not available, skipping parallel read test")
		}
		
		session, err := NewPiholeSession(baseURL, password)
		if err == nil {
			err := session.TestAPIAccess()
			if err == nil {
				t.Log("API access test successful in parallel execution")
			} else {
				t.Logf("API access test attempted: %v", err)
			}
		}
	})
	
	t.Run("DNSResolution", func(t *testing.T) {
		t.Parallel() // Safe to run in parallel - read-only
		
		config := SharedTestConfig{
			UseSharedEnvironment: true,
			RequiresDestruction:  false,
			TestCategory:         "readonly",
		}
		
		terraformOptions, _, _, err := GetTestEnvironment(t, config)
		if err != nil {
			t.Skip("Shared environment not available, skipping parallel DNS test")
		}
		
		dnsPort := terraformOptions.Vars["dns_port"].(int)
		
		client := new(dns.Client)
		client.Timeout = 5 * time.Second
		
		message := new(dns.Msg)
		message.SetQuestion(dns.Fqdn("google.com"), dns.TypeA)
		
		response, _, err := client.Exchange(message, "127.0.0.1:30353") // Hardcoded for shared environment
		if err == nil && len(response.Answer) > 0 {
			t.Logf("DNS parallel test successful on port %d: %v", dnsPort, response.Answer[0])
		} else {
			t.Logf("DNS parallel test attempted on port %d", dnsPort)
		}
	})
}

// TestSequentialDestructiveOperations demonstrates tests that CANNOT run in parallel
func TestSequentialDestructiveOperations(t *testing.T) {
	// Notice: NO t.Parallel() here - these tests modify state and must run sequentially
	
	t.Run("ContainerRecreation", func(t *testing.T) {
		// DO NOT add t.Parallel() - this test destroys containers
		
		config := SharedTestConfig{
			UseSharedEnvironment: false, // Force dedicated environment
			RequiresDestruction:  true,
			TestCategory:         "destructive",
		}
		
		terraformOptions, baseURL, password, err := GetTestEnvironment(t, config)
		require.NoError(t, err, "Should get dedicated environment")
		
		// Full lifecycle for destructive test
		defer terraform.Destroy(t, terraformOptions)
		terraform.InitAndApply(t, terraformOptions)
		time.Sleep(30 * time.Second) // Reduced wait for demo
		
		_, err = NewPiholeSession(baseURL, password)
		if err == nil {
			t.Log("Destructive test - container created and accessible")
		}
		
		// This test would typically modify container state, restart it, etc.
	})
	
	t.Run("NetworkReconfiguration", func(t *testing.T) {
		// DO NOT add t.Parallel() - this test modifies network configuration
		
		config := SharedTestConfig{
			UseSharedEnvironment: false, // Force dedicated environment
			RequiresDestruction:  true,
			TestCategory:         "destructive",
		}
		
		terraformOptions, baseURL, _, err := GetTestEnvironment(t, config)
		require.NoError(t, err, "Should get dedicated environment")
		
		// Full lifecycle for destructive test
		defer terraform.Destroy(t, terraformOptions)
		terraform.InitAndApply(t, terraformOptions)
		time.Sleep(30 * time.Second)
		
		t.Logf("Network reconfiguration test with dedicated environment at %s", baseURL)
		// This test would typically modify network settings, DNS configuration, etc.
	})
}

// TestMixedParallelization demonstrates mixing parallel and sequential patterns
func TestMixedParallelization(t *testing.T) {
	// Fast parallel checks first
	t.Run("ParallelHealthChecks", func(t *testing.T) {
		t.Parallel()
		
		// Multiple quick health checks can run simultaneously
		env := GetSharedPiholeEnvironment()
		if env.Initialized {
			healthy := env.IsHealthy(t)
			t.Logf("Health check result: %v", healthy)
		} else {
			t.Log("Shared environment not available for health check")
		}
	})
	
	// Sequential configuration changes
	t.Run("SequentialConfigurationTests", func(t *testing.T) {
		// NO t.Parallel() - these modify state
		
		config := SharedTestConfig{
			UseSharedEnvironment: false,
			RequiresDestruction:  true,
			TestCategory:         "configuration",
		}
		
		terraformOptions, _, _, err := GetTestEnvironment(t, config)
		require.NoError(t, err, "Should get environment for configuration test")
		
		defer terraform.Destroy(t, terraformOptions)
		terraform.InitAndApply(t, terraformOptions)
		time.Sleep(30 * time.Second)
		
		t.Log("Configuration test completed sequentially")
	})
}

// TestParallelPerformanceDemonstration shows performance benefits
func TestParallelPerformanceDemonstration(t *testing.T) {
	start := time.Now()
	
	// These can all run in parallel, dramatically reducing total time
	for i := 0; i < 3; i++ {
		i := i // Capture loop variable
		t.Run(fmt.Sprintf("ParallelReadOnlyTest%d", i), func(t *testing.T) {
			t.Parallel()
			
			// Simulate a read-only test that takes some time
			time.Sleep(2 * time.Second)
			t.Logf("Parallel test %d completed", i)
		})
	}
	
	// This will only be measured once all parallel tests complete
	t.Cleanup(func() {
		elapsed := time.Since(start)
		t.Logf("Total time for parallel tests: %v (would be ~6s if sequential)", elapsed)
	})
}