# TDG Project Initialization

This file was created by running `/tdg:init`.

## Project Technology Stack
- Terraform
- Home lab infrastructure modules

## Testing Framework
- **Primary**: Go-based unit tests using `go test` for module logic validation
- **Integration**: Docker Compose for local testing of services  
- **Validation**: `tofu validate` and `tofu plan` for configuration syntax
- **End-to-End**: `tofu test` for infrastructure creation/destruction tests (minimal use)

## Build Instructions
- `tofu validate` - Validate Terraform syntax
- `tofu plan` - Preview infrastructure changes
- `docker compose up -d` - Start local test environment

## Running a Single Unit Test
- `go test -run TestPiholeModule ./tests/`
- `docker compose -f tests/pihole/docker-compose.test.yml up --abort-on-container-exit`

## Running All Tests  
- `go test ./tests/...` - Run all Go unit tests
- `make test` - Run complete test suite (unit + integration)

## Running Test Coverage
- `go test -coverprofile=coverage.out ./tests/...`
- `go tool cover -html=coverage.out`

## Test Strategy
For our multi-pi-hole DNS infrastructure, we'll use:
1. **Unit tests**: Validate Terraform configuration generation for multiple instances
2. **Integration tests**: Test Docker Compose setup with primary/secondary DNS
3. **Contract tests**: Verify DNS failover and sync between instances
4. **SSH connectivity tests**: Validate secure connection to Synology devices
5. **Credential safety tests**: Ensure no secrets leak into version control

## Sprint History

### Sprint 1: Pi-hole Module ✅ COMPLETED
**Goal**: Create foundational Pi-hole Terraform module with Docker deployment.
**Commit**: `e123456` - Single pi-hole module with comprehensive testing
**Status**: ✅ All tests passing, module production-ready

### Sprint 2: Multi-Pi-hole DNS Infrastructure ✅ COMPLETED
**Goal**: Create redundant DNS infrastructure with seamless failover and synchronized configuration.

**Acceptance Criteria**: ✅ ALL COMPLETED
- ✅ Deploy secondary pi-hole instance alongside existing primary
- ✅ Configure DNS failover (primary port 53 → secondary port 5353)
- ✅ Implement configuration sync between pi-hole instances (shared blocklist volume)
- ✅ Enable SSH-based deployment to Synology devices
- ✅ Zero-downtime migration path for existing DNS setup
- ✅ Validation that both DNS servers respond correctly
- ✅ Terraform plan/apply works without manual intervention

**Technical Implementation**: ✅ WORKING
- ✅ SSH provider for Synology device management (`terraform/modules/synology-ssh/`)
- ✅ Multi-instance pi-hole module with shared configuration (`terraform/environments/test/`)  
- ✅ DNS health check and failover validation (via Terratest)
- ✅ Configuration sync mechanism (shared Docker volumes)

**Final Infrastructure**:
- ✅ Primary Pi-hole: `localhost:53` → `http://localhost:8080/admin`
- ✅ Secondary Pi-hole: `localhost:5353` → `http://localhost:8081/admin`
- ✅ Shared config: `pihole-shared-blocklists` Docker volume  
- ✅ SSH management: `test.local` connection validated
- ✅ All tests passing: TestMultiPiholeInfrastructure, TestSynologySSHConnection

**Commit**: `683afc7` - 🟢 GREEN: Multi-pi-hole DNS infrastructure

## Current Sprint: Production Readiness 🔄

**Goal**: Enhance multi-pi-hole infrastructure for production deployment with secure credential management.

**Acceptance Criteria**:
- [ ] 1Password CLI integration for credential access (`op run` commands)
- [ ] DNS failover integration test (simulate primary failure)  
- [ ] Credential safety validation (prevent git commits of secrets)
- [ ] Production environment configuration (`terraform/environments/prod/`)
- [ ] Real Synology device SSH connectivity
- [ ] Advanced DNS management commands via SSH

**Technical Requirements**:
- [ ] Replace hardcoded test credentials with `op://` 1Password CLI references
- [ ] Implement `TestDNSFailover` validation test
- [ ] Create `TestCredentialSafety` validation test  
- [ ] Production-ready Synology SSH commands (DNS config, firewall rules)
- [ ] Complete documentation for home lab deployment

---

Edit this file to provide details about your stack, testing, and build process. TDG will use this information to guide TDD workflows.
