# Pi-hole Exporter Module
# Based on start_pihole_exporter script configuration

terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 3.0"
    }
  }
}

# Pi-hole Exporter container for Prometheus metrics
resource "docker_container" "pihole_exporter" {
  name  = var.container_name
  image = var.exporter_image
  
  restart = "unless-stopped"
  
  # Use host networking for Pi-hole access
  network_mode = "host"
  
  ports {
    internal = var.exporter_port
    external = var.exporter_port
    protocol = "tcp"
    ip       = "0.0.0.0"
  }
  
  # Environment variables for Pi-hole connection
  env = concat([
    "PIHOLE_HOSTNAME=${var.pihole_hostname}",
    "PIHOLE_API_TOKEN=${var.pihole_api_token}",
    "INTERVAL=${var.scrape_interval}",
    "PORT=${var.exporter_port}"
  ], var.additional_env_vars)
  
  # Health check for exporter endpoint
  healthcheck {
    test         = ["CMD", "wget", "--quiet", "--tries=1", "--spider", "http://localhost:${var.exporter_port}/metrics"]
    interval     = "30s"
    timeout      = "10s"
    retries      = 3
    start_period = "15s"
  }
  
  # Logging configuration
  log_driver = "json-file"
  log_opts = {
    max-size = "5m"
    max-file = "3"
  }
  
  labels {
    label = "purpose"
    value = "pihole-metrics-exporter"
  }
  
  labels {
    label = "prometheus.io/scrape"
    value = "true"
  }
  
  labels {
    label = "prometheus.io/port"
    value = tostring(var.exporter_port)
  }
  
  labels {
    label = "prometheus.io/path"
    value = "/metrics"
  }
  
  labels {
    label = "managed_by"
    value = "terraform"
  }
}