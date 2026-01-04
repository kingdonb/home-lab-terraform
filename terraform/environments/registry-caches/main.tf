# Multi-Registry Cache Environment
# Creates all registry caches matching current METNOOM setup

# Docker Hub cache (port 5050)
module "docker_cache" {
  source = "../../modules/registry-cache"
  
  registry_name = "docker.io"
  cache_port    = 5050
  
  tags = {
    Environment = "production"
    Purpose     = "container-registry-cache"
    Registry    = "docker-hub"
  }
}

# Kubernetes registry cache (port 5051)
module "k8s_cache" {
  source = "../../modules/registry-cache"
  
  registry_name = "registry.k8s.io"
  cache_port    = 5051
  
  tags = {
    Environment = "production"
    Purpose     = "container-registry-cache"
    Registry    = "kubernetes"
  }
}

# Quay cache (port 5052)
module "quay_cache" {
  source = "../../modules/registry-cache"
  
  registry_name = "quay.io"
  cache_port    = 5052
  
  tags = {
    Environment = "production"
    Purpose     = "container-registry-cache"
    Registry    = "quay"
  }
}

# Google Container Registry cache (port 5053)
module "gcr_cache" {
  source = "../../modules/registry-cache"
  
  registry_name = "gcr.io"
  cache_port    = 5053
  
  tags = {
    Environment = "production"
    Purpose     = "container-registry-cache"
    Registry    = "google"
  }
}

# GitHub Container Registry cache (port 5054)
module "ghcr_cache" {
  source = "../../modules/registry-cache"
  
  registry_name = "ghcr.io"
  cache_port    = 5054
  
  tags = {
    Environment = "production"
    Purpose     = "container-registry-cache"
    Registry    = "github"
  }
}