# AI Coding Agent Instructions

This home lab infrastructure project uses **Test-Driven Development (TDG) methodology** with **OpenTofu/Terraform** for Infrastructure as Code. Focus on the established patterns and working solutions.

## Project Architecture

**Core Stack**: OpenTofu (Terraform alternative) + Docker + Go testing with Terratest
- **Infrastructure**: Multi-Pi-hole DNS with Docker containers on Synology NAS devices
- **Testing**: Go + Terratest for infrastructure validation, session-based API testing
- **TDG Methodology**: RED-GREEN-REFACTOR cycles for infrastructure development

### Key Components
- `terraform/modules/pihole/` - Docker-based Pi-hole deployment with v6+ JSON authentication
- `terraform/environments/test/` - Multi-instance setup with shared configuration volumes  
- `tests/` - Go test suite with Pi-hole v6+ API session management
- `tdg/` - TDG agent skills integration for test-driven infrastructure

## Critical Patterns & Working Solutions

### Pi-hole v6+ Authentication (WORKING)
```go
// Use JSON-based session authentication, NOT legacy API
authPayload := map[string]interface{}{
    "password": password,
    "totp":     nil,
}
// POST to /api/auth, capture 'sid' cookie
```
⚠️ **Never use legacy `/admin/api.php` - removed in Pi-hole v6+**

### Infrastructure Testing Patterns
```bash
# Follow this exact test sequence
go test ./tests/pihole_api_test.go -v     # API authentication 
go test ./tests/network_configuration_test.go -v  # DNS resolution
make test                                 # Full test suite
```

### TDG Development Workflow
1. **RED**: Write failing infrastructure tests (`red: test spec for pi-hole auth (#42)`)
2. **GREEN**: Implement minimal Terraform module (`green: implement pi-hole auth (#42)`) 
3. **REFACTOR**: Optimize configuration (`refactor: extract auth variables (#42)`)

### Terraform Module Structure
```hcl
# Always include these essentials for pi-hole modules:
- Linux capabilities: ["NET_ADMIN", "SYS_TIME", "SYS_NICE"]
- Environment format: FTLCONF_webserver_api_password (v6+ format)
- Health checks: dig command against localhost:53
- Shared volumes: For multi-instance configuration sync
```

## Developer Workflow Commands

### Testing & Validation
```bash
make test          # Run full test suite (unit + integration)
make test-unit     # Go tests only  
tofu validate      # Terraform syntax validation
tofu plan          # Preview infrastructure changes
```

### TDG Commands (Agent Skills)
```bash
/tdg:init          # Initialize TDG for new modules
/tdg:atomic-commit # Create clean, organized commits
```

### Docker Management
```bash
# Always use Docker Compose for integration tests
docker compose -f tests/pihole/docker-compose.test.yml up --abort-on-container-exit
docker compose down -v  # Clean up with volume removal
```

## Critical Infrastructure Dependencies

### Service Deployment Order
```
1. Pi-hole DNS servers (must be first - Kubernetes needs upstream DNS)
2. Registry pull-through cache containers
3. Kubernetes cluster with Cozystack (13-net)
4. Crossplane for advanced multi-tenant management
```
⚠️ **Pi-hole provides split-horizon DNS for Tailnet/cluster/local network resolution**

### Terraform vs Crossplane Scope
- **Terraform**: Foundation services (Pi-hole, registry, pre-Kubernetes infrastructure)
- **Crossplane**: Advanced management, multi-tenant configs (post-Kubernetes)
- **Principle**: Use the right tool - prefer Crossplane for problems it solves better

## Network & Infrastructure Specifics

### Home Lab Network Topology
- **Subnet 1**: 10.17.12.0/24 (DD-WRT Router)  
- **Subnet 2**: 10.17.13.0/24 (Mikrotik Hap-ax2)
- **Synology NAS**: Single-homed and dual-homed devices for long-term storage and production Pi-hole hosting
- **Test Networks**: 172.20.0.0/16, 172.21.0.0/16 (local Docker testing only, uses OrbStack on macOS)

### Port Allocation Pattern
```
Primary Pi-hole: 53 (DNS), 8080 (Web)
Secondary Pi-hole: 5353 (DNS), 8081 (Web)  
Test instances: 25353+ (DNS), 28080+ (Web)
```

## File Conventions & Patterns

### Test Files (`tests/`)
- `*_test.go` - Terratest-based infrastructure tests (local Docker via OrbStack)
- `pihole_api_test.go` - Pi-hole v6+ session authentication patterns
- `docker-compose.test.yml` - Integration test container setup (macOS only)

### Module Structure (`terraform/modules/`)
- `main.tf` - Resource definitions with dynamic blocks for networking
- `variables.tf` - Input parameters with validation
- `outputs.tf` - Module outputs for composition  
- `README.md` - Usage examples and requirements

### Environment Structure (`terraform/environments/`)
- Multi-instance Pi-hole deployment for primary/secondary DNS availability
- Shared configuration volumes for blocklist synchronization
- SSH-based deployment to single-homed and dual-homed Synology devices

## Integration Points

### External Dependencies
- **Docker**: kreuzwerker/docker provider v3.0+
- **Terratest**: gruntwork-io/terratest for Go-based testing
- **DNS**: miekg/dns library for DNS validation
- **Synology**: SSH-based container management on DSM 7.2+

### Authentication Management
- Pi-hole v6+: JSON session-based with cookie capture
- Synology: SSH key-based access (avoid password hardcoding)
- Container management: Docker socket or SSH-based deployment

## Common Gotchas

❌ **Don't use** legacy Pi-hole API endpoints (`/admin/api.php`)  
❌ **Don't hardcode** credentials in Terraform files  
❌ **Don't skip** health checks for DNS containers  
✅ **Do use** session-based authentication for Pi-hole v6+  
✅ **Do validate** DNS resolution in tests before marking green  
✅ **Do follow** TDG RED-GREEN-REFACTOR cycles for infrastructure changes