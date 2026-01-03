terraform {
  required_version = ">= 1.6"
  
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 3.0"
    }
    # TODO: Add additional providers as needed
    # ssh = {
    #   source = "loafoe/ssh" 
    #   version = "~> 2.0"
    # }
  }
  
  # TODO: Configure state backend after decision
  # Options researched:
  # 1. Local files (simple, single-user)
  # 2. Synology NFS share (reliable local storage)  
  # 3. Git repository (version controlled)
  # 4. Cloud backend (future-proof)
  #
  # backend "local" {
  #   path = "terraform.tfstate"
  # }
}

# Provider configurations will be environment-specific
# These will be configured after TDG research determines best approach

provider "docker" {
  # Connection method to be determined:
  # Option A: SSH to Synology devices
  # host = "ssh://user@10.17.12.109:22" # DS923+ subnet 1
  # host = "ssh://user@10.17.13.204:22" # DS923+ subnet 2
  
  # Option B: Docker API over TCP  
  # host = "tcp://10.17.12.109:2376"
  
  # Option C: Local socket (if running from Synology)
  # host = "unix:///var/run/docker.sock"
}

# Additional provider blocks may be needed for:
# - SSH connections for network configuration
# - Synology DSM API integration
# - Custom provider for interface management