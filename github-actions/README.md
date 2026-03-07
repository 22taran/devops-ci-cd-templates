# 🐙 GitHub Actions Workflows

GitHub-native CI/CD workflow configurations for 8 tech stacks.

## Prerequisites

- GitHub repository with Actions enabled
- Repository secrets: `DOCKER_USERNAME`, `DOCKER_PASSWORD`
- Optional: `DEPLOY_TOKEN` for deployment

## Workflow Structure

Each `ci.yml` workflow includes:
- **Triggers**: Push to `main`/`develop`, pull requests to `main`
- **Caching**: Language-specific dependency caching (Maven, npm, pip, etc.)
- **7 jobs**: Build → Lint → Test → Docker Build → Security Scan → Docker Push → Deploy
- **Artifact uploads** for build outputs and coverage reports

## CI/CD Pipeline Flow Diagram

<div align="center">

<pre>
```
┌─────────────────────────────────────────────────────────────────┐
│  Git Push / PR                                                  │
│  GitHub Actions: on: push, pull_request | workflow_dispatch     │
└─────────────────────────────────────────────────────────────────┘
|
▼
┌─────────────────────────────────────────────────────────────────┐
│  Build                                                          │
│  jobs: runs-on ubuntu-latest | actions/setup-*, actions/cache   │
└─────────────────────────────────────────────────────────────────┘
|
▼
┌─────────────────────────────────────────────────────────────────┐
│  Lint                                                           │
│  job: runs in parallel with test (checkstyle/eslint/flake8/etc) │
└─────────────────────────────────────────────────────────────────┘
|
▼
┌─────────────────────────────────────────────────────────────────┐
│  Test                                                           │
│  job: coverage artifact upload | JUnit, Cobertura               │
└─────────────────────────────────────────────────────────────────┘
|
▼
┌─────────────────────────────────────────────────────────────────┐
│  Docker Build                                                   │
│  Docker Buildx: build with load: true, tags: sha + latest       │
└─────────────────────────────────────────────────────────────────┘
|
▼
┌─────────────────────────────────────────────────────────────────┐
│  Security Scan                                                  │
│  Trivy / CodeQL / Snyk / Gitleaks | fail on critical CVEs      │
└─────────────────────────────────────────────────────────────────┘
|
▼
┌─────────────────────────────────────────────────────────────────┐
│  Docker Push                                                    │
│  docker/login-action + docker push to registry                  │
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
| **Triggers** | When to run | `on.push` to main/develop, `on.pull_request` to main. Configurable per workflow. | — |
| **Build** | Compile or install deps | setup-node/java/python/etc, install deps, compile. Uses `actions/cache` or built-in cache. | Uploaded artifacts (JAR, node_modules cache key) |
| **Lint** | Static analysis | Runs in parallel with test. checkstyle, ESLint, flake8, go vet, etc. Fails workflow on violations. | — |
| **Test** | Unit tests + coverage | Runs tests, uploads coverage as artifact. Some workflows use matrix for multi-version testing. | coverage-report artifact |
| **Docker Build** | Build image | Docker Buildx, build with `push: false` and `load: true`. Tags: `sha` + `latest`. | Image in local daemon |
| **Security Scan** | Vulnerability check | Optional. Trivy (container), CodeQL (SAST), Snyk, Gitleaks (secrets). Fail on critical CVEs. | Scan report |
| **Docker Push** | Push to registry | Login, push image. Runs after build. | Image in registry |
| **Deploy** | Deploy to staging | Runs every run. Supports Docker (docker run) or Kubernetes (kubectl/Helm). Replace with your deploy logic. | — |

## Tech Stacks

| Stack | File | Runner | Lint Tool | Test Framework | Extra Jobs |
|-------|------|--------|-----------|----------------|------------|
| Java | [java/ci.yml](java/ci.yml) | `ubuntu-latest` + JDK 17 | Checkstyle | JUnit/JaCoCo | Package in build |
| Node.js | [nodejs/ci.yml](nodejs/ci.yml) | `ubuntu-latest` + Node 18 | ESLint | Jest/npm test | Matrix: Node 18,20,22 |
| Python | [python/ci.yml](python/ci.yml) | `ubuntu-latest` + Python 3.12 | flake8 | pytest | Install in build |
| Go | [go/ci.yml](go/ci.yml) | `ubuntu-latest` + Go 1.21 | go vet, staticcheck | go test | — |
| .NET | [dotnet/ci.yml](dotnet/ci.yml) | `ubuntu-latest` + .NET 8 | dotnet format | xUnit/NUnit | Restore, Publish |
| Ruby | [ruby/ci.yml](ruby/ci.yml) | `ubuntu-latest` + Ruby 3.3 | RuboCop | RSpec | Install in build |
| Rust | [rust/ci.yml](rust/ci.yml) | `ubuntu-latest` + Rust stable | clippy, rustfmt | cargo test | — |
| PHP | [php/ci.yml](php/ci.yml) | `ubuntu-latest` + PHP 8.2 | phpcs, phpstan | PHPUnit | Install in build |

## Usage

1. Copy the desired `ci.yml` to `.github/workflows/` in your repo (e.g., `ci.yml` or `java-ci.yml`)
2. Configure required secrets in GitHub repo settings: `DOCKER_USERNAME`, `DOCKER_PASSWORD`
3. Update `DOCKER_IMAGE` in the workflow to your registry path
4. Push to trigger the workflow

## Resources

- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [Workflow Syntax](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions)
