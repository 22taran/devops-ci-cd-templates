# ⚡ CircleCI Pipelines

CircleCI pipeline configurations for 8 tech stacks.

## Prerequisites

- CircleCI account connected to your VCS
- Context or project-level env vars: `DOCKER_USERNAME`, `DOCKER_PASSWORD`

## Tech Stacks

| Stack | File |
|-------|------|
| Java | [java/config.yml](java/config.yml) |
| Node.js | [nodejs/config.yml](nodejs/config.yml) |
| Python | [python/config.yml](python/config.yml) |
| Go | [go/config.yml](go/config.yml) |
| .NET | [dotnet/config.yml](dotnet/config.yml) |
| Ruby | [ruby/config.yml](ruby/config.yml) |
| Rust | [rust/config.yml](rust/config.yml) |
| PHP | [php/config.yml](php/config.yml) |

## Usage

1. Copy the desired `config.yml` to `.circleci/config.yml` in your project
2. Connect your repo to CircleCI
3. Configure environment variables in CircleCI project settings

## Resources

- [CircleCI Documentation](https://circleci.com/docs/)
- [CircleCI Orbs Registry](https://circleci.com/developer/orbs)
