# Registry Pull-Through Cache Module

Provides Docker registry pull-through caching for multiple upstream registries.

## Overview

This module deploys a registry container configured as a pull-through cache for a specified upstream registry. It helps reduce bandwidth usage and improves pull performance for container images.

## Supported Registries

- `docker.io` - Docker Hub (default registry)
- `registry.k8s.io` - Kubernetes container registry  
- `quay.io` - Red Hat Quay registry
- `gcr.io` - Google Container Registry
- `ghcr.io` - GitHub Container Registry

## Usage

```hcl
module "docker_hub_cache" {
  source = "./modules/registry-cache"
  
  registry_name = "docker.io"
  cache_port    = 5050
  
  # Optional: custom upstream configuration
  upstream_url  = "https://registry-1.docker.io"
  
  # Network and storage
  network_name = "registry-net"
  volume_size  = "100GB"
  
  tags = {
    Purpose = "Container registry caching"
    Registry = "docker.io"
  }
}
```

## Multiple Registry Example

```hcl
# Docker Hub cache
module "docker_cache" {
  source = "./modules/registry-cache"
  registry_name = "docker.io"
  cache_port = 5050
}

# Kubernetes registry cache  
module "k8s_cache" {
  source = "./modules/registry-cache"
  registry_name = "registry.k8s.io"
  cache_port = 5051
}

# Quay cache
module "quay_cache" {
  source = "./modules/registry-cache" 
  registry_name = "quay.io"
  cache_port = 5052
}

# Google Container Registry cache
module "gcr_cache" {
  source = "./modules/registry-cache"
  registry_name = "gcr.io"
  cache_port = 5053
}

# GitHub Container Registry cache
module "ghcr_cache" {
  source = "./modules/registry-cache"
  registry_name = "ghcr.io" 
  cache_port = 5054
}
```

## Configuration

The module automatically generates registry configuration based on the upstream registry specified. For custom configurations, see the variables.tf for additional options.

## Port Allocation

Default port assignments match your current METNOOM setup:
- Docker Hub: 5050
- Kubernetes: 5051  
- Quay: 5052
- GCR: 5053
- GHCR: 5054

## Storage

Each registry cache uses a dedicated Docker volume for persistence. Cache data survives container restarts and updates.