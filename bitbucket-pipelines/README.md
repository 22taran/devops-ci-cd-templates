# 🪣 Bitbucket Pipelines

Bitbucket Pipelines configurations for 8 tech stacks.

## Prerequisites

- Bitbucket Cloud repository
- Repository variables: `DOCKER_USERNAME`, `DOCKER_PASSWORD`
- Deployment environments configured (optional)

## Pipeline Structure

Each `bitbucket-pipelines.yml` uses Bitbucket's native syntax with:
- **Steps**: Build → Parallel (Lint | Test | Security Code) → Docker Build → Security (Image) → Docker Push → Deploy (runs on all branches)
- **Deployment**: `deployment: staging` for environment tracking

## CI/CD Pipeline Flow Diagram

<div align="center">

<pre>
```
┌─────────────────────────────────────────────────────────────────┐
│  Git Push / PR                                                  │
│  Bitbucket: pipelines with branch filters                       │
└─────────────────────────────────────────────────────────────────┘
|
▼
┌─────────────────────────────────────────────────────────────────┐
│  Build                                                          │
│  step: script | image: maven/node/python | caches               │
└─────────────────────────────────────────────────────────────────┘
|
▼
┌─────────────────────────────────────────────────────────────────┐
│  Lint | Test | Security (Code)                                  │
│  Parallel: checkstyle/eslint, junit/pytest, Trivy fs, Semgrep,  │
│  Gitleaks                                                       │
└─────────────────────────────────────────────────────────────────┘
|
▼
┌─────────────────────────────────────────────────────────────────┐
│  Docker Build                                                   │
│  step: script | docker build with tag                           │
└─────────────────────────────────────────────────────────────────┘
|
▼
┌─────────────────────────────────────────────────────────────────┐
│  Security (Image)                                               │
│  Trivy / Snyk Container / Clair | scan built image, fail on CVE │
└─────────────────────────────────────────────────────────────────┘
|
▼
┌─────────────────────────────────────────────────────────────────┐
│  Docker Push                                                    │
│  step: script | docker login, push | DOCKER_USERNAME/PASSWORD   │
└─────────────────────────────────────────────────────────────────┘
|
▼
┌─────────────────────────────────────────────────────────────────┐
│  Deploy                                                         │
│  step: deployment: staging | kubectl / helm / docker run        │
└─────────────────────────────────────────────────────────────────┘
```
</pre>

</div>

## Stage-by-Stage Explanation

| Step | Purpose | What Happens | Artifacts / Output |
|------|---------|--------------|--------------------|
| **Build** | Compile or install deps | Maven, npm, pip, etc. Uses caches. | target/*.jar, node_modules |
| **Lint** | Static analysis | Runs in parallel with Test and Security (Code). checkstyle, ESLint, flake8, etc. | — |
| **Test** | Unit tests | Runs in parallel. JUnit, pytest, go test, etc. Some configs store coverage. | Test results |
| **Security (Code)** | SAST, deps, secrets | Runs in parallel. Trivy fs, Semgrep, Snyk Code, OWASP Dep-Check, Gitleaks. | Scan report |
| **Docker Build** | Build image | Docker build with tag. | Image in local daemon |
| **Security (Image)** | Container scan | After Docker Build. Trivy, Snyk Container, Clair. Scan image for CVEs. | Scan report |
| **Docker Push** | Push to registry | Docker login, push. | Image in registry |
| **Deploy** | Deploy | deployment: staging. Docker: docker run / docker-compose. Kubernetes: kubectl / Helm. | — |

## Security Tools (Industry Options)

| Category | Purpose | Tools |
|----------|---------|-------|
| **Security (Code)** | SAST, dependency scan, secret scan — runs in parallel with Lint/Test | [Trivy fs](https://trivy.dev/), [Semgrep](https://semgrep.dev/), [Snyk Code](https://snyk.io/product/snyk-code/), [OWASP Dependency-Check](https://owasp.org/www-project-dependency-check/), [Gitleaks](https://github.com/gitleaks/gitleaks), [TruffleHog](https://github.com/trufflesecurity/trufflehog) |
| **Security (Image)** | Container scan — runs after Docker Build, before Push | [Trivy](https://trivy.dev/), [Snyk Container](https://snyk.io/product/container-vulnerability-management/), [Clair](https://github.com/quay/clair), [Anchore](https://anchore.com/) |

## Tech Stacks

| Stack | File | Image | Lint Tool | Test Framework |
|-------|------|-------|-----------|----------------|
| Java | [java/bitbucket-pipelines.yml](java/bitbucket-pipelines.yml) | maven:3.9-eclipse-temurin-17 | Checkstyle | JUnit |
| Node.js | [nodejs/bitbucket-pipelines.yml](nodejs/bitbucket-pipelines.yml) | node:18-alpine | ESLint | Jest/npm test |
| Python | [python/bitbucket-pipelines.yml](python/bitbucket-pipelines.yml) | python:3.12-slim | flake8 | pytest |
| Go | [go/bitbucket-pipelines.yml](go/bitbucket-pipelines.yml) | golang:1.21-alpine | go vet | go test |
| .NET | [dotnet/bitbucket-pipelines.yml](dotnet/bitbucket-pipelines.yml) | mcr.microsoft.com/dotnet/sdk:8.0 | dotnet format | xUnit/NUnit |
| Ruby | [ruby/bitbucket-pipelines.yml](ruby/bitbucket-pipelines.yml) | ruby:3.3-slim | RuboCop | RSpec |
| Rust | [rust/bitbucket-pipelines.yml](rust/bitbucket-pipelines.yml) | rust:1.73-slim | clippy, rustfmt | cargo test |
| PHP | [php/bitbucket-pipelines.yml](php/bitbucket-pipelines.yml) | php:8.2-cli | phpcs, phpstan | PHPUnit |

## Usage

1. Copy `bitbucket-pipelines.yml` to your project root
2. Configure repository variables in Bitbucket: Repository settings → Pipelines → Repository variables
3. Push to trigger

## Resources

- [Bitbucket Pipelines Documentation](https://support.atlassian.com/bitbucket-cloud/docs/get-started-with-bitbucket-pipelines/)
