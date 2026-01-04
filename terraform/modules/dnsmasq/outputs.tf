# DNSmasq Module Outputs

output "container_id" {
  description = "Docker container ID of the DNSmasq server"
  value       = docker_container.dnsmasq.id
}

output "container_name" {
  description = "Name of the DNSmasq container"
  value       = docker_container.dnsmasq.name
}

output "static_ip" {
  description = "Static IP address of the DNSmasq server"
  value       = var.static_ip
}

output "dhcp_range" {
  description = "DHCP IP address range"
  value       = var.dhcp_enabled ? "${var.dhcp_range_start}-${var.dhcp_range_end}" : null
}

output "dhcp_config" {
  description = "DHCP configuration summary"
  value = var.dhcp_enabled ? {
    range      = "${var.dhcp_range_start}-${var.dhcp_range_end}"
    router     = var.dhcp_router
    dns_server = var.dhcp_dns_server
  } : null
}

output "tftp_config" {
  description = "TFTP configuration summary"
  value = var.tftp_enabled ? {
    enabled     = true
    volume_name = var.tftp_volume_name
    root_path   = "/var/lib/tftpboot"
  } : null
}

output "pxe_config" {
  description = "PXE boot configuration summary"
  value = var.pxe_enabled ? {
    enabled        = true
    matchbox_server = var.matchbox_server
    matchbox_port   = var.matchbox_port
    boot_url       = "http://${var.matchbox_server}:${var.matchbox_port}/boot.ipxe"
  } : null
}

output "network_config" {
  description = "Network configuration summary"
  value = {
    vlan_network = var.vlan_network_name
    static_ip    = var.static_ip
    services = {
      dhcp = var.dhcp_enabled
      tftp = var.tftp_enabled  
      pxe  = var.pxe_enabled
    }
  }
}

output "service_summary" {
  description = "Complete DNSmasq service configuration"
  value = {
    container_id   = docker_container.dnsmasq.id
    container_name = docker_container.dnsmasq.name
    static_ip      = var.static_ip
    network        = var.vlan_network_name
    services = {
      dhcp_enabled = var.dhcp_enabled
      tftp_enabled = var.tftp_enabled
      pxe_enabled  = var.pxe_enabled
      logging_enabled = var.enable_logging
    }
    dhcp_range = var.dhcp_enabled ? "${var.dhcp_range_start}-${var.dhcp_range_end}" : null
    matchbox_url = var.pxe_enabled ? "http://${var.matchbox_server}:${var.matchbox_port}" : null
  }
}