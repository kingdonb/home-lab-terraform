terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 3.0"
    }
  }
}

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
  shm_size = var.shm_size_mb
  
  # Use host networking if specified (required for multi-subnet DNS)
  network_mode = var.use_host_network ? "host" : "bridge"
  
  # Environment variables (Pi-hole v6+ format)
  env = [
    "TZ=${var.timezone}",
    "FTLCONF_webserver_api_password=${var.web_password}",
    "PIHOLE_DNS_=${var.upstream_dns}", 
    "FTLCONF_dns_listeningMode=${upper(var.dnsmasq_listening)}",
    "WEB_PORT=${var.web_port}",
  ]
  
  # Linux capabilities (required for DNS binding)
  capabilities {
    add = var.capabilities
  }
  
  # Port mappings (only when NOT using host networking)
  dynamic "ports" {
    for_each = var.use_host_network ? [] : [1]
    content {
      internal = 53
      external = var.dns_port
      protocol = "tcp"
    }
  }
  
  dynamic "ports" {
    for_each = var.use_host_network ? [] : [1]
    content {
      internal = 53
      external = var.dns_port
      protocol = "udp"
    }
  }
  
  dynamic "ports" {
    for_each = var.use_host_network ? [] : [1]
    content {
      internal = 80
      external = var.web_port
      protocol = "tcp"
    }
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
  
  # Additional volumes (for shared config, etc.)
  dynamic "volumes" {
    for_each = var.extra_volumes
    content {
      volume_name    = volumes.value.volume_name
      container_path = volumes.value.container_path
    }
  }
  
  # Connect to network (only if not using host networking)
  dynamic "networks_advanced" {
    for_each = var.use_host_network ? [] : [1]
    content {
      name = docker_network.pihole_network.name
    }
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