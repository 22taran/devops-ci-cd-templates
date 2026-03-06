# ☁️ Azure DevOps Pipelines

Azure Pipelines YAML configurations for 8 tech stacks.

## Prerequisites
- Azure DevOps organization and project
- Service connections: Docker Registry, Kubernetes (optional)
- Pipeline variables or variable groups configured

## Tech Stacks

| Stack | File |
|-------|------|
| Java | [java/azure-pipelines.yml](java/azure-pipelines.yml) |
| Node.js | [nodejs/azure-pipelines.yml](nodejs/azure-pipelines.yml) |
| Python | [python/azure-pipelines.yml](python/azure-pipelines.yml) |
| Go | [go/azure-pipelines.yml](go/azure-pipelines.yml) |
| .NET | [dotnet/azure-pipelines.yml](dotnet/azure-pipelines.yml) |
| Ruby | [ruby/azure-pipelines.yml](ruby/azure-pipelines.yml) |
| Rust | [rust/azure-pipelines.yml](rust/azure-pipelines.yml) |
| PHP | [php/azure-pipelines.yml](php/azure-pipelines.yml) |

## Usage
1. Copy `azure-pipelines.yml` to your project root
2. Create pipeline in Azure DevOps pointing to the YAML file
3. Configure service connections and variables

## Resources
- [Azure Pipelines YAML Schema](https://learn.microsoft.com/en-us/azure/devops/pipelines/yaml-schema/)
