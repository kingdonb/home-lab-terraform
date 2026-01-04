# DNSmasq Module Variables

variable "container_name" {
  description = "Name for the DNSmasq container"
  type        = string
  default     = "dnsmasq"
}

variable "dnsmasq_image" {
  description = "DNSmasq Docker image to use"
  type        = string
  default     = "quay.io/poseidon/dnsmasq:v0.5.0-40-g494d4e0"
}

variable "vlan_network_name" {
  description = "Name of the VLAN Docker network to attach to"
  type        = string
}

variable "static_ip" {
  description = "Static IP address for the DNSmasq container"
  type        = string
}

# DHCP Configuration
variable "dhcp_enabled" {
  description = "Enable DHCP server functionality"
  type        = bool
  default     = true
}

variable "dhcp_range_start" {
  description = "Start of DHCP IP range"
  type        = string
  default     = "10.17.13.3"
}

variable "dhcp_range_end" {
  description = "End of DHCP IP range"
  type        = string
  default     = "10.17.13.199"
}

variable "dhcp_router" {
  description = "Default gateway for DHCP clients"
  type        = string
  default     = "10.17.13.249"
}

variable "dhcp_dns_server" {
  description = "DNS server for DHCP clients"
  type        = string
  default     = "10.17.13.254"
}

# TFTP Configuration
variable "tftp_enabled" {
  description = "Enable TFTP server for PXE booting"
  type        = bool
  default     = true
}

variable "tftp_volume_name" {
  description = "Docker volume name for TFTP boot files"
  type        = string
  default     = "tftpboot"
}

# PXE Boot Configuration
variable "pxe_enabled" {
  description = "Enable PXE boot configuration"
  type        = bool
  default     = true
}

variable "matchbox_server" {
  description = "IP address of the Matchbox server for PXE boot"
  type        = string
  default     = "10.17.13.251"
}

variable "matchbox_port" {
  description = "Port of the Matchbox server"
  type        = number
  default     = 8080
}

# Logging
variable "enable_logging" {
  description = "Enable query and DHCP logging"
  type        = bool
  default     = true
}

# Additional Configuration
variable "additional_args" {
  description = "Additional DNSmasq command line arguments"
  type        = list(string)
  default     = []
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

variable "tags" {
  description = "Additional tags for resources"
  type        = map(string)
  default     = {}
}