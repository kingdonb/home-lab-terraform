# Registry Caches Environment Outputs

output "registry_caches" {
  description = "Summary of all deployed registry caches"
  value = {
    docker_hub = {
      endpoint     = module.docker_cache.cache_endpoint
      container_id = module.docker_cache.container_id
      upstream     = module.docker_cache.upstream_url
    }
    kubernetes = {
      endpoint     = module.k8s_cache.cache_endpoint
      container_id = module.k8s_cache.container_id
      upstream     = module.k8s_cache.upstream_url
    }
    quay = {
      endpoint     = module.quay_cache.cache_endpoint
      container_id = module.quay_cache.container_id
      upstream     = module.quay_cache.upstream_url
    }
    google = {
      endpoint     = module.gcr_cache.cache_endpoint
      container_id = module.gcr_cache.container_id
      upstream     = module.gcr_cache.upstream_url
    }
    github = {
      endpoint     = module.ghcr_cache.cache_endpoint
      container_id = module.ghcr_cache.container_id
      upstream     = module.ghcr_cache.upstream_url
    }
  }
}

output "cache_endpoints" {
  description = "All registry cache endpoints for client configuration"
  value = [
    module.docker_cache.cache_endpoint,
    module.k8s_cache.cache_endpoint,
    module.quay_cache.cache_endpoint,
    module.gcr_cache.cache_endpoint,
    module.ghcr_cache.cache_endpoint
  ]
}

output "docker_endpoints_config" {
  description = "Docker daemon configuration for registry mirrors"
  value = {
    "registry-mirrors" = [
      "http://${module.docker_cache.cache_endpoint}"
    ]
    "insecure-registries" = [
      module.docker_cache.cache_endpoint,
      module.k8s_cache.cache_endpoint,
      module.quay_cache.cache_endpoint,
      module.gcr_cache.cache_endpoint,
      module.ghcr_cache.cache_endpoint
    ]
  }
}