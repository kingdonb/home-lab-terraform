# Pi-hole Provider Analysis & Roadmap

## Provider Evolution Journey

### Initial Attempt: ryanwholey/pihole
- **Status**: Unmaintained since 2022
- **Pi-hole v6 Compatibility**: ❌ Failed authentication 
- **Issue**: Uses legacy `/admin/api.php` endpoints removed in v6
- **GitHub Source**: Syntax errors with OpenTofu

### Solution: lukaspustina/pihole  
- **Status**: Actively maintained fork of ryanwholey's work
- **Pi-hole v6 Compatibility**: ✅ JSON session authentication
- **Version**: v0.3+ supports Pi-hole v6+ features
- **Authentication**: Modern session-based with cookie capture

## Current Provider Capabilities

### ✅ Working Features
- **DNS Records**: `pihole_dns_record` for A record management
- **CNAME Records**: `pihole_cname_record` for alias management  
- **Basic Config**: `pihole_config` for webserver settings
- **Authentication**: JSON session with admin password

### 📋 Limited Configuration Management
```hcl
# Available config management (requires admin password)
resource "pihole_config" "enable_app_sudo" {
  key   = "webserver.api.app_sudo" 
  value = "true"
}

# Read configuration status
data "pihole_config" "app_sudo_status" {
  key = "webserver.api.app_sudo"
}
```

### ❌ Missing Advanced Features  
- **Groups Management**: Cannot create/manage client groups
- **Client Management**: No client assignment to groups
- **Blacklist Management**: No domain blacklist automation
- **Adlist Management**: No blocklist source management
- **Whitelist Management**: No domain whitelist automation

## Future Roadmap

### Phase 1: Foundation (COMPLETE)
- ✅ Pi-hole v6 authentication working
- ✅ Basic DNS/CNAME record management
- ✅ TDG methodology demonstrated

### Phase 2: Infrastructure Prerequisites (IN PROGRESS)
- 🔄 Registry pull-through cache modules
- 🔄 DNSmasq deployment for 13-net
- 🔄 Matchbox TFTP boot services
- 🔄 VLAN configuration management

### Phase 3: Advanced Pi-hole Features (PLANNED)
**Approach**: Crossplane Functions with `go-pihole` library

**Rationale**: 
- Runtime API control vs static IaC configuration
- Better suited for dynamic group/client management
- Kubernetes-native approach after cluster is online
- Can leverage `github.com/KimMachineGun/go-pihole` library

**Advanced Features via Crossplane**:
- Group creation and management
- Client assignment to groups  
- Dynamic blacklist/whitelist management
- Adlist source configuration
- Client device registration

### Configuration Backup Strategy
- **Current**: Manual backup/restore of Pi-hole configuration
- **Planned**: IaC-managed backup/restore process
- **Approach**: Volume management + configuration export/import

## Technical Constraints

### Provider Limitations
- Limited to basic DNS record and minimal config management
- No group/client/list management capabilities
- Admin password required for configuration changes
- TLS/connection settings only

### Workaround Approaches
1. **Terraform**: Foundation DNS infrastructure + basic config
2. **Manual Setup**: Initial group/client/list configuration  
3. **Backup/Restore**: IaC-managed configuration preservation
4. **Crossplane**: Future advanced runtime management

## Implementation Notes

### Authentication Requirements
```hcl
# Admin password required (not application password)
provider "pihole" {
  url      = var.pihole_base_url
  password = var.pihole_password  # Must be admin password
}
```

### DNS Foundation Pattern
```hcl
# Homelab services DNS records
resource "pihole_dns_record" "homelab_services" {
  domain = "gateway.homelab.local"
  ip     = "10.17.12.1"
}

# Registry aliases via CNAME
resource "pihole_cname_record" "service_aliases" {
  domain = "docker.homelab.local" 
  target = "registry.homelab.local"
}
```

This provides the DNS foundation needed before Kubernetes deployment while leaving advanced Pi-hole management for future Crossplane implementation.