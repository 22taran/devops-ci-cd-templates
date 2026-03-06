# 🔄 Travis CI Pipelines

Travis CI configurations for 8 tech stacks.

## Prerequisites
- Travis CI account connected to GitHub
- Repository enabled in Travis CI settings
- Environment variables: `DOCKER_USERNAME`, `DOCKER_PASSWORD`

## Tech Stacks

| Stack | File |
|-------|------|
| Java | [java/.travis.yml](java/.travis.yml) |
| Node.js | [nodejs/.travis.yml](nodejs/.travis.yml) |
| Python | [python/.travis.yml](python/.travis.yml) |
| Go | [go/.travis.yml](go/.travis.yml) |
| .NET | [dotnet/.travis.yml](dotnet/.travis.yml) |
| Ruby | [ruby/.travis.yml](ruby/.travis.yml) |
| Rust | [rust/.travis.yml](rust/.travis.yml) |
| PHP | [php/.travis.yml](php/.travis.yml) |

## Usage
1. Copy the desired `.travis.yml` to your project root
2. Enable the repo in Travis CI
3. Push to trigger

## Resources
- [Travis CI Documentation](https://docs.travis-ci.com/)
