# Sample Apps

Minimal application stubs for each tech stack. Use these to test CI/CD pipelines or as reference for expected project structure.

## Overview

Each sample app includes:
- **Source code** and tests that pass
- **Dockerfile** for containerization
- **Config files** expected by pipelines (pom.xml, package.json, requirements.txt, etc.)

## Tech Stacks

| Stack | Path | Build Command | Test Command | Lint Tool |
|-------|------|---------------|--------------|-----------|
| Java | [java/](java/) | `mvn clean package` | `mvn test` | Checkstyle |
| Node.js | [nodejs/](nodejs/) | `npm ci` | `npm test` | ESLint |
| Python | [python/](python/) | `pip install -r requirements.txt` | `pytest` | flake8 |
| Go | [go/](go/) | `go build` | `go test` | go vet |
| .NET | [dotnet/](dotnet/) | `dotnet build` | `dotnet test` | dotnet format |
| Ruby | [ruby/](ruby/) | `bundle install` | `bundle exec rspec` | RuboCop |
| Rust | [rust/](rust/) | `cargo build` | `cargo test` | clippy, rustfmt |
| PHP | [php/](php/) | `composer install` | `phpunit` | phpcs, phpstan |

## Try It: Run a Pipeline Locally

### Option 1: Jenkins (with Docker)

1. Copy a Jenkinsfile to a sample app:
   ```bash
   cp jenkins/java/Jenkinsfile sample-apps/java/
   cd sample-apps/java
   ```
2. Run with Jenkins or use the Jenkinsfile as reference for your project.

### Option 2: GitHub Actions

1. Create a repo with a sample app:
   ```bash
   cp -r sample-apps/java my-java-demo
   cd my-java-demo
   git init && git add . && git commit -m "Initial commit"
   ```
2. Copy `github-actions/java/ci.yml` to `.github/workflows/ci.yml`
3. Push to GitHub and watch the workflow run.

### Option 3: Local Build & Test (No CI)

Verify a sample app builds and tests locally:

```bash
# Java
cd sample-apps/java && mvn clean test

# Node.js
cd sample-apps/nodejs && npm ci && npm test

# Python
cd sample-apps/python && pip install -r requirements.txt && pytest

# Go
cd sample-apps/go && go build ./... && go test ./...
```

## Project Structure per Stack

- **Java**: `pom.xml`, `src/main/`, `src/test/`, `Dockerfile`
- **Node.js**: `package.json`, `package-lock.json`, `src/`, `test/`, `Dockerfile`
- **Python**: `requirements.txt`, `app.py`, `tests/`, `Dockerfile`
- **Go**: `go.mod`, `main.go`, `main_test.go`, `Dockerfile`
- **.NET**: `*.csproj`, `Program.cs`, `Tests.cs`, `Dockerfile`
- **Ruby**: `Gemfile`, `Gemfile.lock`, `app.rb`, `spec/`, `Dockerfile`
- **Rust**: `Cargo.toml`, `src/main.rs`, `Dockerfile`
- **PHP**: `composer.json`, `src/`, `tests/`, `Dockerfile`

## Customization

When using these as a starting point for your project:
1. Update package names, app names, and metadata
2. Add your dependencies and business logic
3. Ensure lint/test configs match what your pipeline expects (e.g., `checkstyle.xml`, `.eslintrc`)
4. Update the Dockerfile `EXPOSE` and `ENTRYPOINT` for your app
