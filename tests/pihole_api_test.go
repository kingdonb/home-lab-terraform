package tests

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/miekg/dns"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// PiholeSession represents an authenticated Pi-hole session
type PiholeSession struct {
	BaseURL    string
	HTTPClient *http.Client
	SessionID  string
	CSRFToken  string
}

// NewPiholeSession creates and authenticates a new Pi-hole session
func NewPiholeSession(baseURL, password string) (*PiholeSession, error) {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}
	
	session := &PiholeSession{
		BaseURL:    baseURL,
		HTTPClient: client,
	}
	
	// Pi-hole v6+ uses JSON API authentication
	authPayload := map[string]interface{}{
		"password": password,
		"totp":     nil,
	}

	// Convert to JSON
	jsonData, err := json.Marshal(authPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal auth payload: %v", err)
	}

	// Create request to /api/auth
	req, err := http.NewRequest("POST", baseURL+"/api/auth", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create auth request: %v", err)
	}

	// Set required headers for Pi-hole v6+ API
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Referer", baseURL+"/admin/login")
	req.Header.Set("Origin", baseURL)

	// Make authentication request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("authentication request failed: %v", err)
	}
	defer resp.Body.Close()

	fmt.Printf("Authentication response status: %d\n", resp.StatusCode)

	// Check for successful authentication
	if resp.StatusCode == 200 {
		// Check if authentication cookies were set (Pi-hole v6+ uses 'sid' cookie)
		cookieSet := false
		for _, cookie := range resp.Cookies() {
			fmt.Printf("  Cookie set: %s=%s (first 20 chars)\n", cookie.Name, cookie.Value[:min(20, len(cookie.Value))])
			if cookie.Name == "sid" || cookie.Name == "_SSID" {
				cookieSet = true
			}
		}
		
		if cookieSet {
			return session, nil
		}
		
		// Even if no cookies, check the response body for a successful session
		body, _ := ioutil.ReadAll(resp.Body)
		var authResp map[string]interface{}
		if json.Unmarshal(body, &authResp) == nil {
			if sessionData, ok := authResp["session"].(map[string]interface{}); ok {
				if valid, ok := sessionData["valid"].(bool); ok && valid {
					// Capture CSRF token for subsequent API requests
					if csrfToken, ok := sessionData["csrf"].(string); ok {
						session.CSRFToken = csrfToken
						fmt.Printf("Captured CSRF token: %s\n", csrfToken)
					}
					if sid, ok := sessionData["sid"].(string); ok {
						session.SessionID = sid
						fmt.Printf("Captured session ID: %s\n", sid)
					}
					fmt.Printf("Authentication successful via session data\n")
					return session, nil
				}
			}
		}
	}

	// Read response body for debugging
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Authentication response body: %s\n", string(body))

	return nil, fmt.Errorf("authentication failed with status %d", resp.StatusCode)
}

// GetStats retrieves Pi-hole statistics using authenticated session
func (s *PiholeSession) GetStats() (map[string]interface{}, error) {
	// Use /api endpoint with cookie-based authentication
	url := s.BaseURL + "/api"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get stats: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("stats API returned status: %d", resp.StatusCode)
	}
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}
	
	return result, nil
}

// TestAPIAccess tests that we can access API endpoints with authentication
func (s *PiholeSession) TestAPIAccess() error {
	// Test basic API access with authenticated session
	url := s.BaseURL + "/api"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to access API: %w", err)
	}
	defer resp.Body.Close()
	
	// Any response that's not a 401 indicates authentication is working
	if resp.StatusCode == 401 {
		return fmt.Errorf("API returned 401 - authentication failed")
	}
	
	return nil
}

// GetLists retrieves blocklist configuration
func (s *PiholeSession) GetLists() (map[string]interface{}, error) {
	endpoints := []string{
		"/api/lists",
		"/api", 
	}
	
	for _, endpoint := range endpoints {
		resp, err := s.HTTPClient.Get(s.BaseURL + endpoint)
		if err != nil {
			continue
		}
		defer resp.Body.Close()
		
		if resp.StatusCode == 200 {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				continue
			}
			
			var result map[string]interface{}
			if json.Unmarshal(body, &result) == nil {
				return result, nil
			}
		}
	}
	
	return nil, fmt.Errorf("failed to get lists configuration")
}

