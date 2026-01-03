# Next Steps and Actions Required

*Action Plan - January 3, 2026*

## Immediate Actions (Before TDG Setup)

### 1. Environment Setup
- [ ] **Switch to VSCode Insiders** (required for TDG skill)
- [ ] **Configure TDG skill** in VSCode Insiders environment
- [ ] **Test TDG functionality** with basic Terraform module search

### 2. Current Configuration Backup
- [ ] **Export Pi-hole configuration** from existing Docker Compose setup
  - Copy docker-compose.yml file
  - Export custom blocklists and whitelists  
  - Document current volume mounts and network configuration
  - Export Pi-hole admin settings and configurations

- [ ] **Document existing Docker run scripts** for other services:
  - Tailscale container script
  - Plex container script
  - Registry pull-through cache script
  - Matchbox container script
  - Any additional container scripts

## TDG Skill Research Tasks

### 3. Module Discovery with TDG
Once TDG is configured, use it to research:

#### Docker Management Modules
- [ ] **Search for Docker provider patterns**: `kreuzwerker/docker` usage examples
- [ ] **Find Synology-specific modules**: DSM integration approaches
- [ ] **Research remote Docker management**: SSH vs API vs local execution patterns

#### Pi-hole Specific Research
- [ ] **Pi-hole Terraform modules**: Existing community modules
- [ ] **DNS service patterns**: Multi-subnet DNS configuration
- [ ] **Container networking**: Cross-subnet service accessibility

#### Infrastructure Patterns
- [ ] **Static IP management**: Network interface configuration via Terraform
- [ ] **Service dependencies**: Container startup ordering and health checks
- [ ] **Configuration management**: Backup, restore, and drift detection

### 4. Technical Decisions Needed

#### Docker Management Approach
**Decision Point**: How will Terraform connect to Docker on Synology devices?
- **Option A**: SSH connection with docker commands
- **Option B**: Docker API over TCP
- **Option C**: Local execution from one of the Synology devices

**Research Required**: Test connectivity and authentication methods

#### State Storage Decision  
**Decision Point**: Where to store Terraform state?
- **Option A**: Local files (simple, single-user)
- **Option B**: Synology NFS share (reliable, local)
- **Option C**: Git repository (version controlled)
- **Option D**: Cloud backend (future-proof)

#### Network Architecture Decision
**Decision Point**: DHCP strategy affects DNS configuration
- **Option A**: Router-based DHCP per subnet (simpler)
- **Option B**: Centralized dnsmasq DHCP (more control)

**Impact**: Affects Pi-hole configuration and network service dependencies

## File and Configuration Collection

### 5. Configuration Export Tasks
- [ ] **Locate Pi-hole Docker Compose directory**
- [ ] **Export current Pi-hole admin configuration**:
  - Custom DNS settings
  - Blocklist URLs and local lists
  - Whitelist entries
  - Client group configurations
  - Query logging settings

- [ ] **Document current network binding**:
  - Which interfaces Pi-hole listens on
  - Current DNS forwarding configuration
  - Any custom dnsmasq configuration

- [ ] **Collect Docker run scripts** for migration reference

### 6. Network Documentation
- [ ] **Map current service dependencies**:
  - Which services depend on Pi-hole DNS
  - Which services are needed for Kubernetes cluster
  - Service startup order requirements

- [ ] **Document current IP assignments**:
  - Confirm DS1517+ IP address(es)
  - Document any other static assignments
  - Identify DHCP range usage

## Development Workflow Setup

### 7. Repository Workflow
- [ ] **Initialize Git repository** (if not already done)
- [ ] **Create development branch** for TDG experimentation
- [ ] **Setup .gitignore** for Terraform files (state, secrets, etc.)
- [ ] **Document commit and testing procedures**

### 8. Testing Environment
- [ ] **Plan parallel deployment strategy** for Pi-hole testing
- [ ] **Identify test DNS clients** for validation
- [ ] **Create rollback procedures** for failed deployments

## Expected Deliverables

### After TDG Research Phase
- [ ] **Module selection document** with recommended Terraform modules
- [ ] **Provider configuration** with authentication setup
- [ ] **Initial Pi-hole module** with current configuration preserved
- [ ] **Migration plan** with specific steps and timelines

### After Implementation Phase
- [ ] **Working Terraform configuration** for Pi-hole
- [ ] **Automated deployment scripts** for complete infrastructure
- [ ] **Documentation and runbooks** for operations
- [ ] **Backup and recovery procedures** for configurations

---

## Current Status Summary

**✅ Completed:**
- Repository structure created
- Infrastructure inventory documented  
- Implementation plan defined
- Research strategy outlined

**🔄 In Progress:**
- Switching to VSCode Insiders for TDG support

**⏳ Next Up:**
- TDG skill configuration and testing
- Current configuration backup and documentation
- Module research and selection with TDG

---

*This action plan will be updated as tasks are completed and new requirements emerge.*