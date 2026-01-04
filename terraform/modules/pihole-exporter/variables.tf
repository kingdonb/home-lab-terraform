# Pi-hole Exporter Module Variables

variable "container_name" {
  description = "Name for the Pi-hole Exporter container"
  type        = string
  default     = "pihole-exporter"
}

variable "exporter_image" {
  description = "Pi-hole Exporter Docker image to use"
  type        = string
  default     = "ekofr/pihole-exporter:v0.4.0"
}

variable "pihole_hostname" {
  description = "Pi-hole hostname or IP address"
  type        = string
}

variable "pihole_api_token" {
  description = "Pi-hole API token for metrics access"
  type        = string
  sensitive   = true
}

variable "exporter_port" {
  description = "Port for the exporter metrics endpoint"
  type        = number
  default     = 9617
  
  validation {
    condition     = var.exporter_port >= 1024 && var.exporter_port <= 65535
    error_message = "Exporter port must be between 1024 and 65535."
  }
}

variable "scrape_interval" {
  description = "Interval for scraping Pi-hole metrics"
  type        = string
  default     = "10s"
  
  validation {
    condition = can(regex("^[0-9]+[smh]$", var.scrape_interval))
    error_message = "Scrape interval must be in format like '10s', '1m', '1h'."
  }
}

variable "additional_env_vars" {
  description = "Additional environment variables for the exporter"
  type        = list(string)
  default     = []
}

variable "enable_prometheus_labels" {
  description = "Add Prometheus discovery labels to the container"
  type        = bool
  default     = true
}

variable "metrics_path" {
  description = "Path for metrics endpoint"
  type        = string
  default     = "/metrics"
}

variable "tags" {
  description = "Additional tags for resources"
  type        = map(string)
  default     = {}
}