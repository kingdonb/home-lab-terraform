variable "primary_dns_port" {
  description = "DNS port for primary pi-hole instance"
  type        = number
  default     = 53
}

variable "secondary_dns_port" {
  description = "DNS port for secondary pi-hole instance" 
  type        = number
  default     = 5353
}

variable "sync_enabled" {
  description = "Enable configuration synchronization between instances"
  type        = bool
  default     = true
}

variable "synology_host" {
  description = "Hostname of Synology device for deployment"
  type        = string
  default     = ""
}

variable "pihole_password" {
  description = "Admin password for Pi-hole web interface"
  type        = string
  sensitive   = true
  default     = "admin"
}