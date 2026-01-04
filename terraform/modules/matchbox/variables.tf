# Matchbox Module Variables

variable "container_name" {
  description = "Name for the Matchbox container"
  type        = string
  default     = "matchbox"
}

variable "matchbox_image" {
  description = "Matchbox Docker image to use"
  type        = string
  default     = "kingdonb/matchbox:v1.10.5-cozy-spin-tailscale"
  
  validation {
    condition = contains([
      "kingdonb/matchbox:v1.9.5-cozy-spin",
      "kingdonb/matchbox:v1.10.5-cozy-spin-tailscale",
      "ghcr.io/aenix-io/cozystack/matchbox:v1.7.6-v0.15.0",
      "ghcr.io/aenix-io/cozystack/matchbox:v1.9.3-v0.25.1"
    ], var.matchbox_image)
    error_message = "Use a supported Matchbox image variant."
  }
}

variable "vlan_network_name" {
  description = "Name of the VLAN Docker network to attach to"
  type        = string
}

variable "static_ip" {
  description = "Static IP address for the Matchbox container"
  type        = string
  default     = "10.17.13.251"
}

variable "matchbox_port" {
  description = "Port for the Matchbox HTTP API"
  type        = number
  default     = 8080
}

variable "log_level" {
  description = "Log level for Matchbox server"
  type        = string
  default     = "debug"
  
  validation {
    condition = contains(["debug", "info", "warn", "error"], var.log_level)
    error_message = "Log level must be one of: debug, info, warn, error"
  }
}

variable "data_volume_name" {
  description = "Docker volume name for Matchbox data (profiles, groups, assets)"
  type        = string
  default     = null
}

variable "additional_volumes" {
  description = "Additional volume mounts for the container"
  type = list(object({
    host_path      = string
    container_path = string
    read_only      = optional(bool, false)
  }))
  default = []
}

variable "additional_env_vars" {
  description = "Additional environment variables for Matchbox"
  type        = list(string)
  default     = []
}

variable "enable_https" {
  description = "Enable HTTPS for Matchbox server"
  type        = bool
  default     = false
}

variable "cert_file_path" {
  description = "Path to TLS certificate file (if HTTPS enabled)"
  type        = string
  default     = null
}

variable "key_file_path" {
  description = "Path to TLS private key file (if HTTPS enabled)"
  type        = string
  default     = null
}

variable "tags" {
  description = "Additional tags for resources"
  type        = map(string)
  default     = {}
}