# 🚀 DevOps CI/CD Pipeline Reference

<div align="center">

![CI/CD](https://img.shields.io/badge/CI%2FCD-Pipeline%20Reference-blue?style=for-the-badge&logo=githubactions)
![Tools](https://img.shields.io/badge/Tools-10-green?style=for-the-badge)
![Tech Stacks](https://img.shields.io/badge/Tech%20Stacks-8-orange?style=for-the-badge)
![Pipelines](https://img.shields.io/badge/Pipelines-80%2B-red?style=for-the-badge)
![License](https://img.shields.io/badge/License-MIT-yellow?style=for-the-badge)

**A comprehensive collection of CI/CD pipeline configurations across 10 popular DevOps tools and 8 tech stacks.**

*Production-ready templates with build, lint, test, Docker, and deploy stages.*

[Jenkins](#jenkins) · [GitHub Actions](#github-actions) · [GitLab CI](#gitlab-ci) · [CircleCI](#circleci) · [Travis CI](#travis-ci) · [Azure DevOps](#azure-devops) · [Bitbucket Pipelines](#bitbucket-pipelines) · [ArgoCD](#argocd) · [Tekton](#tekton) · [AWS CodePipeline](#aws-codepipeline)

</div>

---

## 📋 Coverage Matrix

| Tech Stack | Jenkins | GitHub Actions | GitLab CI | CircleCI | Travis CI | Azure DevOps | Bitbucket | ArgoCD | Tekton | AWS CodePipeline |
|:----------:|:-------:|:--------------:|:---------:|:--------:|:---------:|:------------:|:---------:|:------:|:------:|:----------------:|
| ☕ Java     | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| 🟢 Node.js | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| 🐍 Python  | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| 🔵 Go      | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| 🟣 .NET    | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| 💎 Ruby    | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| 🦀 Rust    | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |
| 🐘 PHP     | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ | ✅ |

---

## 📂 Repository Structure

```
├── sample-apps/              # Minimal app stubs per tech stack
│   ├── java/
│   ├── nodejs/
│   ├── python/
│   ├── go/
│   ├── dotnet/
│   ├── ruby/
│   ├── rust/
│   └── php/
│
├── jenkins/                  # Declarative Jenkinsfiles
├── github-actions/           # GitHub Actions YAML workflows
├── gitlab-ci/                # GitLab CI/CD YAML configs
├── circleci/                 # CircleCI config.yml files
├── travis-ci/                # Travis CI .travis.yml files
├── azure-devops/             # Azure Pipelines YAML
├── bitbucket-pipelines/      # Bitbucket pipeline configs
├── argocd/                   # ArgoCD Application manifests
├── tekton/                   # Tekton Pipeline & Task CRDs
└── aws-codepipeline/         # AWS CodeBuild buildspec files
```

---

## 🔄 Pipeline Stages

Every pipeline follows a consistent **5-stage pattern**:

```
┌─────────┐    ┌─────────┐    ┌─────────┐    ┌─────────┐    ┌─────────┐
│  Build  │───▶│  Lint   │───▶│  Test   │───▶│ Docker  │───▶│ Deploy  │
└─────────┘    └─────────┘    └─────────┘    └─────────┘    └─────────┘
```

| Stage    | Description                                     |
|----------|-------------------------------------------------|
| **Build**   | Compile source code / install dependencies     |
| **Lint**    | Static code analysis & style checks            |
| **Test**    | Run unit tests with coverage reports           |
| **Docker**  | Build & push Docker image                      |
| **Deploy**  | Deploy to staging environment                  |

---

## 🛠️ Tools Overview

### Jenkins
Industry-standard automation server with declarative pipeline syntax. Each tech stack has a `Jenkinsfile` with Docker agent, stages, and post-build actions.

### GitHub Actions
GitHub-native CI/CD with YAML workflows. Includes caching, matrix builds, and artifact uploads.

### GitLab CI
GitLab-integrated pipelines with `.gitlab-ci.yml`. Features Docker-in-Docker, artifacts, and environment deployments.

### CircleCI
Cloud-native CI/CD with orbs, executors, and workflow orchestration.

### Travis CI
Simple CI/CD with `.travis.yml` configuration. Great for open-source projects.

### Azure DevOps
Microsoft's DevOps platform with `azure-pipelines.yml`. Supports multi-stage pipelines with approval gates.

### Bitbucket Pipelines
Atlassian's integrated CI/CD with `bitbucket-pipelines.yml`. Built-in Docker support and deployment environments.

### ArgoCD
Declarative GitOps continuous delivery for Kubernetes. Application manifests that sync with Git repositories.

### Tekton
Kubernetes-native CI/CD with Pipeline and Task CRDs. Cloud-native, vendor-neutral pipeline engine.

### AWS CodePipeline
AWS-native CI/CD with CodeBuild `buildspec.yml`. Integrates with ECR, ECS, Lambda, and other AWS services.

---

## 🚀 Quick Start

1. **Clone the repository:**
   ```bash
   git clone https://github.com/yourusername/DevOps-CI-CD.git
   cd DevOps-CI-CD
   ```

2. **Choose a tool and tech stack:**
   Navigate to the desired tool folder and tech stack subfolder.

3. **Copy the pipeline config** into your project and customize:
   - Update project-specific values (repo URL, Docker registry, etc.)
   - Adjust deployment targets and environment variables
   - Modify test and build commands as needed

4. **Reference the sample apps** in `sample-apps/` for the expected project structure.

---

## 📚 Resources

| Tool | Official Docs |
|------|---------------|
| Jenkins | [jenkins.io/doc](https://www.jenkins.io/doc/) |
| GitHub Actions | [docs.github.com/actions](https://docs.github.com/en/actions) |
| GitLab CI | [docs.gitlab.com/ee/ci](https://docs.gitlab.com/ee/ci/) |
| CircleCI | [circleci.com/docs](https://circleci.com/docs/) |
| Travis CI | [docs.travis-ci.com](https://docs.travis-ci.com/) |
| Azure DevOps | [learn.microsoft.com](https://learn.microsoft.com/en-us/azure/devops/pipelines/) |
| Bitbucket | [support.atlassian.com](https://support.atlassian.com/bitbucket-cloud/docs/get-started-with-bitbucket-pipelines/) |
| ArgoCD | [argo-cd.readthedocs.io](https://argo-cd.readthedocs.io/) |
| Tekton | [tekton.dev/docs](https://tekton.dev/docs/) |
| AWS CodePipeline | [docs.aws.amazon.com](https://docs.aws.amazon.com/codepipeline/) |

---

## 🤝 Contributing

Contributions are welcome! Please read the [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## 📄 License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.

---

<div align="center">

**⭐ Star this repo if you find it useful!**

Made with ❤️ for the DevOps community

</div>
