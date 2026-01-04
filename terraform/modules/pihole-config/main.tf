terraform {
  required_providers {
    pihole = {
      source  = "lukaspustina/pihole"
      version = "~> 0.3"
    }
  }
}

# Configure the Pi-hole provider
provider "pihole" {
  url      = var.pihole_base_url
  password = var.pihole_password
}

# DNS records for internal services
resource "pihole_dns_record" "homelab_gateway" {
  domain = "gateway.homelab.local"
  ip     = "10.17.12.1"
}

resource "pihole_dns_record" "homelab_nas" {
  domain = "nas.homelab.local"
  ip     = "10.17.12.100"
}

resource "pihole_dns_record" "homelab_registry" {
  domain = "registry.homelab.local"
  ip     = "10.17.12.101"
}

# CNAME records for service aliases
resource "pihole_cname_record" "docker_registry" {
  domain = "docker.homelab.local"
  target = "registry.homelab.local"
}

resource "pihole_cname_record" "container_registry" {
  domain = "containers.homelab.local"
  target = "registry.homelab.local"
}