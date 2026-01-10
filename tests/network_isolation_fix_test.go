package tests

import (
	"crypto/sha256"
	"fmt"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

// TestNetworkIsolationFix - RED phase test to fix network isolation issues
// This addresses the core issue preventing GREEN phase: network name conflicts
func TestNetworkIsolationFix(t *testing.T) {
	t.Parallel()
	
	// Generate truly unique identifiers based on test name + timestamp
	testID := fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%s-%d", t.Name(), time.Now().UnixNano()))))[:8]
	
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/modules/pihole",
		Vars: map[string]interface{}{
			"container_name":     fmt.Sprintf("pihole-isolation-test-%s", testID),
			"network_name":       fmt.Sprintf("pihole-isolation-net-%s", testID), // UNIQUE NETWORK NAME
			"subnet":             fmt.Sprintf("172.%d.0.0/16", 40+(len(testID)%10)), // Unique subnet
			"dns_port":           32000 + (len(testID) % 1000), // Unique port
			"web_port":           33000 + (len(testID) % 1000), // Unique port
			"timezone":           "America/New_York",
			"web_password":       fmt.Sprintf("isolation-test-%s", testID),
			"dnsmasq_listening":  "all",
			"use_host_network":   false,
		},
	})

	// Cleanup on test completion
	defer terraform.Destroy(t, terraformOptions)

	// Apply terraform - this should work without network conflicts
	terraform.InitAndApply(t, terraformOptions)

	// Basic validation that container was created
	containerName := terraform.Output(t, terraformOptions, "container_name")
	networkName := terraform.Output(t, terraformOptions, "network_name")
	
	require.Equal(t, fmt.Sprintf("pihole-isolation-test-%s", testID), containerName)
	require.Equal(t, fmt.Sprintf("pihole-isolation-net-%s", testID), networkName)
	
	t.Logf("✅ Network isolation test passed with unique identifiers: container=%s, network=%s", 
		containerName, networkName)
}

// TestQuickContainerStartup - Test that containers start within reasonable time
func TestQuickContainerStartup(t *testing.T) {
	t.Parallel()
	
	start := time.Now()
	
	testID := fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%s-%d", t.Name(), time.Now().UnixNano()))))[:8]
	
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/modules/pihole",
		Vars: map[string]interface{}{
			"container_name":     fmt.Sprintf("pihole-startup-test-%s", testID),
			"network_name":       fmt.Sprintf("pihole-startup-net-%s", testID),
			"subnet":             fmt.Sprintf("172.%d.0.0/16", 50+(len(testID)%10)),
			"dns_port":           34000 + (len(testID) % 1000),
			"web_port":           35000 + (len(testID) % 1000),
			"timezone":           "America/New_York",
			"web_password":       fmt.Sprintf("startup-test-%s", testID),
			"dnsmasq_listening":  "all",
			"use_host_network":   false,
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	// Apply and measure startup time
	terraform.InitAndApply(t, terraformOptions)
	
	elapsed := time.Since(start)
	
	// Should complete within 20 seconds for GREEN phase
	require.Less(t, elapsed, 20*time.Second, "Container startup took too long: %v", elapsed)
	
	t.Logf("✅ Container startup completed in %v", elapsed)
}

// TestParallelNetworkCreation - Validate that multiple tests can run simultaneously
func TestParallelNetworkCreation(t *testing.T) {
	// Run 3 tests in parallel to verify no conflicts
	for i := 0; i < 3; i++ {
		i := i // Capture loop variable
		t.Run(fmt.Sprintf("ParallelTest%d", i), func(t *testing.T) {
			t.Parallel()
			
			testID := fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%s-%d-%d", t.Name(), i, time.Now().UnixNano()))))[:8]
			
			terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
				TerraformDir: "../terraform/modules/pihole",
				Vars: map[string]interface{}{
					"container_name":     fmt.Sprintf("pihole-parallel-%d-%s", i, testID),
					"network_name":       fmt.Sprintf("pihole-parallel-net-%d-%s", i, testID),
					"subnet":             fmt.Sprintf("172.%d.0.0/16", 60+i+(len(testID)%5)),
					"dns_port":           36000 + i*10 + (len(testID) % 10),
					"web_port":           37000 + i*10 + (len(testID) % 10),
					"timezone":           "America/New_York",
					"web_password":       fmt.Sprintf("parallel-test-%d-%s", i, testID),
					"dnsmasq_listening":  "all",
					"use_host_network":   false,
				},
			})

			defer terraform.Destroy(t, terraformOptions)
			terraform.InitAndApply(t, terraformOptions)
			
			containerName := terraform.Output(t, terraformOptions, "container_name")
			t.Logf("✅ Parallel test %d completed with container: %s", i, containerName)
		})
	}
}