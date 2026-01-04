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
For our pi-hole module, we'll use:
1. **Unit tests**: Validate Terraform configuration generation
2. **Integration tests**: Test Docker Compose setup locally
3. **Contract tests**: Verify API responses from pi-hole service

---

Edit this file to provide details about your stack, testing, and build process. TDG will use this information to guide TDD workflows.
