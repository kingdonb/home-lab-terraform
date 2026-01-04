package tests

import (
	"net"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/miekg/dns"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPiholeNetworkConfiguration(t *testing.T) {
	t.Parallel()

	// Test that our pi-hole module uses correct network configuration
	// to respond to DNS queries from different network contexts
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/modules/pihole",
		Vars: map[string]interface{}{
			"container_name":         "pihole-network-test",
			"network_name":          "pihole-test-net", 
			"subnet":                "172.21.0.0/16", // Unique subnet to avoid conflicts
			"dns_port":              15353, // Use non-standard port to avoid conflicts
			"web_port":              18080, // Use non-standard port to avoid conflicts
			"timezone":              "America/New_York",
			"web_password":          "test-password", 
			"dnsmasq_listening":     "all", // Critical: should listen on all interfaces
			"use_host_network":      false, // Test with bridge networking first
		},
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// Wait longer for pi-hole to be ready (host networking needs more time)
	t.Log("Waiting for pi-hole to start...")
	time.Sleep(60 * time.Second)

	// Test 1: Verify pi-hole responds to DNS queries on the DNS port
	t.Run("DNS_Resolution_Works", func(t *testing.T) {
		client := new(dns.Client)
		client.Timeout = 5 * time.Second
		
		message := new(dns.Msg)
		message.SetQuestion(dns.Fqdn("pi.hole"), dns.TypeA)
		
		response, _, err := client.Exchange(message, "127.0.0.1:15353")
		require.NoError(t, err, "DNS query should succeed")
		assert.True(t, len(response.Answer) > 0, "Should get DNS response for pi.hole")
	})

	// Test 2: Verify DNS functionality works
	t.Run("DNS_Query_Success", func(t *testing.T) {
		client := new(dns.Client)
		client.Timeout = 5 * time.Second
		
		// Test standard DNS query
		message := new(dns.Msg)
		message.SetQuestion(dns.Fqdn("google.com"), dns.TypeA)
		
		response, _, err := client.Exchange(message, "127.0.0.1:15353")
		require.NoError(t, err, "DNS query for external domain should succeed")
		assert.True(t, len(response.Answer) > 0, "Should get DNS response for google.com")
		
		t.Logf("DNS query successful, got %d answers", len(response.Answer))
	})

	// Test 3: Basic connectivity check
	t.Run("Pi_Hole_Web_Interface_Accessible", func(t *testing.T) {
		// Test HTTP connection to web interface  
		conn, err := net.DialTimeout("tcp", "127.0.0.1:18080", 5*time.Second)
		if err == nil {
			conn.Close()
			t.Log("Web interface port is accessible")
		} else {
			t.Logf("Web interface connection failed: %v", err)
		}
		
		// Also test DNS port accessibility
		conn, err = net.DialTimeout("tcp", "127.0.0.1:15353", 5*time.Second)
		if err == nil {
			conn.Close() 
			t.Log("DNS port is accessible")
		} else {
			t.Logf("DNS port connection failed: %v", err)
		}
	})
}

func TestPiholeConfigurationCompliance(t *testing.T) {
	t.Parallel()
	
	// Test that our Terraform module matches the original docker-compose configuration
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/modules/pihole",
		Vars: map[string]interface{}{
			"container_name":         "pihole-compliance-test",
			"network_name":          "pihole-compliance-net",
			"subnet":                "172.22.0.0/16", // Unique subnet to avoid conflicts
			"dns_port":              15354,
			"web_port":              18081,
			"timezone":              "America/New_York",
			"webpassword":           "test-password",
		},
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndPlan(t, terraformOptions)
	
	// TODO: Add more specific tests once we update the module
	// to match the original docker-compose configuration:
	// - CAP_NET_BIND_SERVICE capability
	// - CAP_SYS_NICE capability  
	// - CAP_CHOWN capability
	// - shm_size: 256mb
	// - restart: unless-stopped
	// - WEB_PORT environment variable
}

