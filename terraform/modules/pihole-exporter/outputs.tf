# Pi-hole Exporter Module Outputs

output "container_id" {
  description = "Docker container ID of the Pi-hole Exporter"
  value       = docker_container.pihole_exporter.id
}

output "container_name" {
  description = "Name of the Pi-hole Exporter container"
  value       = docker_container.pihole_exporter.name
}

output "metrics_endpoint" {
  description = "Prometheus metrics endpoint URL"
  value       = "http://localhost:${var.exporter_port}${var.metrics_path}"
}

output "exporter_port" {
  description = "Port for the metrics endpoint"
  value       = var.exporter_port
}

output "pihole_target" {
  description = "Pi-hole instance being monitored"
  value       = var.pihole_hostname
}

output "scrape_config" {
  description = "Prometheus scrape configuration"
  value = {
    job_name        = "pihole-exporter"
    static_configs = [{
      targets = ["localhost:${var.exporter_port}"]
    }]
    metrics_path    = var.metrics_path
    scrape_interval = var.scrape_interval
  }
}

output "prometheus_labels" {
  description = "Container labels for Prometheus discovery"
  value = var.enable_prometheus_labels ? {
    "prometheus.io/scrape" = "true"
    "prometheus.io/port"   = tostring(var.exporter_port)
    "prometheus.io/path"   = var.metrics_path
  } : null
}

output "service_summary" {
  description = "Complete Pi-hole Exporter service configuration"
  value = {
    container_id     = docker_container.pihole_exporter.id
    container_name   = docker_container.pihole_exporter.name
    metrics_endpoint = "http://localhost:${var.exporter_port}${var.metrics_path}"
    pihole_target    = var.pihole_hostname
    scrape_interval  = var.scrape_interval
    image           = var.exporter_image
  }
}