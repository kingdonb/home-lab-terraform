# Matchbox Module
# Based on start_matchbox script configuration for PXE boot services

terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 3.0"
    }
  }
}

# Matchbox container for PXE boot management
resource "docker_container" "matchbox" {
  name  = var.container_name
  image = var.matchbox_image
  
  restart = "unless-stopped"
  
  # Network configuration - requires external VLAN network
  networks_advanced {
    name = var.vlan_network_name
    ipv4_address = var.static_ip
  }
  
  # Matchbox configuration arguments
  command = [
    "-address=:${var.matchbox_port}",
    "-log-level=${var.log_level}"
  ]
  
  # Optional volumes for profiles, groups, and assets
  dynamic "volumes" {
    for_each = var.data_volume_name != null ? [1] : []
    content {
      volume_name    = var.data_volume_name
      container_path = "/var/lib/matchbox"
      read_only      = false
    }
  }
  
  # Additional volumes for custom configuration
  dynamic "volumes" {
    for_each = var.additional_volumes
    content {
      host_path      = volumes.value.host_path
      container_path = volumes.value.container_path
      read_only      = try(volumes.value.read_only, false)
    }
  }
  
  # Environment variables for configuration
  env = concat([
    "MATCHBOX_ADDRESS=0.0.0.0:${var.matchbox_port}",
    "MATCHBOX_LOG_LEVEL=${var.log_level}"
  ], var.additional_env_vars)
  
  # Health check for Matchbox API
  healthcheck {
    test         = ["CMD", "wget", "--quiet", "--tries=1", "--spider", "http://localhost:${var.matchbox_port}"]
    interval     = "30s"
    timeout      = "10s"
    retries      = 3
    start_period = "30s"
  }
  
  # Logging configuration
  log_driver = "json-file"
  log_opts = {
    max-size = "10m"
    max-file = "3"
  }
  
  labels {
    label = "purpose"
    value = "pxe-boot-server"
  }
  
  labels {
    label = "network"
    value = var.vlan_network_name
  }
  
  labels {
    label = "managed_by"
    value = "terraform"
  }
}