# 🪣 Bitbucket Pipelines

Bitbucket Pipelines configurations for 8 tech stacks.

## Prerequisites
- Bitbucket Cloud repository
- Repository variables: `DOCKER_USERNAME`, `DOCKER_PASSWORD`
- Deployment environments configured

## Tech Stacks

| Stack | File |
|-------|------|
| Java | [java/bitbucket-pipelines.yml](java/bitbucket-pipelines.yml) |
| Node.js | [nodejs/bitbucket-pipelines.yml](nodejs/bitbucket-pipelines.yml) |
| Python | [python/bitbucket-pipelines.yml](python/bitbucket-pipelines.yml) |
| Go | [go/bitbucket-pipelines.yml](go/bitbucket-pipelines.yml) |
| .NET | [dotnet/bitbucket-pipelines.yml](dotnet/bitbucket-pipelines.yml) |
| Ruby | [ruby/bitbucket-pipelines.yml](ruby/bitbucket-pipelines.yml) |
| Rust | [rust/bitbucket-pipelines.yml](rust/bitbucket-pipelines.yml) |
| PHP | [php/bitbucket-pipelines.yml](php/bitbucket-pipelines.yml) |

## Usage
1. Copy `bitbucket-pipelines.yml` to your project root
2. Configure repository variables
3. Push to trigger

## Resources
- [Bitbucket Pipelines Documentation](https://support.atlassian.com/bitbucket-cloud/docs/get-started-with-bitbucket-pipelines/)
