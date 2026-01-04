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