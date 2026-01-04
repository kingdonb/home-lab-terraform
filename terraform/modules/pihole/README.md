# Pi-hole Terraform Module

This module provides a Docker-based Pi-hole deployment for DNS filtering and ad blocking.

## Features

- **Container Management**: Creates and manages a Pi-hole Docker container
- **Network Isolation**: Dedicated Docker network with configurable subnet
- **Persistent Storage**: Volumes for configuration and DNS data persistence
- **Health Monitoring**: Built-in health checks for service availability
- **Configurable Ports**: Customizable DNS and web interface ports
- **Environment Configuration**: Timezone, passwords, and upstream DNS settings

## Usage

```hcl
module "pihole" {
  source = "./modules/pihole"
  
  container_name = "home-pihole"
  network_name   = "home-dns-net"
  dns_port      = 53
  web_port      = 8080
  timezone      = "America/New_York"
  web_password  = var.pihole_password
  upstream_dns  = "1.1.1.1;1.0.0.1"
}
```

## Outputs

- `dns_endpoint` - DNS server endpoint (e.g., localhost:53)
- `web_endpoint` - Web interface URL (e.g., http://localhost:8080/admin)
- `container_id` - Docker container ID
- `network_id` - Docker network ID
- `volumes` - Created volume names for backup/restore operations

## Testing

This module includes comprehensive tests:

```bash
# Run all tests
make test

# Run unit tests only
make test-unit

# Run integration tests
make test-integration
```

## Requirements

| Name | Version |
|------|---------|
| terraform | >= 1.0 |
| docker | ~> 3.0 |

## Providers

| Name | Version |
|------|---------|
| docker | ~> 3.0 |

## Resources

| Name | Type |
|------|------|
| docker_container.pihole | resource |
| docker_image.pihole | resource |
| docker_network.pihole_network | resource |
| docker_volume.pihole_data | resource |
| docker_volume.pihole_dnsmasq | resource |