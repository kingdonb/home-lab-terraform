# Registry Cache Module Variables

variable "registry_name" {
  description = "Name of the upstream registry to cache (e.g., docker.io, registry.k8s.io)"
  type        = string
  
  validation {
    condition = contains([
      "docker.io",
      "registry.k8s.io", 
      "quay.io",
      "gcr.io",
      "ghcr.io"
    ], var.registry_name)
    error_message = "Supported registries: docker.io, registry.k8s.io, quay.io, gcr.io, ghcr.io"
  }
}

variable "cache_port" {
  description = "External port for the registry cache"
  type        = number
  
  validation {
    condition     = var.cache_port >= 1024 && var.cache_port <= 65535
    error_message = "Cache port must be between 1024 and 65535."
  }
}

variable "upstream_url" {
  description = "Custom upstream registry URL (overrides default for registry_name)"
  type        = string
  default     = null
}

variable "registry_image" {
  description = "Docker registry image to use"
  type        = string
  default     = "registry:2"
}

variable "create_network" {
  description = "Whether to create a dedicated network for the registry"
  type        = bool
  default     = false
}

variable "network_name" {
  description = "Name of the Docker network (used if create_network is true)"
  type        = string
  default     = "registry-cache-net"
}

variable "network_subnet" {
  description = "Subnet for the Docker network (used if create_network is true)"
  type        = string
  default     = "172.22.0.0/16"
}

variable "tags" {
  description = "Additional tags for resources"
  type        = map(string)
  default     = {}
}