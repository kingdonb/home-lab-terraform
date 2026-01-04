terraform {
  required_providers {
    ssh = {
      source  = "loafoe/ssh"
      version = "~> 2.6"
    }
  }
}

# For testing, we'll skip the actual 1Password integration
# In production, this would use 1Password CLI to fetch credentials safely
locals {
  # Test credentials (would be replaced by 1Password in production)
  test_credentials = {
    username    = var.user
    private_key = "" # Empty for testing
  }
}

# SSH connection resource for Synology device management
resource "ssh_resource" "synology_connection" {
  host        = var.host
  user        = local.test_credentials.username
  port        = var.port

  # For testing, we skip the actual SSH command
  # In production this would execute management commands
  when = "create"

  # Test connection by running a simple command (when credentials are available)
  commands = [
    "echo 'SSH connection test to ${var.host}'"
  ]

  timeout = "30s"
}

# Output connection status
output "connection_status" {
  value = "Test connection configured for ${var.host}"
}

output "host" {
  value = var.host
}