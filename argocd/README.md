# 🔄 ArgoCD Application Manifests

ArgoCD GitOps application configurations for 8 tech stacks.

## Prerequisites
- Kubernetes cluster with ArgoCD installed
- Container registry with built images
- Git repository accessible by ArgoCD

## What is ArgoCD?
ArgoCD is a declarative, GitOps continuous delivery tool for Kubernetes. It monitors Git repositories for changes and automatically syncs the desired state to the cluster.

## Tech Stacks

| Stack | File |
|-------|------|
| Java | [java/application.yaml](java/application.yaml) |
| Node.js | [nodejs/application.yaml](nodejs/application.yaml) |
| Python | [python/application.yaml](python/application.yaml) |
| Go | [go/application.yaml](go/application.yaml) |
| .NET | [dotnet/application.yaml](dotnet/application.yaml) |
| Ruby | [ruby/application.yaml](ruby/application.yaml) |
| Rust | [rust/application.yaml](rust/application.yaml) |
| PHP | [php/application.yaml](php/application.yaml) |

## Usage
1. Install ArgoCD in your cluster
2. Apply the desired `application.yaml`: `kubectl apply -f application.yaml`
3. Update the `repoURL` and `targetRevision` in the manifest

## Resources
- [ArgoCD Documentation](https://argo-cd.readthedocs.io/)
- [ArgoCD Getting Started](https://argo-cd.readthedocs.io/en/stable/getting_started/)
