# Questions and Answers Log

*Session Date: January 3, 2026*

## Questions Asked & Answered ✅

### Network & Hardware
**Q: What models are your Synology devices?**
A: DS1517+ (Intel Atom C2538) running DSM 7.2.1-69057, DS923+ (AMD Ryzen R1600) running DSM 7.2.2-72806

**Q: Which Synology runs Pi-hole? Is it on the 13-net only?**
A: DS923+ runs Pi-hole, yes it's on the 13-net only

**Q: Do the Synology devices have static IP assignments, or DHCP reservations?**
A: Static assignments because they're infrastructure and they might run DHCP at times

**Q: What version of Pi-hole are you currently running? What target version?**
A: Current version doesn't matter, target is latest version with completely new configuration

**Q: Do you have current Pi-hole configuration documented? Custom blocklists?**
A: Yes, in the Docker Compose tree. Yes, lots of custom configuration that needs to be captured

**Q: Are you using Docker Compose for all containers?**
A: Only for Pi-hole. Everything else uses scripts and should move to Terraform. One existing compose file but followed first tutorial found - can be replaced

## Outstanding Questions ❓

### Docker Management & Access
**Q: How do you currently manage Docker on the Synology devices?**
- SSH access with docker commands?
- Synology Docker UI?
- API access?
- *This determines which Terraform provider approach we use*

### Terraform State Management  
**Q: Where would you prefer to store Terraform state?**
Options:
- Local files (simple, but no collaboration)
- One of the Synology devices (reliable local storage)
- Cloud backend (Terraform Cloud, S3, etc.)

### Current Configuration Details
**Q: Can you provide the existing Pi-hole Docker Compose file?**
- Need to see current configuration for reference
- Need to understand volume mounts and network configuration
- Need to capture custom configurations

**Q: What are the current Docker run scripts for other services?**
- Tailscale container configuration
- Plex configuration  
- Registry pull-through cache setup
- Matchbox configuration
- Other container configurations

### DHCP Strategy Decision
**Q: Router-based DHCP vs centralized dnsmasq?**
- Current preference unclear
- Affects overall network architecture
- May impact Pi-hole configuration

### Authentication & Access
**Q: How will Terraform authenticate to Synology devices?**
- SSH keys?
- API tokens?
- Local management from one of the Synology devices?

## Agent Skills Setup ❓

**Q: TDG skill configuration in VSCode Insiders?**
A: Not familiar with setup, needs to switch to VSCode Insiders first

## Research Questions for TDG

### Terraform Providers to Investigate
- `kreuzwerker/docker` provider capabilities
- Synology DSM API integration options
- SSH-based Docker management modules
- Network interface configuration management

### Module Discovery Targets
- Pi-hole specific Terraform modules
- Docker Compose to Terraform migration tools
- Multi-subnet DNS configuration modules
- Container health checking and dependency management

## Implementation Priority Questions

**Q: Bootstrap order - what should come online first after reboot?**
- Network interfaces
- Pi-hole (critical for DNS)
- Other services in dependency order

**Q: Backup and restore strategy for configurations?**
- How to backup current Pi-hole custom configs before migration
- How to ensure rollback capability
- Configuration drift detection

---

*This log will be updated as more questions are answered and new ones arise.*