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

output "pihole_config_summary" {
  description = "Pi-hole configuration summary"
  value       = module.pihole_config.pihole_config_summary
}

output "dns_records_created" {
  description = "DNS records created in Pi-hole"
  value       = module.pihole_config.dns_records_created
}

output "cname_records_created" {
  description = "CNAME records created in Pi-hole"
  value       = module.pihole_config.cname_records_created
}