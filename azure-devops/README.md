# ☁️ Azure DevOps Pipelines

Azure Pipelines YAML configurations for 8 tech stacks.

## Prerequisites

- Azure DevOps organization and project
- Service connections: Docker Registry, Kubernetes (optional)
- Pipeline variables or variable groups configured

## Pipeline Structure

Each `azure-pipelines.yml` uses multi-stage pipelines with:
- **Stages**: Build → Parallel (Lint | Test | Security Code) → Docker Build → Security (Image) → Docker Push → Deploy
- **Tasks**: Maven, Node, Python, DotNet, etc.
- **Artifacts**: publish between stages
- **Approval gates**: Optional for deploy stage

## CI/CD Pipeline Flow Diagram

<div align="center">

<pre>
```
┌─────────────────────────────────────────────────────────────────┐
│  Git Push / PR                                                  │
│  Azure DevOps: trigger: branches                                │
└─────────────────────────────────────────────────────────────────┘
|
▼
┌─────────────────────────────────────────────────────────────────┐
│  Build                                                          │
│  stage: Build | task: Maven, Node, DotNet, etc. | publish       │
└─────────────────────────────────────────────────────────────────┘
|
▼
┌─────────────────────────────────────────────────────────────────┐
│  Lint | Test | Security (Code)                                  │
│  Parallel: checkstyle/eslint, junit, Trivy fs, Semgrep, Gitleaks│
└─────────────────────────────────────────────────────────────────┘
|
▼
┌─────────────────────────────────────────────────────────────────┐
│  Docker Build                                                   │
│  stage: Docker | Docker task build                              │
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
│  stage: Docker | Docker task push | service connection          │
└─────────────────────────────────────────────────────────────────┘
|
▼
┌─────────────────────────────────────────────────────────────────┐
│  Deploy                                                         │
│  stage: Deploy | approval gates optional | kubectl / helm       │
└─────────────────────────────────────────────────────────────────┘
```
</pre>

</div>

## Stage-by-Stage Explanation

| Stage | Purpose | What Happens | Artifacts / Output |
|-------|---------|--------------|--------------------|
| **Build** | Compile or install deps | Maven, npm, dotnet, etc. publish artifact. | java-build, publish/ |
| **Lint** | Static analysis | Runs in parallel with Test and Security (Code). checkstyle, ESLint, flake8, etc. | — |
| **Test** | Unit tests + coverage | Runs in parallel. publishJUnitResults, JaCoCo/Cobertura. | Test results |
| **Security (Code)** | SAST, deps, secrets | Runs in parallel. Trivy fs, Semgrep, Snyk Code, OWASP Dep-Check, Gitleaks. | Scan report |
| **Docker Build** | Build image | Docker task build. | Image in local daemon |
| **Security (Image)** | Container scan | After Docker Build. Trivy, Snyk Container, Clair. Scan image for CVEs. | Scan report |
| **Docker Push** | Push to registry | Docker task push. Runs on every branch. | Image in registry |
| **Deploy** | Deploy to staging | Runs on every branch. Supports Docker or Kubernetes. Replace with kubectl/Helm/docker run. | — |

## Security Tools (Industry Options)

| Category | Purpose | Tools |
|----------|---------|-------|
| **Security (Code)** | SAST, dependency scan, secret scan — runs in parallel with Lint/Test | [Trivy fs](https://trivy.dev/), [Semgrep](https://semgrep.dev/), [Snyk Code](https://snyk.io/product/snyk-code/), [OWASP Dependency-Check](https://owasp.org/www-project-dependency-check/), [Gitleaks](https://github.com/gitleaks/gitleaks), [TruffleHog](https://github.com/trufflesecurity/trufflehog) |
| **Security (Image)** | Container scan — runs after Docker Build, before Push | [Trivy](https://trivy.dev/), [Snyk Container](https://snyk.io/product/container-vulnerability-management/), [Clair](https://github.com/quay/clair), [Anchore](https://anchore.com/) |

## Tech Stacks

| Stack | File | VM Image | Lint Tool | Test Framework |
|-------|------|----------|-----------|----------------|
| Java | [java/azure-pipelines.yml](java/azure-pipelines.yml) | ubuntu-latest | Checkstyle | JUnit/JaCoCo |
| Node.js | [nodejs/azure-pipelines.yml](nodejs/azure-pipelines.yml) | ubuntu-latest | ESLint | Jest/npm test |
| Python | [python/azure-pipelines.yml](python/azure-pipelines.yml) | ubuntu-latest | flake8 | pytest |
| Go | [go/azure-pipelines.yml](go/azure-pipelines.yml) | ubuntu-latest | go vet | go test |
| .NET | [dotnet/azure-pipelines.yml](dotnet/azure-pipelines.yml) | ubuntu-latest | dotnet format | xUnit/NUnit |
| Ruby | [ruby/azure-pipelines.yml](ruby/azure-pipelines.yml) | ubuntu-latest | RuboCop | RSpec |
| Rust | [rust/azure-pipelines.yml](rust/azure-pipelines.yml) | ubuntu-latest | clippy, rustfmt | cargo test |
| PHP | [php/azure-pipelines.yml](php/azure-pipelines.yml) | ubuntu-latest | phpcs, phpstan | PHPUnit |

## Usage

1. Copy `azure-pipelines.yml` to your project root
2. Create pipeline in Azure DevOps pointing to the YAML file
3. Configure service connections and variables
4. Run the pipeline

## Resources

- [Azure Pipelines YAML Schema](https://learn.microsoft.com/en-us/azure/devops/pipelines/yaml-schema/)
