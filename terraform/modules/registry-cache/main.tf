# Registry Pull-Through Cache Module
# Based on start_caching_docker_proxies script configuration

terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 3.0"
    }
  }
}

# Local variables for registry configuration
locals {
  registry_configs = {
    "docker.io"         = "https://registry-1.docker.io"
    "registry.k8s.io"   = "https://registry.k8s.io"
    "quay.io"           = "https://quay.io"
    "gcr.io"            = "https://gcr.io"
    "ghcr.io"           = "https://ghcr.io"
  }
  
  upstream_url = var.upstream_url != null ? var.upstream_url : local.registry_configs[var.registry_name]
  container_name = "registry-${var.registry_name}"
}

# Docker network for registry isolation (optional)
resource "docker_network" "registry_network" {
  count = var.create_network ? 1 : 0
  name  = var.network_name
  
  ipam_config {
    subnet = var.network_subnet
  }
}

# Docker volume for registry cache storage
resource "docker_volume" "registry_cache" {
  name = "${local.container_name}-cache"
  
  labels = {
    purpose = "registry-cache"
    registry = var.registry_name
    managed_by = "terraform"
  }
}

# Registry container with pull-through cache configuration
resource "docker_container" "registry_cache" {
  name  = local.container_name
  image = var.registry_image
  
  restart = "always"
  
  ports {
    internal = 5000
    external = var.cache_port
    protocol = "tcp"
    ip       = "0.0.0.0"
  }
  
  env = [
    "REGISTRY_PROXY_REMOTEURL=${local.upstream_url}",
    "REGISTRY_STORAGE_FILESYSTEM_ROOTDIRECTORY=/var/lib/registry",
    "REGISTRY_HTTP_ADDR=0.0.0.0:5000"
  ]
  
  volumes {
    volume_name    = docker_volume.registry_cache.name
    container_path = "/var/lib/registry"
    read_only      = false
  }
  
  # Optional network attachment
  dynamic "networks_advanced" {
    for_each = var.create_network ? [docker_network.registry_network[0].name] : []
    content {
      name = networks_advanced.value
    }
  }
  
  # Health check
  healthcheck {
    test         = ["CMD", "wget", "--quiet", "--tries=1", "--spider", "http://localhost:5000/v2/"]
    interval     = "30s"
    timeout      = "10s"
    retries      = 3
    start_period = "10s"
  }
  
  # Logging configuration
  log_driver = "json-file"
  log_opts = {
    max-size = "10m"
    max-file = "3"
  }
  
  labels {
    label = "purpose"
    value = "registry-cache"
  }
  
  labels {
    label = "registry"
    value = var.registry_name
  }
  
  labels {
    label = "managed_by"
    value = "terraform"
  }
}