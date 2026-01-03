# Variable definitions for home lab infrastructure
# These will be refined after TDG research and current config analysis

# Network Configuration
variable "subnet_1_cidr" {
  description = "Primary subnet CIDR"
  type        = string
  default     = "10.17.12.0/24"
}

variable "subnet_2_cidr" {
  description = "Secondary subnet CIDR" 
  type        = string
  default     = "10.17.13.0/24"
}

# Synology Device Configuration
variable "ds923_ip_subnet1" {
  description = "DS923+ IP address on subnet 1"
  type        = string
  default     = "10.17.12.109"
}

variable "ds923_ip_subnet2" {
  description = "DS923+ IP address on subnet 2"
  type        = string
  default     = "10.17.13.204"
}

variable "ds1517_ip" {
  description = "DS1517+ IP address (to be determined)"
  type        = string
  default     = "" # TBD based on current config
}

# Router Configuration
variable "mikrotik_ip_subnet1" {
  description = "Mikrotik router IP on subnet 1"
  type        = string
  default     = "10.17.12.249"
}

variable "mikrotik_ip_subnet2" {
  description = "Mikrotik router IP on subnet 2"
  type        = string
  default     = "10.17.13.249"
}

# Pi-hole Configuration
variable "pihole_version" {
  description = "Pi-hole container version"
  type        = string
  default     = "latest" # Will be pinned after research
}

variable "pihole_admin_password" {
  description = "Pi-hole admin password"
  type        = string
  sensitive   = true
  # Will be set via environment variable or .tfvars
}

variable "pihole_timezone" {
  description = "Timezone for Pi-hole"
  type        = string
  default     = "America/New_York" # Adjust as needed
}

# Docker Configuration
variable "docker_host_ds923" {
  description = "Docker host connection string for DS923+"
  type        = string
  # Will be determined during TDG research phase
  # Examples:
  # "ssh://user@10.17.12.109:22"
  # "tcp://10.17.12.109:2376" 
  # "unix:///var/run/docker.sock"
}

variable "docker_host_ds1517" {
  description = "Docker host connection string for DS1517+"
  type        = string
  # TBD based on final service distribution
}

# Service Configuration  
variable "container_restart_policy" {
  description = "Default restart policy for containers"
  type        = string
  default     = "unless-stopped"
}

variable "enable_container_logging" {
  description = "Enable container logging"
  type        = bool
  default     = true
}

# Backup and Storage
variable "config_backup_path" {
  description = "Path for configuration backups"
  type        = string
  default     = "/volume1/docker/backups"
}

variable "pihole_config_path" {
  description = "Host path for Pi-hole configuration persistence"
  type        = string
  default     = "/volume1/docker/pihole"
}

# Network Services
variable "custom_dns_servers" {
  description = "Upstream DNS servers for Pi-hole"
  type        = list(string)
  default     = ["1.1.1.1", "8.8.8.8"] # Will be customized based on current config
}

variable "enable_dhcp" {
  description = "Enable DHCP server in Pi-hole"
  type        = bool
  default     = false # TBD based on architecture decision
}