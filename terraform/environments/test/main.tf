terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker" 
      version = "~> 3.0"
    }
  }
}

# Shared network for DNS infrastructure
resource "docker_network" "dns_network" {
  name = "dns-infrastructure"
  
  ipam_config {
    subnet = "172.21.0.0/16"
  }
}

# Shared configuration volume for pi-hole sync
resource "docker_volume" "pihole_shared_config" {
  name = "pihole-shared-blocklists"
}

# Primary Pi-hole instance
module "primary_pihole" {
  source = "../../modules/pihole"
  
  container_name = "pihole-primary"
  network_name   = docker_network.dns_network.name
  dns_port      = var.primary_dns_port
  web_port      = 8080
  timezone      = "America/New_York"
  
  # Additional volume for shared configuration
  extra_volumes = [{
    volume_name    = docker_volume.pihole_shared_config.name
    container_path = "/shared"
  }]
}

# Secondary Pi-hole instance  
module "secondary_pihole" {
  source = "../../modules/pihole"
  
  container_name = "pihole-secondary"
  network_name   = docker_network.dns_network.name
  dns_port      = var.secondary_dns_port
  web_port      = 8081
  timezone      = "America/New_York"
  
  # Additional volume for shared configuration
  extra_volumes = [{
    volume_name    = docker_volume.pihole_shared_config.name
    container_path = "/shared"
  }]
}

# Configure DNS records in primary Pi-hole
module "pihole_config" {
  source = "../../modules/pihole-config"
  
  pihole_base_url = "http://localhost:8080"
  pihole_password = var.pihole_password
  
  # Wait for Pi-hole to be ready
  depends_on = [module.primary_pihole]
}