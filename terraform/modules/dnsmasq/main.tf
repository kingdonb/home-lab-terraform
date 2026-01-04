# DNSmasq Module
# Based on start_dnsmasq script configuration for 13-net VLAN

terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 3.0"
    }
  }
}

# DNSmasq container for DHCP/DNS/TFTP services
resource "docker_container" "dnsmasq" {
  name  = var.container_name
  image = var.dnsmasq_image
  
  restart = "unless-stopped"
  
  # Network configuration - requires external VLAN network
  networks_advanced {
    name = var.vlan_network_name
    ipv4_address = var.static_ip
  }
  
  # Required capabilities for networking
  capabilities {
    add = ["NET_ADMIN"]
  }
  
  # DNSmasq configuration arguments
  command = concat([
    "-d",           # Run in foreground for Docker
    "-q",           # Quiet mode
    "-p0"           # No DNS port binding (handled by network)
  ], var.dhcp_enabled ? [
    "--dhcp-range=${var.dhcp_range_start},${var.dhcp_range_end}",
    "--dhcp-option=option:router,${var.dhcp_router}",
    "--dhcp-option=option:dns-server,${var.dhcp_dns_server}"
  ] : [], var.tftp_enabled ? [
    "--enable-tftp",
    "--tftp-root=/var/lib/tftpboot"
  ] : [], var.pxe_enabled ? [
    # PXE Boot configuration
    "--dhcp-match=set:bios,option:client-arch,0",
    "--dhcp-boot=tag:bios,undionly.kpxe",
    "--dhcp-match=set:efi32,option:client-arch,6", 
    "--dhcp-boot=tag:efi32,ipxe.efi",
    "--dhcp-match=set:efibc,option:client-arch,7",
    "--dhcp-boot=tag:efibc,ipxe.efi",
    "--dhcp-match=set:efi64,option:client-arch,9",
    "--dhcp-boot=tag:efi64,ipxe.efi",
    "--dhcp-userclass=set:ipxe,iPXE",
    "--dhcp-boot=tag:ipxe,http://${var.matchbox_server}:${var.matchbox_port}/boot.ipxe"
  ] : [], var.enable_logging ? [
    "--log-queries",
    "--log-dhcp"
  ] : [], var.additional_args)\n  \n  # TFTP volume for PXE boot files\n  dynamic \"volumes\" {\n    for_each = var.tftp_enabled ? [1] : []\n    content {\n      volume_name    = var.tftp_volume_name\n      container_path = \"/var/lib/tftpboot\"\n      read_only      = false\n    }\n  }\n  \n  # Additional volumes\n  dynamic \"volumes\" {\n    for_each = var.additional_volumes\n    content {\n      host_path      = volumes.value.host_path\n      container_path = volumes.value.container_path\n      read_only      = try(volumes.value.read_only, false)\n    }\n  }\n  \n  # Logging configuration\n  log_driver = \"json-file\"\n  log_opts = {\n    max-size = \"10m\"\n    max-file = \"3\"\n  }\n  \n  labels {\n    label = \"purpose\"\n    value = \"dns-dhcp-tftp\"\n  }\n  \n  labels {\n    label = \"network\"\n    value = var.vlan_network_name\n  }\n  \n  labels {\n    label = \"managed_by\"\n    value = \"terraform\"\n  }\n}