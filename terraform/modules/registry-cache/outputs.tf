# Registry Cache Module Outputs

output "container_id" {
  description = "Docker container ID of the registry cache"
  value       = docker_container.registry_cache.id
}

output "container_name" {
  description = "Name of the registry cache container"
  value       = docker_container.registry_cache.name
}

output "cache_endpoint" {
  description = "Registry cache endpoint URL"
  value       = "localhost:${var.cache_port}"
}

output "upstream_registry" {
  description = "Upstream registry being cached"
  value       = var.registry_name
}

output "upstream_url" {
  description = "Upstream registry URL"
  value       = local.upstream_url
}

output "cache_port" {
  description = "External port for the cache"
  value       = var.cache_port
}

output "volume_name" {
  description = "Docker volume name for cache storage"
  value       = docker_volume.registry_cache.name
}

output "network_id" {
  description = "Docker network ID (if network was created)"
  value       = var.create_network ? docker_network.registry_network[0].id : null
}

output "cache_summary" {
  description = "Summary of registry cache configuration"
  value = {
    registry_name = var.registry_name
    cache_port    = var.cache_port
    upstream_url  = local.upstream_url
    endpoint      = "localhost:${var.cache_port}"
    volume        = docker_volume.registry_cache.name
    container_id  = docker_container.registry_cache.id
  }
}