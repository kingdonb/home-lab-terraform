# Home Lab Terraform Infrastructure

This repository manages the Infrastructure as Code (IaC) configuration for a home lab environment using Terraform.

## Infrastructure Overview

### Network Topology
- **Subnet 1**: 10.17.12.0/24 (DD-WRT Router Gateway)
- **Subnet 2**: 10.17.13.0/24 (Mikrotik Hap-ax2 Router Gateway)
- **Router Interconnect**: Mikrotik at 10.17.12.249 ↔ 10.17.13.249

### Hardware Inventory

#### Synology NAS Devices
- **DS1517+** (Intel Atom C2538) - DSM 7.2.1-69057
- **DS923+** (AMD Ryzen R1600) - DSM 7.2.2-72806
  - Interfaces: 10.17.12.109, 10.17.13.204
  - Services: Tailscale, Plex

#### Network Equipment
- DD-WRT Router (10.17.12.0/24 gateway)
- Mikrotik Hap-ax2 Router (10.17.13.0/24 gateway)

### Current Services
- **Pi-hole DNS** (Docker container on 13-net Synology)
- **Tailscale** (Docker container on DS923+)
- **Plex Media Server** (Docker container on DS923+)
- **Additional containers** (~7 total across both Synology devices)

## Goals

1. **Centralize Configuration**: Move from manual setup to Terraform-managed IaC
2. **Improve Reliability**: Eliminate manual post-reboot configuration steps
3. **Modernize Pi-hole**: Upgrade from current outdated version
4. **Bootstrap Process**: Define clear deployment and recovery procedures
5. **Documentation**: Maintain clear infrastructure documentation

## Repository Structure

```
├── README.md              # This file
├── PLAN.md               # Implementation roadmap
├── terraform/            # Terraform configurations
│   ├── environments/     # Environment-specific configs
│   ├── modules/          # Reusable Terraform modules
│   └── providers/        # Provider configurations
├── docker-compose/       # Docker Compose files
├── configs/              # Service configuration files
└── scripts/              # Automation and utility scripts
```

## Quick Start

> **Note**: This repository is under active development. See [PLAN.md](PLAN.md) for current status and roadmap.

## Prerequisites

- Terraform >= 1.6
- Docker access to Synology devices
- Network access to home lab subnets

## Agent Skills Integration

This repository leverages Agent Skills for enhanced Terraform development:
- **TDG (Terraform Documentation Generator)** - Module discovery and documentation
- Additional skills to be configured as needed

---

*Last updated: January 3, 2026*