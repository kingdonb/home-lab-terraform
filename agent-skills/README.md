# Agent Skills Configuration

This directory contains configuration and documentation for Agent Skills used in this Terraform project.

## TDG (Terraform Documentation Generator) Integration

The TDG skill from Chanwit Kaewkasi helps with:
- Terraform module discovery and selection
- Automated documentation generation
- Best practice recommendations


### Setup & Integration Notes

**What we actually did:**

- We followed the official documentation at https://code.visualstudio.com/docs/copilot/customization/agent-skills for local agent skills.
- Instead of using `agent.yaml`, we cloned the `tdg` skill repository and placed the relevant skills in `.github/skills/` as described in the docs.
- After reloading VS Code, Copilot Chat can now recognize and use the local skills (TDG and Atomic Commit).

**Next step: Initialize TDG**

The TDG skill expects a `TDG.md` file in your project. If it does not exist, you should initialize it by running the following command in Copilot Chat:

      /tdg:init

This will create the required TDG.md file and set up the project for TDG workflows.

You can then proceed with the TDD workflow as described in the skill's instructions.

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