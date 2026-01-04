## Makefile for Home Lab Terraform Infrastructure

.PHONY: init validate plan test test-unit test-integration clean

# Initialize Terraform
init:
	tofu init

# Validate Terraform configuration
validate:
	tofu validate

# Plan infrastructure changes
plan:
	tofu plan

# Run all tests
test: test-unit test-integration

# Run unit tests
test-unit:
	go test ./tests/... -v

# Run integration tests  
test-integration:
	docker compose -f tests/pihole/docker-compose.test.yml up --abort-on-container-exit

# Clean up test artifacts
clean:
	docker compose -f tests/pihole/docker-compose.test.yml down -v
	go clean -testcache