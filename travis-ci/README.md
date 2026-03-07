# 🔄 Travis CI Pipelines

Travis CI configurations for 8 tech stacks.

## Prerequisites

- Travis CI account connected to GitHub
- Repository enabled in Travis CI settings
- Environment variables: `DOCKER_USERNAME`, `DOCKER_PASSWORD`

## Pipeline Structure

Each `.travis.yml` uses Travis stages with:
- **Stages**: build → parallel (lint, test, security-code) → docker build → security (image) → docker push → deploy
- **Caching**: Maven, npm, pip, etc. for dependency speedup
- **Security**: Code scan (Trivy fs, Semgrep, Gitleaks) in parallel with Lint/Test; Image scan (Trivy, Snyk Container) after Docker Build

## CI/CD Pipeline Flow Diagram

<div align="center">

<pre>
```
┌─────────────────────────────────────────────────────────────────┐
│  Git Push / PR                                                  │
│  Travis CI: on: branches                                        │
└─────────────────────────────────────────────────────────────────┘
|
▼
┌─────────────────────────────────────────────────────────────────┐
│  Build                                                          │
│  stage: build | cache: maven, npm, pip | compile/install deps   │
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
│  stage: docker | docker build with tag                          │
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
│  stage: docker | docker login, push to registry                 │
└─────────────────────────────────────────────────────────────────┘
|
▼
┌─────────────────────────────────────────────────────────────────┐
│  Deploy                                                         │
│  stage: deploy | kubectl / helm / docker run                    │
└─────────────────────────────────────────────────────────────────┘
```
</pre>

</div>

## Stage-by-Stage Explanation

| Stage | Purpose | What Happens | Artifacts / Output |
|-------|---------|--------------|--------------------|
| **build** | Compile or install deps | Maven compile/package, npm ci, pip install, etc. Caches deps. | JAR, build artifacts |
| **lint** | Static analysis | Runs in parallel with test and security-code. checkstyle, ESLint, flake8, go vet, etc. | — |
| **test** | Unit tests | Runs in parallel. Some configs publish coverage. | — |
| **security-code** | SAST, deps, secrets | Runs in parallel. Trivy fs, Semgrep, Snyk Code, OWASP Dep-Check, Gitleaks. | Scan report |
| **docker (build)** | Build image | Docker build with tag. | Image in local daemon |
| **security (image)** | Container scan | After Docker Build. Trivy, Snyk Container, Clair. Scan image for CVEs. | Scan report |
| **docker (push)** | Push to registry | Docker login, push. Runs on every branch. | Image in registry |
| **deploy** | Deploy to staging | Runs on every branch. Supports Docker or Kubernetes. Replace echo with kubectl/Helm/docker run. | — |

## Security Tools (Industry Options)

| Category | Purpose | Tools |
|----------|---------|-------|
| **Security (Code)** | SAST, dependency scan, secret scan — runs in parallel with Lint/Test | [Trivy fs](https://trivy.dev/), [Semgrep](https://semgrep.dev/), [Snyk Code](https://snyk.io/product/snyk-code/), [OWASP Dependency-Check](https://owasp.org/www-project-dependency-check/), [Gitleaks](https://github.com/gitleaks/gitleaks), [TruffleHog](https://github.com/trufflesecurity/trufflehog) |
| **Security (Image)** | Container scan — runs after Docker Build, before Push | [Trivy](https://trivy.dev/), [Snyk Container](https://snyk.io/product/container-vulnerability-management/), [Clair](https://github.com/quay/clair), [Anchore](https://anchore.com/) |

## Tech Stacks

| Stack | File | Language / JDK | Lint Tool | Test Framework |
|-------|------|-----------------|-----------|----------------|
| Java | [java/.travis.yml](java/.travis.yml) | openjdk17 | Checkstyle | JUnit |
| Node.js | [nodejs/.travis.yml](nodejs/.travis.yml) | node_js 18 | ESLint | Jest/npm test |
| Python | [python/.travis.yml](python/.travis.yml) | python 3.12 | flake8 | pytest |
| Go | [go/.travis.yml](go/.travis.yml) | go 1.21 | go vet | go test |
| .NET | [dotnet/.travis.yml](dotnet/.travis.yml) | dotnet 8.0 | dotnet format | xUnit/NUnit |
| Ruby | [ruby/.travis.yml](ruby/.travis.yml) | ruby 3.3 | RuboCop | RSpec |
| Rust | [rust/.travis.yml](rust/.travis.yml) | rust stable | clippy, rustfmt | cargo test |
| PHP | [php/.travis.yml](php/.travis.yml) | php 8.2 | phpcs, phpstan | PHPUnit |

## Usage

1. Copy the desired `.travis.yml` to your project root
2. Enable the repo in Travis CI
3. Configure environment variables in Travis CI repo settings
4. Push to trigger

## Resources

- [Travis CI Documentation](https://docs.travis-ci.com/)
