variable "container_name" {
  description = "Name for the pi-hole Docker container"
  type        = string
  default     = "pihole"
}

variable "network_name" {
  description = "Name for the Docker network"
  type        = string
  default     = "pihole-net"
}

variable "subnet" {
  description = "Subnet CIDR for the Docker network"
  type        = string
  default     = "172.20.0.0/16"
}

variable "dns_port" {
  description = "External port for DNS service"
  type        = number
  default     = 53
}

variable "web_port" {
  description = "External port for web interface"
  type        = number
  default     = 8080
}

variable "timezone" {
  description = "Timezone for the pi-hole container"
  type        = string
  default     = "UTC"
}

variable "web_password" {
  description = "Password for pi-hole web interface"
  type        = string
  default     = "admin"
  sensitive   = true
}

variable "upstream_dns" {
  description = "Upstream DNS servers"
  type        = string
  default     = "1.1.1.1;1.0.0.1"
}

variable "pihole_version" {
  description = "Pi-hole Docker image version"
  type        = string
  default     = "latest"
}

variable "dnsmasq_listening" {
  description = "Dnsmasq listening mode"
  type        = string
  default     = "all"
}

variable "extra_volumes" {
  description = "Additional volumes to mount in the container"
  type = list(object({
    volume_name    = string
    container_path = string
  }))
  default = []
}

variable "use_host_network" {
  description = "Use host networking mode instead of bridge (required for multi-subnet DNS)"
  type        = bool
  default     = false
}

variable "shm_size_mb" {
  description = "Shared memory size in MB"
  type        = number
  default     = 256
}

variable "capabilities" {
  description = "Linux capabilities to add to the container"
  type        = list(string)
  default     = ["NET_ADMIN", "SYS_TIME", "SYS_NICE"]
}