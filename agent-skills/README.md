# Agent Skills Configuration

This directory contains configuration and documentation for Agent Skills used in this Terraform project.

## TDG (Terraform Documentation Generator) Integration

The TDG skill from Chanwit Kaewkasi helps with:
- Terraform module discovery and selection
- Automated documentation generation
- Best practice recommendations

### Setup Instructions

1. **Configure TDG Skill**:
   1. Ensure you have the latest GitHub Copilot extension installed in VS Code (Insiders or Stable).
   2. Create a file named `agent.yaml` in this directory with the following content:
      ```yaml
      skills:
        - id: chanwit/tdg
          enabled: true
      ```
   3. Save the file.
   4. Reload your VS Code window (Command Palette → "Developer: Reload Window").
   5. Copilot will automatically detect and activate the TDG skill.

   You can now use TDG features via Copilot Chat or the Command Palette.

2. **Configure for Home Lab Use**:
   - Focus on Docker container management modules
   - Synology NAS infrastructure modules
   - Networking and DNS management modules

### Module Research Targets

The following Terraform modules need to be researched and evaluated:

#### Docker Management
- [ ] `kreuzwerker/docker` provider
- [ ] `terraform-docker-modules/*` collection
- [ ] Custom SSH-based Docker management modules

#### DNS and Networking
- [ ] Pi-hole specific modules
- [ ] DNS server management modules
- [ ] Multi-interface network configuration

#### Synology Integration
- [ ] DSM API integration modules
- [ ] Docker Compose to Terraform migration tools
- [ ] Configuration backup and restore modules

### Usage Workflow

1. **Module Discovery**: Use TDG to identify suitable modules for each service
2. **Evaluation**: Test modules in isolated environment
3. **Documentation**: Generate module documentation automatically
4. **Integration**: Incorporate selected modules into home lab configuration

## Future Agent Skills

Additional skills to consider:
- Infrastructure testing and validation
- Configuration drift detection
- Automated backup verification

---

*This configuration will be updated as Agent Skills are implemented and tested.*