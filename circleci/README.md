# ⚡ CircleCI Pipelines

CircleCI pipeline configurations for 8 tech stacks.

## Prerequisites

- CircleCI account connected to your VCS (GitHub, Bitbucket)
- Context or project-level env vars: `DOCKER_USERNAME`, `DOCKER_PASSWORD`

## Pipeline Structure

Each `config.yml` uses CircleCI 2.1 syntax with:
- **Executors**: Docker images per tech stack (cimg/openjdk, cimg/node, etc.)
- **Orbs**: Reusable packages (e.g., docker orb for build/push)
- **Workflows**: Job dependencies, branch filters
- **Caching**: restore_cache / save_cache for deps
- **Security**: Parallel (Lint | Test | Security Code) after Build; Security (Image) after Docker Build, before Push

## CI/CD Pipeline Flow Diagram

<div align="center">

<pre>
```
┌─────────────────────────────────────────────────────────────────┐
│  Git Push / PR                                                  │
│  CircleCI: workflows with branch filters                        │
└─────────────────────────────────────────────────────────────────┘
|
▼
┌─────────────────────────────────────────────────────────────────┐
│  Build                                                          │
│  executor: cimg/* | restore_cache, persist_to_workspace         │
└─────────────────────────────────────────────────────────────────┘
|
▼
┌─────────────────────────────────────────────────────────────────┐
│  Lint | Test | Security (Code)                                  │
│  Parallel: checkstyle/eslint, JUnit, Trivy fs, Semgrep, Gitleaks│
└─────────────────────────────────────────────────────────────────┘
|
▼
┌─────────────────────────────────────────────────────────────────┐
│  Docker Build                                                   │
│  setup_remote_docker | docker orb build                         │
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
│  docker orb push | DOCKER_USERNAME, DOCKER_PASSWORD             │
└─────────────────────────────────────────────────────────────────┘
|
▼
┌─────────────────────────────────────────────────────────────────┐
│  Deploy                                                         │
│  job: kubectl / helm / docker run                               │
└─────────────────────────────────────────────────────────────────┘
```
</pre>

</div>

## Stage-by-Stage Explanation

| Job | Purpose | What Happens | Artifacts / Output |
|-----|---------|--------------|--------------------|
| **build** | Compile or install deps | checkout, restore_cache, run build commands, save_cache, persist_to_workspace | Workspace: target/, node_modules, etc. |
| **lint** | Static analysis | Runs in parallel with test and security-code. checkstyle, eslint, etc. | — |
| **test** | Unit tests + coverage | Runs in parallel. store_test_results, store_artifacts. | JUnit XML, coverage reports |
| **security-code** | SAST, deps, secrets | Runs in parallel. Trivy fs, Semgrep, Snyk, OWASP Dep-Check, Gitleaks. | Scan report |
| **docker-build** | Build image | setup_remote_docker, docker orb build. | Image in local daemon |
| **Security (Image)** | Container scan | After Docker Build. Trivy, Snyk Container, Clair. Scan image for CVEs. | Scan report |
| **docker-push** | Push to registry | docker orb push. Runs after security. | Image in registry |
| **deploy** | Deploy to staging | Runs on every branch. Supports Docker or Kubernetes. Replace echo with kubectl/Helm/docker run. | — |

## Security Tools (Industry Options)

| Category | Purpose | Tools |
|----------|---------|-------|
| **Security (Code)** | SAST, dependency scan, secret scan — runs in parallel with Lint/Test | [Trivy fs](https://trivy.dev/), [Semgrep](https://semgrep.dev/), [Snyk Code](https://snyk.io/product/snyk-code/), [OWASP Dependency-Check](https://owasp.org/www-project-dependency-check/), [Gitleaks](https://github.com/gitleaks/gitleaks), [TruffleHog](https://github.com/trufflesecurity/trufflehog) |
| **Security (Image)** | Container scan — runs after Docker Build, before Push | [Trivy](https://trivy.dev/), [Snyk Container](https://snyk.io/product/container-vulnerability-management/), [Clair](https://github.com/quay/clair), [Anchore](https://anchore.com/) |

## Tech Stacks

| Stack | File | Executor Image | Lint Tool | Test Framework |
|-------|------|---------------|-----------|----------------|
| Java | [java/config.yml](java/config.yml) | cimg/openjdk:17.0 | Checkstyle | JUnit/JaCoCo |
| Node.js | [nodejs/config.yml](nodejs/config.yml) | cimg/node:18.0 | ESLint | Jest/npm test |
| Python | [python/config.yml](python/config.yml) | cimg/python:3.12 | flake8 | pytest |
| Go | [go/config.yml](go/config.yml) | cimg/go:1.21 | go vet | go test |
| .NET | [dotnet/config.yml](dotnet/config.yml) | cimg/dotnet:8.0 | dotnet format | xUnit/NUnit |
| Ruby | [ruby/config.yml](ruby/config.yml) | cimg/ruby:3.3 | RuboCop | RSpec |
| Rust | [rust/config.yml](rust/config.yml) | cimg/rust:1.73 | clippy, rustfmt | cargo test |
| PHP | [php/config.yml](php/config.yml) | cimg/php:8.2 | phpcs, phpstan | PHPUnit |

## Usage

1. Copy the desired `config.yml` to `.circleci/config.yml` in your project
2. Connect your repo to CircleCI
3. Configure environment variables in CircleCI project settings (or use a Context)
4. Update `image` in docker/build and docker/push to your registry path

## Resources

- [CircleCI Documentation](https://circleci.com/docs/)
- [CircleCI Orbs Registry](https://circleci.com/developer/orbs)