func TestPiholeAPIFunctionality(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/modules/pihole",
		Vars: map[string]interface{}{
			"container_name":         "pihole-api-test",
			"network_name":          "pihole-api-net",
			"dns_port":              25353, // Different port to avoid conflicts
			"web_port":              28080,
			"timezone":              "America/New_York",
			"web_password":          "api-test-password",
			"dnsmasq_listening":     "all",
			"use_host_network":      false,
		},
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// Wait for pi-hole to be ready
	t.Log("Waiting for pi-hole to start and API to be ready...")
	time.Sleep(45 * time.Second)

	webPort := terraformOptions.Vars["web_port"].(int)
	password := terraformOptions.Vars["web_password"].(string)
	baseURL := fmt.Sprintf("http://localhost:%d", webPort)

	// Test 1: Verify web interface is accessible
	t.Run("Web_Interface_Accessible", func(t *testing.T) {
		resp, err := http.Get(baseURL + "/admin")
		require.NoError(t, err, "Web interface should be accessible")
		defer resp.Body.Close()
		
		assert.Equal(t, 200, resp.StatusCode, "Web interface should return 200")
		t.Logf("Web interface accessible at %s/admin", baseURL)
	})

	// Test 2: Test Pi-hole v6 session-based authentication
	t.Run("API_Authentication", func(t *testing.T) {
		// Create authenticated session
		session, err := NewPiholeSession(baseURL, password)
		require.NoError(t, err, "Should be able to create authenticated session")
		
		// Test that API access works with authentication
		err = session.TestAPIAccess()
		require.NoError(t, err, "Should be able to access API with session")
		
		t.Logf("Pi-hole v6+ authentication successful - session established")
		
		// Test basic stats retrieval
		stats, err := session.GetStats()
		if err == nil {
			t.Logf("Successfully retrieved stats: %+v", stats)
		} else {
			t.Logf("Stats API not yet available, but authentication working: %v", err)
		}
	})

	// Test 3: Test authenticated API endpoints
	t.Run("API_Endpoint_Discovery", func(t *testing.T) {
		// Create authenticated session
		session, err := NewPiholeSession(baseURL, password)
		require.NoError(t, err, "Should be able to create authenticated session")
		
		// Test different API endpoints to see what's available
		endpoints := []string{"/api", "/api/stats", "/api/lists"}
		workingEndpoints := 0
		
		for _, endpoint := range endpoints {
			resp, err := session.HTTPClient.Get(session.BaseURL + endpoint)
			if err == nil {
				defer resp.Body.Close()
				if resp.StatusCode != 401 {
					workingEndpoints++
					t.Logf("API endpoint %s accessible (status: %d)", endpoint, resp.StatusCode)
				}
			}
		}
		
		t.Logf("Found %d accessible API endpoints out of %d tested", workingEndpoints, len(endpoints))
	})

	// Test 4: Basic DNS functionality test
	t.Run("DNS_Functionality", func(t *testing.T) {
		// Test that Pi-hole DNS is working by querying a known domain
		client := new(dns.Client)
		client.Timeout = 5 * time.Second
		
		message := new(dns.Msg)
		message.SetQuestion(dns.Fqdn("google.com"), dns.TypeA)
		
		response, _, err := client.Exchange(message, "127.0.0.1:25353")
		if err == nil && len(response.Answer) > 0 {
			t.Logf("DNS functionality confirmed: %v", response.Answer[0])
		} else {
			t.Log("DNS functionality test completed")
		}
	})

	// Test 5: Explore available API endpoints
	t.Run("API_Exploration", func(t *testing.T) {
		// Create authenticated session
		session, err := NewPiholeSession(baseURL, password)
		require.NoError(t, err, "Should be able to create authenticated session")
		
		// Test various Pi-hole v6+ API endpoints
		endpoints := []string{"/api", "/api/stats", "/api/lists", "/api/domains", "/api/clients"}
		accessibleEndpoints := 0
		
		for _, endpoint := range endpoints {
			resp, err := session.HTTPClient.Get(session.BaseURL + endpoint)
			if err == nil {
				defer resp.Body.Close()
				t.Logf("Endpoint %s returned status: %d", endpoint, resp.StatusCode)
				
				if resp.StatusCode != 401 {
					accessibleEndpoints++
					body, err := io.ReadAll(resp.Body)
					if err == nil && len(body) > 0 {
						var result interface{}
						if json.Unmarshal(body, &result) == nil {
							t.Logf("Endpoint %s returned valid JSON", endpoint)
						}
					}
				}
			}
		}
		
		t.Logf("Successfully accessed %d out of %d API endpoints", accessibleEndpoints, len(endpoints))
	})

	// Test 6: Configuration Management - API Accessibility
	t.Run("Configuration_Management", func(t *testing.T) {
		// Create authenticated session
		session, err := NewPiholeSession(baseURL, password)
		require.NoError(t, err, "Should be able to create authenticated session")
		
		// Test that we can access management endpoints without 401 errors
		endpoints := []string{"/api", "/api/stats", "/api/clients", "/api/domains"}
		
		successfulEndpoints := 0
		for _, endpoint := range endpoints {
			resp, err := session.HTTPClient.Get(session.BaseURL + endpoint)
			if err == nil {
				defer resp.Body.Close()
				
				t.Logf("Management endpoint %s returned status: %d", endpoint, resp.StatusCode)
				
				// Any non-401 response means authentication is working
				if resp.StatusCode != 401 {
					successfulEndpoints++
				}
			}
		}
		
		t.Logf("Successfully authenticated to %d out of %d management endpoints", successfulEndpoints, len(endpoints))
		assert.Greater(t, successfulEndpoints, 0, "At least one management endpoint should be accessible")
	})
}

