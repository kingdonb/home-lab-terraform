# METNOOM 13-net Environment Outputs

output "registry_caches" {
  description = "Summary of all registry cache endpoints"
  value = {
    docker_hub = module.docker_cache.cache_summary
    kubernetes = module.k8s_cache.cache_summary
    quay       = module.quay_cache.cache_summary
    google     = module.gcr_cache.cache_summary
    github     = module.ghcr_cache.cache_summary
  }
}

output "dns_dhcp_tftp" {
  description = "DNSmasq service configuration on 13-net"
  value = module.dnsmasq_13net.service_summary
}

output "pxe_boot_server" {
  description = "Matchbox PXE boot server configuration"
  value = module.matchbox_13net.service_summary
}

output "monitoring" {
  description = "Pi-hole monitoring configuration"
  value = module.pihole_exporter.service_summary
}

output "network_topology" {
  description = "Complete 13-net network service topology"
  value = {
    vlan_network = "vlan-13net"
    services = {
      dnsmasq = {
        ip = "10.17.13.252"
        services = ["DNS", "DHCP", "TFTP"]
        dhcp_range = "10.17.13.3-10.17.13.199"
      }
      matchbox = {
        ip = "10.17.13.251"
        services = ["PXE Boot", "API"]
        port = 8080
      }
      registry_caches = {
        services = ["Docker Registry Cache"]
        ports = [5050, 5051, 5052, 5053, 5054]
      }
    }
    dhcp_config = {
      range = "10.17.13.3-10.17.13.199"
      router = "10.17.13.249"
      dns_server = "10.17.13.254"
    }
  }
}

output "kubernetes_prerequisites" {
  description = "Summary of Kubernetes infrastructure dependencies"
  value = {
    dns_resolution = {
      status = "configured"
      server = "10.17.13.254"
      backup_server = "10.17.12.109" # Pi-hole on 12-net
    }
    registry_caching = {
      status = "configured"
      endpoints = [
        "localhost:5050", # docker.io
        "localhost:5051", # registry.k8s.io  
        "localhost:5052", # quay.io
        "localhost:5053", # gcr.io
        "localhost:5054"  # ghcr.io
      ]
    }
    pxe_infrastructure = {
      status = "configured"
      dnsmasq_ip = "10.17.13.252"
      matchbox_ip = "10.17.13.251"
      boot_url = "http://10.17.13.251:8080/boot.ipxe"
    }
    monitoring = {
      status = "configured"
      pihole_exporter = "localhost:9617"
      metrics_path = "/metrics"
    }
  }
}