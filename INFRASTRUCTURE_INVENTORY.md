# Infrastructure Inventory & Current State

*Captured: January 3, 2026*

## Network Configuration

### Subnets
- **10.17.12.0/24** - Primary subnet with DD-WRT router as gateway
- **10.17.13.0/24** - Secondary subnet with Mikrotik Hap-ax2 as gateway

### Router Interconnection
- **Mikrotik Hap-ax2**: 10.17.12.249 ↔ 10.17.13.249 (multi-interface)

## Hardware Inventory

### Synology NAS Devices

#### DS1517+ (Intel Atom C2538)
- **DSM Version**: 7.2.1-69057
- **Network**: TBD - interface configuration needed
- **Services**: TBD
- **Notes**: Older model, but solid for infrastructure services

#### DS923+ (AMD Ryzen R1600) 
- **DSM Version**: 7.2.2-72806
- **Network Interfaces**: 
  - 10.17.12.109 (subnet 1)
  - 10.17.13.204 (subnet 2)
- **Current Services**:
  - **Pi-hole** (Docker Compose) - 13-net only
  - **Tailscale** (Docker script)
  - **Plex Media Server** (Docker script)
- **IP Assignment**: Static (infrastructure requirement, may run DHCP)
- **Notes**: Primary multi-homed server

## Current Service Deployment Methods

### Docker Compose Services
- **Pi-hole**: Currently managed via Docker Compose
  - Location: Has existing compose file and configuration tree
  - Custom blocklists and configurations in place
  - **CRITICAL**: Must serve DNS to both subnets (current issue)
  - Target: Upgrade to latest Pi-hole version

### Docker Script-Managed Services (Target for Terraform Migration)
- **Tailscale container**
- **Plex Media Server**  
- **Registry** (pull-through caching)
- **Matchbox** (network boot services)
- **Dnsmasq** (sometimes used, from Cozystack/Talos guide)
- **Additional containers** (~2 more, total ~7 containers)

### Infrastructure Services Context
- **Matchbox**: Used for network boot (likely PXE boot for Talos Linux)
- **Registry**: Pull-through cache (reduces external Docker pulls)
- **Dnsmasq**: Alternative DHCP option vs router-based DHCP

## Network Boot & Kubernetes Context
- Uses Cozystack methodology
- Talos Linux netboot capability
- Indicates sophisticated home lab with K8s infrastructure
- Pi-hole critical for proper DNS resolution for K8s

## Current Pain Points

### Pi-hole Issues
- **Outdated version** running
- **Misconfigured** - doesn't properly serve both subnets
- **Manual intervention required** after reboots
- **Manual interface configuration** in Pi-hole UI needed

### General Infrastructure Issues  
- **Manual post-reboot verification** required
- **Fragile configuration** - hand-cobbled over time
- **No centralized configuration management**
- **Interface configuration** not persistent/automated

## IP Address Assignments

### Static Assignments (Infrastructure)
- **DS923+**: 10.17.12.109 & 10.17.13.254
- **Mikrotik**: 10.17.12.249 & 10.17.13.249
- **DS1517+**: TBD
- **Other services**: TBD

### DHCP Strategy (Undecided)
- Option 1: Router-based DHCP per subnet
- Option 2: Centralized DHCP via dnsmasq
- Consideration: Infrastructure devices need static assignments regardless

---

*This inventory will be updated as more details are discovered during implementation.*