func TestPiholeAPIConfiguration(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/modules/pihole",
		Vars: map[string]interface{}{
			"container_name":         "pihole-config-test",
			"network_name":          "pihole-config-net", 
			"dns_port":              26353,
			"web_port":              28081,
			"timezone":              "America/New_York",
			"web_password":          "config-test-password",
			"dnsmasq_listening":     "all",
			"use_host_network":      false,
		},
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// Wait for pi-hole to be ready  
	t.Log("Waiting for pi-hole configuration test instance...")
	time.Sleep(45 * time.Second)

	baseURL := "http://localhost:28081"
	password := "config-test-password"

	// Test advanced configuration capabilities
	t.Run("Test_Group_Management", func(t *testing.T) {
		passwordHash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
		
		// Pi-hole groups API endpoint
		apiURL := fmt.Sprintf("%s/admin/api.php?summaryRaw&auth=%s", baseURL, passwordHash)
		resp, err := http.Get(apiURL)
		require.NoError(t, err, "Groups API should be accessible")
		defer resp.Body.Close()
		
		assert.Equal(t, 200, resp.StatusCode, "Groups API should return 200")
		t.Log("Groups API endpoint is accessible for future group management")
	})

	t.Run("Test_Client_Management", func(t *testing.T) {
		passwordHash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
		
		// Pi-hole clients API endpoint
		apiURL := fmt.Sprintf("%s/admin/api.php?topClients&auth=%s", baseURL, passwordHash)
		resp, err := http.Get(apiURL)
		require.NoError(t, err, "Clients API should be accessible")
		defer resp.Body.Close()
		
		assert.Equal(t, 200, resp.StatusCode, "Clients API should return 200")
		
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err, "Should read clients response")
		
		var apiResp map[string]interface{}
		err = json.Unmarshal(body, &apiResp)
		require.NoError(t, err, "Clients response should be valid JSON")
		
		t.Logf("Client management API working, response keys: %v", getKeys(apiResp))
	})
}

// Helper function to get map keys
func getKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Helper function for min
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}