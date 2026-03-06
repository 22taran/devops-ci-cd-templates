# 🐙 GitHub Actions Workflows

GitHub-native CI/CD workflow configurations for 8 tech stacks.

## Prerequisites

- GitHub repository with Actions enabled
- Repository secrets: `DOCKER_USERNAME`, `DOCKER_PASSWORD`
- Optional: `DEPLOY_TOKEN` for deployment

## Workflow Structure

Each `ci.yml` workflow includes:
- **Triggers**: Push to `main`, pull requests
- **Caching**: Language-specific dependency caching
- **5 jobs**: Build → Lint → Test → Docker → Deploy
- **Artifact uploads** for build outputs and coverage

## Tech Stacks

| Stack | File | Runner |
|-------|------|--------|
| Java | [java/ci.yml](java/ci.yml) | `ubuntu-latest` + JDK 17 |
| Node.js | [nodejs/ci.yml](nodejs/ci.yml) | `ubuntu-latest` + Node 18 |
| Python | [python/ci.yml](python/ci.yml) | `ubuntu-latest` + Python 3.12 |
| Go | [go/ci.yml](go/ci.yml) | `ubuntu-latest` + Go 1.21 |
| .NET | [dotnet/ci.yml](dotnet/ci.yml) | `ubuntu-latest` + .NET 8 |
| Ruby | [ruby/ci.yml](ruby/ci.yml) | `ubuntu-latest` + Ruby 3.3 |
| Rust | [rust/ci.yml](rust/ci.yml) | `ubuntu-latest` + Rust stable |
| PHP | [php/ci.yml](php/ci.yml) | `ubuntu-latest` + PHP 8.2 |

## Usage

1. Copy the desired `ci.yml` to `.github/workflows/` in your repo
2. Configure required secrets in GitHub repo settings
3. Push to trigger the workflow

## Resources

- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [Workflow Syntax](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions)
