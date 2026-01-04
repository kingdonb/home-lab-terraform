output "primary_dns_endpoint" {
  description = "Primary DNS server endpoint"
  value       = module.primary_pihole.dns_endpoint
}

output "secondary_dns_endpoint" {
  description = "Secondary DNS server endpoint"  
  value       = module.secondary_pihole.dns_endpoint
}

output "primary_web_interface" {
  description = "Primary Pi-hole web interface"
  value       = module.primary_pihole.web_endpoint
}

output "secondary_web_interface" {
  description = "Secondary Pi-hole web interface"
  value       = module.secondary_pihole.web_endpoint
}

output "shared_config_volume" {
  description = "Shared configuration volume name"
  value       = docker_volume.pihole_shared_config.name
}

output "synology_connection_status" {
  description = "Synology SSH connection status"
  value       = length(module.synology_connection) > 0 ? module.synology_connection[0].connection_status : "Not configured"
}