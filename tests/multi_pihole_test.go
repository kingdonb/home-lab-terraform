package tests

import (
	"path/filepath"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// TestMultiPiholeInfrastructure verifies that we can deploy multiple pi-hole
// instances with proper DNS failover configuration
func TestMultiPiholeInfrastructure(t *testing.T) {
	// Arrange: Set up test inputs for multi-instance deployment
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: filepath.Join("..", "terraform", "environments", "test"),
		Vars: map[string]interface{}{
			"primary_dns_port":   53,
			"secondary_dns_port": 5353, // Non-standard port for testing
			"sync_enabled":       true,
			"synology_host":      "test.local",
		},
		NoColor: true,
	})

	// Clean up resources at the end of the test
	defer terraform.Destroy(t, terraformOptions)

	// Act & Assert: Initialize and validate the infrastructure
	terraform.InitAndValidate(t, terraformOptions)

	// Act: Plan the infrastructure
	planResult := terraform.Plan(t, terraformOptions)

	// Assert: Verify expected resources are planned
	assert.Contains(t, planResult, "module.primary_pihole")
	assert.Contains(t, planResult, "module.secondary_pihole")
	assert.Contains(t, planResult, "docker_network.dns_network")
	
	// Verify shared configuration volume
	assert.Contains(t, planResult, "docker_volume.pihole_shared_config")
}

// TestSynologySSHConnection verifies SSH connectivity to Synology devices
// using credentials from 1Password without exposing them
func TestSynologySSHConnection(t *testing.T) {
	// This test verifies SSH provider configuration
	// without actually connecting during test
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: filepath.Join("..", "terraform", "modules", "synology-ssh"),
		Vars: map[string]interface{}{
			"host": "test.synology.local",
			"user": "testuser",
		},
		NoColor: true,
	})

	// Act: Validate configuration
	terraform.InitAndValidate(t, terraformOptions)
	
	// Assert: Plan should include SSH connection resource
	planResult := terraform.Plan(t, terraformOptions)
	assert.Contains(t, planResult, "ssh_resource.synology_connection")
}

// TestCredentialSafety verifies that no secrets are committed to version control
func TestCredentialSafety(t *testing.T) {
	// Test that credential access uses 1Password CLI and doesn't leak
	// This would be implemented to scan for patterns that shouldn't exist
	t.Skip("Credential safety validation - to be implemented")
}

// TestDNSFailover tests that secondary DNS takes over when primary fails
func TestDNSFailover(t *testing.T) {
	// Integration test for DNS failover functionality
	t.Skip("DNS failover integration test - to be implemented")
}