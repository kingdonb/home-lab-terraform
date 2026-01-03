# Terraform Module Research & Strategy

*Research Planning Document - January 3, 2026*

## Primary Research Goals with TDG

### 1. Docker Container Management
**Target Providers:**
- `kreuzwerker/docker` - Primary Docker provider for Terraform
- `terraform-docker-modules/*` - Community modules collection
- Custom SSH-based solutions for Synology integration

**Key Research Questions:**
- How to manage Docker containers on remote Synology hosts?
- Best practices for container dependencies and startup order?
- Health checking and restart policies in Terraform?
- Volume and network management strategies?

### 2. Pi-hole Specific Modules
**Research Targets:**
- Existing Pi-hole Terraform modules
- DNS server configuration modules
- Multi-interface/multi-subnet DNS serving
- Pi-hole configuration backup/restore automation

**Critical Requirements:**
- Must serve DNS requests from both 10.17.12.0/24 and 10.17.13.0/24
- Must preserve custom blocklists and configurations
- Must handle graceful upgrades without DNS outage
- Must auto-start after host reboot

### 3. Synology Integration
**Investigation Areas:**
- DSM API access and authentication methods
- Network interface configuration via Terraform
- Docker daemon management on Synology
- Static IP configuration persistence

**Platform Specifics:**
- DSM 7.2.1-69057 (DS1517+) capabilities
- DSM 7.2.2-72806 (DS923+) capabilities
- Cross-platform compatibility considerations

### 4. Network Infrastructure Management
**Modules to Research:**
- Static IP configuration management
- Multi-homed device configuration
- Network interface validation and health checks
- Cross-subnet service accessibility

## Current Service Migration Strategy

### Phase 1: Pi-hole (Critical Priority)
**Current State:** Docker Compose on DS923+
**Target State:** Terraform-managed Docker container
**Migration Approach:**
1. Export current configuration and custom lists
2. Build parallel Terraform deployment
3. Test DNS resolution on both subnets
4. Cutover with minimal downtime

**TDG Research Focus:**
- Pi-hole container configuration patterns
- DNS service health validation
- Multi-subnet binding configuration

### Phase 2: Infrastructure Services
**Services to Migrate:**
- Tailscale (networking critical)
- Registry (pull-through cache)
- Matchbox (netboot services)

**TDG Research Focus:**
- Container dependency management
- Network service ordering
- Infrastructure service patterns

### Phase 3: Media & Optional Services
**Services to Migrate:**
- Plex Media Server
- Dnsmasq (if needed)
- Additional containers

## Terraform Architecture Decisions

### State Management Research
**Options to Evaluate:**
- Local state files (simplest for home lab)
- Remote state on Synology device (reliable local storage)
- Cloud backends (overkill but future-proof)

### Provider Strategy
**Primary Approach:** Use `kreuzwerker/docker` with SSH or TCP connection
**Fallback Approach:** API-based management or local execution

### Module Organization
```
terraform/
├── modules/
│   ├── pihole/           # Pi-hole specific module
│   ├── synology-docker/  # Synology Docker management
│   ├── network-service/  # Generic network service module
│   └── infrastructure/   # Base infrastructure module
├── environments/
│   └── homelab/         # Single environment for now
└── configs/             # Service-specific configurations
```

## Research Methodology with TDG

### 1. Module Discovery Phase
- Use TDG to search for relevant Docker management modules
- Identify Pi-hole specific Terraform implementations
- Research Synology integration approaches

### 2. Evaluation Phase  
- Test promising modules in isolated environment
- Document compatibility with Synology DSM
- Validate multi-subnet network requirements

### 3. Integration Phase
- Adapt selected modules for home lab requirements
- Create custom modules where needed
- Document configuration patterns and best practices

## Critical Success Factors

### Must-Have Capabilities
- [ ] Automated Pi-hole deployment and configuration
- [ ] Cross-subnet DNS serving without manual configuration
- [ ] Container restart after host reboot
- [ ] Configuration preservation during upgrades

### Nice-to-Have Features
- [ ] Container health monitoring and alerting
- [ ] Configuration drift detection
- [ ] Automated backup and restore
- [ ] Service dependency management

---

*This research plan will guide the TDG skill usage and module selection process.*