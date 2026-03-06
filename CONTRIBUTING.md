# Contributing to DevOps CI/CD Pipeline Reference

Thank you for your interest in contributing! This project aims to be the most comprehensive CI/CD reference on GitHub.

## How to Contribute

### Adding a New CI/CD Tool

1. Create a new folder at the root level with the tool name (lowercase, hyphenated)
2. Add a `README.md` explaining the tool
3. Create subfolders for each tech stack
4. Add pipeline config files following the 5-stage pattern (Build → Lint → Test → Docker → Deploy)

### Adding a New Tech Stack

1. Add a sample app in `sample-apps/<tech-stack>/`
2. Add pipeline configs for the new stack in **every** tool folder
3. Update the coverage matrix in the root `README.md`

### Improving Existing Pipelines

- Follow best practices for the specific tool
- Include caching where possible
- Use tool-specific features (e.g., GitHub Actions matrix, CircleCI orbs)
- Add comments explaining non-obvious configurations

## Pipeline Standards

Every pipeline should include these stages:

| Stage | Required |
|-------|----------|
| Build | ✅ |
| Lint | ✅ |
| Test | ✅ |
| Docker | ✅ |
| Deploy | ✅ |

## Pull Request Process

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/add-new-tool`)
3. Commit your changes (`git commit -m 'Add: New CI/CD tool configs'`)
4. Push to the branch (`git push origin feature/add-new-tool`)
5. Open a Pull Request with a clear description

## Code Style

- YAML: 2-space indentation
- Groovy (Jenkinsfile): 4-space indentation
- Use descriptive stage/job names
- Add inline comments for complex configurations
- Keep configs DRY — use variables and anchors where possible

## Questions?

Open an issue and we'll be happy to help!
