# Matchbox Module Outputs

output "container_id" {
  description = "Docker container ID of the Matchbox server"
  value       = docker_container.matchbox.id
}

output "container_name" {
  description = "Name of the Matchbox container"
  value       = docker_container.matchbox.name
}

output "static_ip" {
  description = "Static IP address of the Matchbox server"
  value       = var.static_ip
}

output "api_endpoint" {
  description = "Matchbox API endpoint URL"
  value       = "http://${var.static_ip}:${var.matchbox_port}"
}

output "boot_ipxe_url" {
  description = "iPXE boot script URL for PXE clients"
  value       = "http://${var.static_ip}:${var.matchbox_port}/boot.ipxe"
}

output "matchbox_port" {
  description = "Matchbox server port"
  value       = var.matchbox_port
}

output "network_config" {
  description = "Network configuration summary"
  value = {
    vlan_network = var.vlan_network_name
    static_ip    = var.static_ip
    port         = var.matchbox_port
  }
}

output "server_config" {
  description = "Matchbox server configuration summary"
  value = {
    image      = var.matchbox_image
    log_level  = var.log_level
    https_enabled = var.enable_https
    data_volume = var.data_volume_name
  }
}

output "service_summary" {
  description = "Complete Matchbox service configuration"
  value = {
    container_id   = docker_container.matchbox.id
    container_name = docker_container.matchbox.name
    static_ip      = var.static_ip
    network        = var.vlan_network_name
    api_endpoint   = "http://${var.static_ip}:${var.matchbox_port}"
    boot_url       = "http://${var.static_ip}:${var.matchbox_port}/boot.ipxe"
    image          = var.matchbox_image
    log_level      = var.log_level
  }
}