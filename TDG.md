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

## Current Sprint: Multi-Pi-hole DNS Infrastructure
**Goal**: Create redundant DNS infrastructure with seamless failover and synchronized configuration.

**Acceptance Criteria**:
- [ ] Deploy secondary pi-hole instance alongside existing primary
- [ ] Configure DNS failover (primary → secondary)
- [ ] Implement configuration sync between pi-hole instances
- [ ] Enable SSH-based deployment to Synology devices
- [ ] Secure credential management via 1Password CLI
- [ ] Zero-downtime migration path for existing DNS setup
- [ ] Validation that both DNS servers respond correctly
- [ ] Terraform plan/apply works without manual intervention

**Technical Requirements**:
- SSH provider for Synology device management
- 1Password CLI integration for credential access
- Multi-instance pi-hole module with shared configuration
- DNS health check and failover validation
- Configuration sync mechanism (shared volumes or API sync)

---

Edit this file to provide details about your stack, testing, and build process. TDG will use this information to guide TDD workflows.
