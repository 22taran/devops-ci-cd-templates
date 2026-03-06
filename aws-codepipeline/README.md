# ☁️ AWS CodePipeline / CodeBuild

AWS CodeBuild buildspec configurations for 8 tech stacks.

## Prerequisites
- AWS Account with CodePipeline and CodeBuild
- ECR repository for Docker images
- IAM roles with appropriate permissions
- S3 bucket for artifacts (auto-created by CodePipeline)

## Tech Stacks

| Stack | File |
|-------|------|
| Java | [java/buildspec.yml](java/buildspec.yml) |
| Node.js | [nodejs/buildspec.yml](nodejs/buildspec.yml) |
| Python | [python/buildspec.yml](python/buildspec.yml) |
| Go | [go/buildspec.yml](go/buildspec.yml) |
| .NET | [dotnet/buildspec.yml](dotnet/buildspec.yml) |
| Ruby | [ruby/buildspec.yml](ruby/buildspec.yml) |
| Rust | [rust/buildspec.yml](rust/buildspec.yml) |
| PHP | [php/buildspec.yml](php/buildspec.yml) |

## Usage
1. Create a CodePipeline in AWS Console or via CloudFormation
2. Add CodeBuild as a build stage
3. Place `buildspec.yml` in your project root
4. Configure environment variables in CodeBuild project

## Resources
- [AWS CodePipeline Documentation](https://docs.aws.amazon.com/codepipeline/)
- [CodeBuild Buildspec Reference](https://docs.aws.amazon.com/codebuild/latest/userguide/build-spec-ref.html)
