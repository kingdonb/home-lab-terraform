# METNOOM 13-net Infrastructure Environment
# Replicates the complete container setup on Synology single-homed in 13-net

# Registry pull-through caches (ports 5050-5054)
module "docker_cache" {
  source = "../../modules/registry-cache"
  registry_name = "docker.io"
  cache_port    = 5050
}

module "k8s_cache" {
  source = "../../modules/registry-cache"
  registry_name = "registry.k8s.io"
  cache_port    = 5051
}

module "quay_cache" {
  source = "../../modules/registry-cache"
  registry_name = "quay.io"
  cache_port    = 5052
}

module "gcr_cache" {
  source = "../../modules/registry-cache"
  registry_name = "gcr.io"
  cache_port    = 5053
}

module "ghcr_cache" {
  source = "../../modules/registry-cache"
  registry_name = "ghcr.io"
  cache_port    = 5054
}

# DNSmasq for DHCP/DNS/TFTP on 13-net VLAN
module "dnsmasq_13net" {
  source = "../../modules/dnsmasq"
  
  container_name    = "dnsmasq"
  vlan_network_name = "vlan-13net"
  static_ip         = "10.17.13.252"
  
  # DHCP Configuration
  dhcp_enabled      = true
  dhcp_range_start  = "10.17.13.3"
  dhcp_range_end    = "10.17.13.199"
  dhcp_router       = "10.17.13.249"
  dhcp_dns_server   = "10.17.13.254"
  
  # TFTP and PXE Configuration
  tftp_enabled      = true
  pxe_enabled       = true
  matchbox_server   = "10.17.13.251"
  matchbox_port     = 8080
  
  # Logging
  enable_logging    = true
}

# Matchbox for PXE boot services on 13-net VLAN
module "matchbox_13net" {
  source = "../../modules/matchbox"
  
  container_name    = "matchbox"
  matchbox_image    = "kingdonb/matchbox:v1.10.5-cozy-spin-tailscale"
  vlan_network_name = "vlan-13net"
  static_ip         = "10.17.13.251"
  matchbox_port     = 8080
  log_level         = "debug"
}

# Pi-hole Exporter for monitoring (from docker-compose.yml config)
module "pihole_exporter" {
  source = "../../modules/pihole-exporter"
  
  container_name    = "pihole-exporter"
  pihole_hostname   = "10.17.12.109"  # Pi-hole ServerIP from docker-compose
  pihole_api_token  = var.pihole_api_token
  exporter_port     = 9617
  scrape_interval   = "10s"
}