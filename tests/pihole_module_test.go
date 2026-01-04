package tests

import (
	"path/filepath"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// TestPiholeModuleGeneratesValidConfig verifies that the pi-hole module
// generates valid Terraform configuration with expected outputs
func TestPiholeModuleGeneratesValidConfig(t *testing.T) {
	// Arrange: Set up test inputs
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: filepath.Join("..", "terraform", "modules", "pihole"),
		Vars: map[string]interface{}{
			"container_name": "pihole-test",
			"network_name":   "pihole-net",
			"dns_port":      53,
			"web_port":      8080,
			"timezone":      "America/New_York",
		},
		NoColor: true,
	})

	// Clean up resources at the end of the test
	defer terraform.Destroy(t, terraformOptions)

	// Act & Assert: Initialize and validate the module
	terraform.InitAndValidate(t, terraformOptions)
	
	// Act: Plan the infrastructure
	planResult := terraform.Plan(t, terraformOptions)
	
	// Assert: Verify expected resources are planned
	assert.Contains(t, planResult, "docker_container.pihole")
	assert.Contains(t, planResult, "docker_network.pihole_network")
	assert.Contains(t, planResult, "docker_volume.pihole_data")
}

// TestPiholeDockerComposeIntegration tests the Docker Compose setup
// This is our integration test to verify the pi-hole service starts correctly
func TestPiholeDockerComposeIntegration(t *testing.T) {
	// This test will be implemented once we have the Docker Compose file
	t.Skip("Docker Compose integration test - to be implemented")
}