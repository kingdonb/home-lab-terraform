output "container_id" {
  description = "ID of the pi-hole Docker container"
  value       = docker_container.pihole.id
}

output "container_name" {
  description = "Name of the pi-hole Docker container"
  value       = docker_container.pihole.name
}

output "network_id" {
  description = "ID of the Docker network"
  value       = docker_network.pihole_network.id
}

output "network_name" {
  description = "Name of the Docker network"
  value       = docker_network.pihole_network.name
}

output "dns_endpoint" {
  description = "DNS endpoint address"
  value       = "localhost:${var.dns_port}"
}

output "web_endpoint" {
  description = "Web interface endpoint"
  value       = "http://localhost:${var.web_port}/admin"
}

output "volumes" {
  description = "Created Docker volumes"
  value = {
    data    = docker_volume.pihole_data.name
    dnsmasq = docker_volume.pihole_dnsmasq.name
  }
}