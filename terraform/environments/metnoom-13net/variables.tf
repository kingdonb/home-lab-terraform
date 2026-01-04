# METNOOM 13-net Environment Variables

variable "pihole_api_token" {
  description = "API token for Pi-hole metrics access"
  type        = string
  sensitive   = true
}