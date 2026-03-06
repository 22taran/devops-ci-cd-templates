# 🦊 GitLab CI Pipelines

GitLab CI/CD pipeline configurations for 8 tech stacks.

## Prerequisites

- GitLab instance or GitLab.com account
- Container registry access for Docker builds
- CI/CD variables: `CI_REGISTRY`, `CI_REGISTRY_USER`, `CI_REGISTRY_PASSWORD`

## Tech Stacks

| Stack | File |
|-------|------|
| Java | [java/.gitlab-ci.yml](java/.gitlab-ci.yml) |
| Node.js | [nodejs/.gitlab-ci.yml](nodejs/.gitlab-ci.yml) |
| Python | [python/.gitlab-ci.yml](python/.gitlab-ci.yml) |
| Go | [go/.gitlab-ci.yml](go/.gitlab-ci.yml) |
| .NET | [dotnet/.gitlab-ci.yml](dotnet/.gitlab-ci.yml) |
| Ruby | [ruby/.gitlab-ci.yml](ruby/.gitlab-ci.yml) |
| Rust | [rust/.gitlab-ci.yml](rust/.gitlab-ci.yml) |
| PHP | [php/.gitlab-ci.yml](php/.gitlab-ci.yml) |

## Usage

1. Copy the desired `.gitlab-ci.yml` to your project root
2. Update variables and Docker registry settings
3. Push to trigger the pipeline

## Resources

- [GitLab CI/CD Documentation](https://docs.gitlab.com/ee/ci/)
- [Pipeline Architecture](https://docs.gitlab.com/ee/ci/pipelines/pipeline_architectures.html)
