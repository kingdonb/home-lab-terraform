terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 3.0"
    }
  }
}

# Configure the Docker Provider
provider "docker" {}

# Create a Docker network for pi-hole
resource "docker_network" "pihole_network" {
  name = var.network_name
  
  ipam_config {
    subnet = var.subnet
  }
}

# Create volumes for pi-hole data persistence
resource "docker_volume" "pihole_data" {
  name = "${var.container_name}-data"
}

resource "docker_volume" "pihole_dnsmasq" {
  name = "${var.container_name}-dnsmasq"
}

# Pull pi-hole Docker image
resource "docker_image" "pihole" {
  name = "pihole/pihole:${var.pihole_version}"
}

# Create pi-hole container
resource "docker_container" "pihole" {
  name  = var.container_name
  image = docker_image.pihole.image_id
  
  restart = "unless-stopped"
  
  # Environment variables
  env = [
    "TZ=${var.timezone}",
    "WEBPASSWORD=${var.web_password}",
    "PIHOLE_DNS_=${var.upstream_dns}",
    "DNSMASQ_LISTENING=${var.dnsmasq_listening}",
  ]
  
  # Port mappings
  ports {
    internal = 53
    external = var.dns_port
    protocol = "tcp"
  }
  
  ports {
    internal = 53
    external = var.dns_port
    protocol = "udp"
  }
  
  ports {
    internal = 80
    external = var.web_port
    protocol = "tcp"
  }
  
  # Volume mounts
  volumes {
    volume_name    = docker_volume.pihole_data.name
    container_path = "/etc/pihole"
  }
  
  volumes {
    volume_name    = docker_volume.pihole_dnsmasq.name
    container_path = "/etc/dnsmasq.d"
  }
  
  # Connect to network
  networks_advanced {
    name = docker_network.pihole_network.name
  }
  
  # Healthcheck
  healthcheck {
    test         = ["CMD", "dig", "@127.0.0.1", "-p", "53", "pi.hole", "+short"]
    interval     = "30s"
    timeout      = "10s"
    start_period = "60s"
    retries      = 3
  }
}