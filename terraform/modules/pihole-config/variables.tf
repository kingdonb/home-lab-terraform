variable "pihole_base_url" {
  description = "Base URL for the Pi-hole instance (e.g., http://localhost:8080)"
  type        = string
}

variable "pihole_password" {
  description = "Pi-hole admin password for authentication"
  type        = string
  sensitive   = true
}

# Client device configuration
variable "work_laptop_ip" {
  description = "IP address for work laptop"
  type        = string
  default     = "10.17.12.100"
}

variable "work_laptop_mac" {
  description = "MAC address for work laptop"
  type        = string
  default     = "00:11:22:33:44:55"
}

variable "work_phone_ip" {
  description = "IP address for work phone"
  type        = string
  default     = "10.17.12.101"
}

variable "work_phone_mac" {
  description = "MAC address for work phone"
  type        = string
  default     = "00:11:22:33:44:56"
}

variable "other_laptop_ip" {
  description = "IP address for other laptop"
  type        = string
  default     = "10.17.13.100"
}

variable "other_laptop_mac" {
  description = "MAC address for other laptop"
  type        = string
  default     = "00:11:22:33:44:57"
}

variable "phone_24ghz_ip" {
  description = "IP address for phone on 2.4GHz"
  type        = string
  default     = "10.17.13.101"
}

variable "phone_24ghz_mac" {
  description = "MAC address for phone on 2.4GHz"
  type        = string
  default     = "00:11:22:33:44:58"
}