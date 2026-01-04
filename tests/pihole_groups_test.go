package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// PiholeGroup represents a Pi-hole group configuration
type PiholeGroup struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
}

// PiholeClient represents a Pi-hole client configuration
type PiholeClient struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	IP       string   `json:"ip"`
	MAC      string   `json:"mac"`
	Groups   []int    `json:"groups"`
	Comment  string   `json:"comment"`
}

// PiholeDomain represents a Pi-hole domain/regex entry
type PiholeDomain struct {
	ID        int    `json:"id"`
	Domain    string `json:"domain"`
	Type      string `json:"type"` // "regex", "exact", "wildcard"
	Groups    []int  `json:"groups"`
	Comment   string `json:"comment"`
	Enabled   bool   `json:"enabled"`
}

// CreateGroup creates a new group via Pi-hole API
func (s *PiholeSession) CreateGroup(name, description string, enabled bool) (*PiholeGroup, error) {
	// Pi-hole v6+ group creation endpoint
	payload := map[string]interface{}{
		"name":        name,
		"description": description,
		"enabled":     enabled,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal group data: %v", err)
	}

	req, err := http.NewRequest("POST", s.BaseURL+"/api/groups", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create group request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	if s.CSRFToken != "" {
		req.Header.Set("X-Pi-hole-Token", s.CSRFToken)
	}

	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("group creation request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("group creation failed with status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	var group PiholeGroup
	if err := json.Unmarshal(body, &group); err != nil {
		return nil, fmt.Errorf("failed to parse group response: %v", err)
	}

	return &group, nil
}

// GetGroups retrieves all groups from Pi-hole
func (s *PiholeSession) GetGroups() ([]PiholeGroup, error) {
	req, err := http.NewRequest("GET", s.BaseURL+"/api/groups", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create groups request: %v", err)
	}

	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("groups request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("groups API returned status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	var result struct {
		Groups []PiholeGroup `json:"groups"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		// Try direct array format
		var groups []PiholeGroup
		if err := json.Unmarshal(body, &groups); err != nil {
			return nil, fmt.Errorf("failed to parse groups response: %v", err)
		}
		return groups, nil
	}

	return result.Groups, nil
}

// CreateClient creates a new client via Pi-hole API
func (s *PiholeSession) CreateClient(name, ip, mac string, groups []int, comment string) (*PiholeClient, error) {
	payload := map[string]interface{}{
		"name":    name,
		"ip":      ip,
		"mac":     mac,
		"groups":  groups,
		"comment": comment,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal client data: %v", err)
	}

	req, err := http.NewRequest("POST", s.BaseURL+"/api/clients", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create client request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	if s.CSRFToken != "" {
		req.Header.Set("X-Pi-hole-Token", s.CSRFToken)
	}

	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client creation request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("client creation failed with status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	var client PiholeClient
	if err := json.Unmarshal(body, &client); err != nil {
		return nil, fmt.Errorf("failed to parse client response: %v", err)
	}

	return &client, nil
}

// CreateDomainRegex creates a regex domain entry via Pi-hole API
func (s *PiholeSession) CreateDomainRegex(domain string, groups []int, comment string) (*PiholeDomain, error) {
	payload := map[string]interface{}{
		"domain":  domain,
		"type":    "regex",
		"groups":  groups,
		"comment": comment,
		"enabled": true,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal domain data: %v", err)
	}

	req, err := http.NewRequest("POST", s.BaseURL+"/api/domains", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create domain request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	if s.CSRFToken != "" {
		req.Header.Set("X-Pi-hole-Token", s.CSRFToken)
	}

	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("domain creation request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("domain creation failed with status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	var domainResult PiholeDomain
	if err := json.Unmarshal(body, &domainResult); err != nil {
		return nil, fmt.Errorf("failed to parse domain response: %v", err)
	}

	return &domainResult, nil
}

func TestPiholeGroupManagement(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: filepath.Join("..", "terraform", "modules", "pihole"),
		Vars: map[string]interface{}{
			"container_name":     "pihole-groups-test",
			"network_name":       "pihole-groups-net",
			"subnet":             "172.26.0.0/16", // Unique subnet to avoid conflicts
			"dns_port":           27353, // Unique port for this test
			"web_port":           27080,
			"timezone":           "America/New_York",
			"web_password":       "groups-test-password",
			"dnsmasq_listening":  "all",
			"use_host_network":   false,
		},
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// Wait for Pi-hole to be ready
	t.Log("Waiting for Pi-hole groups test instance...")
	time.Sleep(45 * time.Second)

	baseURL := "http://localhost:27080"
	password := "groups-test-password"

	// Create authenticated session
	session, err := NewPiholeSession(baseURL, password)
	require.NoError(t, err, "Should be able to create authenticated session")

	// Test 1: Create custom groups
	t.Run("Create_Custom_Groups", func(t *testing.T) {
		// Expected groups with descriptions
		expectedGroups := map[string]string{
			"Socials":     "Social media sites",
			"Cryptos":     "Gambling and stuff", 
			"Advertising": "Ad safety (the default adlist)",
		}

		createdGroups := make(map[string]*PiholeGroup)
		
		for name, description := range expectedGroups {
			group, err := session.CreateGroup(name, description, true)
			require.NoError(t, err, "Should be able to create group %s", name)
			require.NotNil(t, group, "Created group should not be nil")
			assert.Equal(t, name, group.Name, "Group name should match")
			assert.Equal(t, description, group.Description, "Group description should match")
			assert.True(t, group.Enabled, "Group should be enabled")
			
			createdGroups[name] = group
			t.Logf("Created group: %s (ID: %d) - %s", group.Name, group.ID, group.Description)
		}

		// Verify groups can be retrieved
		groups, err := session.GetGroups()
		require.NoError(t, err, "Should be able to get groups")
		
		// Check that our custom groups exist
		foundGroups := 0
		for _, group := range groups {
			if expectedDesc, exists := expectedGroups[group.Name]; exists {
				assert.Equal(t, expectedDesc, group.Description, "Group %s description should match", group.Name)
				foundGroups++
			}
		}
		
		assert.Equal(t, len(expectedGroups), foundGroups, "Should find all created custom groups")
	})

	// Test 2: Create client configurations
	t.Run("Create_Client_Configurations", func(t *testing.T) {
		// First get the groups we need
		groups, err := session.GetGroups()
		require.NoError(t, err, "Should be able to get groups for client setup")
		
		groupMap := make(map[string]int)
		for _, group := range groups {
			groupMap[group.Name] = group.ID
		}
		
		// Define test clients with their group assignments
		testClients := []struct {
			name    string
			ip      string
			mac     string
			groups  []string // Group names
			comment string
		}{
			{
				name:    "work-laptop",
				ip:      "10.17.12.100",
				mac:     "00:11:22:33:44:55",
				groups:  []string{"Socials", "Cryptos", "Advertising"},
				comment: "Work laptop - full restrictions",
			},
			{
				name:    "work-phone", 
				ip:      "10.17.12.101",
				mac:     "00:11:22:33:44:56",
				groups:  []string{"Socials", "Cryptos", "Advertising"},
				comment: "Work phone - full restrictions",
			},
			{
				name:    "other-laptop",
				ip:      "10.17.13.100",
				mac:     "00:11:22:33:44:57",
				groups:  []string{"Socials", "Cryptos", "Advertising"},
				comment: "Other laptop - full restrictions",
			},
			{
				name:    "phone-2.4ghz",
				ip:      "10.17.13.101", 
				mac:     "00:11:22:33:44:58",
				groups:  []string{"Socials", "Cryptos", "Advertising"},
				comment: "Phone on 2.4GHz - full restrictions",
			},
		}

		for _, clientDef := range testClients {
			// Convert group names to IDs
			var groupIDs []int
			for _, groupName := range clientDef.groups {
				if groupID, exists := groupMap[groupName]; exists {
					groupIDs = append(groupIDs, groupID)
				}
			}
			
			client, err := session.CreateClient(
				clientDef.name,
				clientDef.ip,
				clientDef.mac,
				groupIDs,
				clientDef.comment,
			)
			
			require.NoError(t, err, "Should be able to create client %s", clientDef.name)
			require.NotNil(t, client, "Created client should not be nil")
			assert.Equal(t, clientDef.name, client.Name, "Client name should match")
			assert.Equal(t, clientDef.ip, client.IP, "Client IP should match")
			
			t.Logf("Created client: %s (%s) with groups %v", client.Name, client.IP, client.Groups)
		}
	})

	// Test 3: Create domain regex entries
	t.Run("Create_Domain_Regex_Entries", func(t *testing.T) {
		// Get group IDs for assignment
		groups, err := session.GetGroups()
		require.NoError(t, err, "Should be able to get groups for domain setup")
		
		groupMap := make(map[string]int)
		for _, group := range groups {
			groupMap[group.Name] = group.ID
		}
		
		// Sample crypto and social media regex patterns (representative of the ~15 entries)
		regexEntries := []struct {
			pattern string
			groups  []string
			comment string
		}{
			{
				pattern: `^(.+\.)?coinbase\.com$`,
				groups:  []string{"Cryptos"},
				comment: "Block Coinbase - crypto trading",
			},
			{
				pattern: `^(.+\.)?binance\.(com|us)$`,
				groups:  []string{"Cryptos"},
				comment: "Block Binance - crypto exchange",
			},
			{
				pattern: `^(.+\.)?facebook\.com$`,
				groups:  []string{"Socials"},
				comment: "Block Facebook - social media",
			},
			{
				pattern: `^(.+\.)?instagram\.com$`,
				groups:  []string{"Socials"},
				comment: "Block Instagram - social media",
			},
			{
				pattern: `^(.+\.)?twitter\.com$`,
				groups:  []string{"Socials"},
				comment: "Block Twitter - social media",
			},
			{
				pattern: `^(.+\.)?x\.com$`,
				groups:  []string{"Socials"},
				comment: "Block X.com - social media",
			},
			{
				pattern: `^(.+\.)?tiktok\.com$`,
				groups:  []string{"Socials"},
				comment: "Block TikTok - social media",
			},
			{
				pattern: `^(.+\.)?reddit\.com$`,
				groups:  []string{"Socials"},
				comment: "Block Reddit - social media",
			},
		}

		for _, entry := range regexEntries {
			// Convert group names to IDs
			var groupIDs []int
			for _, groupName := range entry.groups {
				if groupID, exists := groupMap[groupName]; exists {
					groupIDs = append(groupIDs, groupID)
				}
			}
			
			domain, err := session.CreateDomainRegex(
				entry.pattern,
				groupIDs,
				entry.comment,
			)
			
			require.NoError(t, err, "Should be able to create regex entry: %s", entry.pattern)
			require.NotNil(t, domain, "Created domain should not be nil")
			assert.Equal(t, entry.pattern, domain.Domain, "Domain pattern should match")
			assert.True(t, domain.Enabled, "Domain should be enabled")
			
			t.Logf("Created regex domain: %s for groups %v", domain.Domain, domain.Groups)
		}
	})
}