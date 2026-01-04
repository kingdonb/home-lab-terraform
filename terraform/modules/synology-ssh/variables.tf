variable "host" {
  description = "Hostname or IP address of the Synology device"
  type        = string
}

variable "port" {
  description = "SSH port for connection"
  type        = number
  default     = 22
}

variable "user" {
  description = "SSH username (overridden by 1Password credentials)"
  type        = string
  default     = "admin"
}