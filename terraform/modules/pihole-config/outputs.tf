output "dns_records_created" {
  description = "DNS records created for homelab services"
  value = {
    gateway  = pihole_dns_record.homelab_gateway.domain
    nas      = pihole_dns_record.homelab_nas.domain
    registry = pihole_dns_record.homelab_registry.domain
  }
}

output "cname_records_created" {
  description = "CNAME records created for service aliases"
  value = {
    docker_registry     = pihole_cname_record.docker_registry.domain
    container_registry  = pihole_cname_record.container_registry.domain
  }
}

output "pihole_config_summary" {
  description = "Summary of Pi-hole configuration created"
  value = {
    total_dns_records   = length([pihole_dns_record.homelab_gateway, pihole_dns_record.homelab_nas, pihole_dns_record.homelab_registry])
    total_cname_records = length([pihole_cname_record.docker_registry, pihole_cname_record.container_registry])
    pihole_url         = var.pihole_base_url
  }
}