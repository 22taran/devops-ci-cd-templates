# ⚙️ Jenkins Pipelines

Declarative Jenkins pipeline configurations for 8 tech stacks.

## Prerequisites

- Jenkins 2.387+ with Pipeline plugin
- Docker Pipeline plugin (for Docker agents)
- Credentials configured: `docker-hub-credentials`, `deploy-credentials`

## Pipeline Structure

Each `Jenkinsfile` uses declarative pipeline syntax with:
- **Docker-based agents** for consistent builds
- **5 stages**: Build → Lint → Test → Docker → Deploy
- **Post-build actions** for cleanup and notifications
- **Environment variables** for configuration

## Tech Stacks

| Stack | File | Docker Agent |
|-------|------|-------------|
| Java | [java/Jenkinsfile](java/Jenkinsfile) | `maven:3.9-eclipse-temurin-17` |
| Node.js | [nodejs/Jenkinsfile](nodejs/Jenkinsfile) | `node:18-alpine` |
| Python | [python/Jenkinsfile](python/Jenkinsfile) | `python:3.12-slim` |
| Go | [go/Jenkinsfile](go/Jenkinsfile) | `golang:1.21-alpine` |
| .NET | [dotnet/Jenkinsfile](dotnet/Jenkinsfile) | `mcr.microsoft.com/dotnet/sdk:8.0` |
| Ruby | [ruby/Jenkinsfile](ruby/Jenkinsfile) | `ruby:3.3-slim` |
| Rust | [rust/Jenkinsfile](rust/Jenkinsfile) | `rust:1.73-slim` |
| PHP | [php/Jenkinsfile](php/Jenkinsfile) | `php:8.2-cli` |

## Usage

1. Copy the desired `Jenkinsfile` to your project root
2. Update `DOCKER_IMAGE` and `DEPLOY_ENV` environment variables
3. Configure a Jenkins Pipeline job pointing to your repo
4. Ensure required credentials are configured in Jenkins

## Resources

- [Jenkins Pipeline Syntax](https://www.jenkins.io/doc/book/pipeline/syntax/)
- [Docker Pipeline Plugin](https://plugins.jenkins.io/docker-workflow/)
