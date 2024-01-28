# CI/CD

Continuous deployment (CD) is the automatic deployment of successful builds to production.

A typical pipeline for containerized applications:
1. Changes in code are pushed into repository.
2. The auto-build system re-builds the current version of the code and runs tests.
3. If all tests pass, the container image is published into the container registry.
4. The newly built container is deployed to a staging environment for UAT tests
5. The staging environment undergoes some automated and manual acceptance/e2e tests
6. The verified container image is auto-deployed to production (there might be a manual approval step - human in the loop)

## CI/CD Tools:

## Hosted CI/CD Tools
* Google Cloud Build
* Codefresh
* GitHub Actions
* GitLab CI

## Self-Hosted CI/CD Tools

* Argo
* Concourse

# Key notions

* Deciding which CI/CD tools to use is an important process when building a new pipeline. 
  All tools can be incorporated into almost any existing environment.
* Jenkins, GitHub Actions, GitLab, Drone, Cloud Build, and Spinnaker are popular CI/CD tools that work well with K8s.
* Defining the build pipeline steps with code allows you to track and modify these steps alongside application code.
* Containers enable developers to promote build artifacts up through envs, such as testing, staging, and eventually production
* GitOps is a newer term used when talking about CI/CD pipelines. The main idea is that deployments and infrastructure changes should be managed using code
that is tracked in source control (Git).
* Flux and Argo have popular GitOps tools that can automatically apply changes to your clusters whenever you push code changes to a Git repo.

# Best practice

* Use automation and version tracking wherever possible

# Annotated resources

* [GitHub Actions](https://github.com/features/actions)
* [ArgoCD](https://argo-cd.readthedocs.io/en/stable/)