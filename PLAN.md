# Home Lab Terraform Implementation Plan

*Status: Planning Phase*  
*Last Updated: January 3, 2026*

## Phase 1: Foundation Setup 🔧

### Repository Structure
- [x] Initialize Git repository
- [x] Create README.md with infrastructure overview
- [x] Create PLAN.md (this document)
- [ ] Setup Terraform directory structure
- [ ] Configure Agent Skills integration

### Terraform Foundation
- [ ] Define provider requirements
- [ ] Configure state backend
- [ ] Setup environment structure (dev/staging/prod if needed)
- [ ] Create base modules directory structure

### Agent Skills Integration
- [ ] Configure TDG (Terraform Documentation Generator) skill
- [ ] Research and document module selection workflow
- [ ] Setup automated documentation generation

## Phase 2: Docker Container Management 🐳

### Research & Planning
- [ ] Evaluate Terraform Docker providers:
  - [ ] `kreuzwerker/docker` provider
  - [ ] `terraform-docker-modules` options
  - [ ] Direct SSH/API integration approaches
- [ ] Document current Docker Compose configurations
- [ ] Plan migration strategy from manual to IaC

### Pi-hole Priority Implementation
- [ ] **Current State Assessment**:
  - [ ] Document existing Pi-hole version and configuration
  - [ ] Export current blocklists and custom settings
  - [ ] Document network interface binding requirements
- [ ] **New Pi-hole Deployment**:
  - [ ] Define target Pi-hole version
  - [ ] Create Terraform module for Pi-hole container
  - [ ] Configure multi-subnet DNS serving
  - [ ] Implement configuration backup/restore
- [ ] **Migration & Testing**:
  - [ ] Test new deployment in parallel
  - [ ] Validate DNS resolution across both subnets
  - [ ] Plan cutover procedure

## Phase 3: Network Interface Management 🌐

### Synology Interface Configuration
- [ ] Research Synology DSM API capabilities
- [ ] Define static IP configuration management
- [ ] Implement interface state validation
- [ ] Create post-reboot verification scripts

### Multi-homed Device Support
- [ ] DS923+ dual-interface configuration (10.17.12.109 & 10.17.13.204)
- [ ] Pi-hole Synology interface configuration
- [ ] Cross-subnet service accessibility

## Phase 4: Service Expansion 📈

### Additional Container Services
- [ ] Tailscale container management
- [ ] Plex Media Server configuration
- [ ] Remaining ~4 container services
- [ ] Service dependency mapping

### Monitoring & Health Checks
- [ ] Service availability monitoring
- [ ] Network connectivity validation
- [ ] Automated alerting setup

## Phase 5: Automation & Orchestration 🤖

### Bootstrap Process
- [ ] Complete infrastructure deployment from scratch
- [ ] Post-reboot recovery procedures
- [ ] Configuration drift detection
- [ ] Backup and disaster recovery

### CI/CD Integration
- [ ] Automated testing of infrastructure changes
- [ ] Validation pipelines
- [ ] Documentation generation automation

## Key Questions to Resolve 🤔

### Pi-hole Specific
- [ ] Current Pi-hole version and target upgrade version
- [ ] Existing custom blocklists and configurations
- [ ] DNS forwarding and caching requirements

### Infrastructure Management
- [ ] Terraform state storage preference (local/cloud)
- [ ] Authentication method for Synology management
- [ ] Backup strategy for configurations

### Docker Strategy
- [ ] Keep Docker Compose + Terraform orchestration, or migrate to pure Terraform?
- [ ] Container update and rollback procedures
- [ ] Volume and data persistence management

## Success Criteria ✅

### Phase 1 Success
- [ ] Repository properly structured
- [ ] Agent Skills operational
- [ ] Development workflow established

### Phase 2 Success
- [ ] Pi-hole deployed via Terraform
- [ ] DNS serving correctly on both subnets
- [ ] No manual post-reboot intervention required

### Final Success
- [ ] Complete infrastructure deployable via `terraform apply`
- [ ] All services automatically available after reboot
- [ ] Configuration changes managed through code
- [ ] Clear documentation and runbooks

---

## Next Actions 🎯

1. **Immediate**: Configure Agent Skills and TDG integration
2. **Next**: Research Docker Terraform providers and choose approach
3. **Then**: Begin Pi-hole configuration documentation and migration planning

*This plan will be updated as we progress through each phase.*