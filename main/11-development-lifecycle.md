# Dev tools and special cases

## Skaffold

Skaffold is a command line tool that facilitates continuous development for Kubernetes applications. 
Skaffold handles the workflow for building, pushing and deploying your application. 
It also provides building blocks and describe customizations for a CI/CD pipeline.

It automatically rebuilds your containers as you develop locally, and deploys those changes to either a local or remote cluster.

## Helm Hooks

Helm hooks allow to control the order in which things happen during a deployment. They also let you bail out of an upgrade if things go wrong.

# Key notions

* The default `RollingUpdate` deployment strategy in Kubernetes upgrades a few Pods at a time, waiting for each replacement  
  Pod to become ready before shutting down the old one.
* Rolling updates avoid downtime at the expense of making the rollout take longer. It also means that both old and new versions of your application will be running
simultaneously during the rollout period.
* `maxSurge` and `maxUnavailable` fields are to fine tune rolling updates. Depending on the versions of the Kubernetes API you are using, the
defaults may or may not be appropriate for your situation.
* The `Recreate` strategy deletes all the old Pods and starts up new ones all at once. 
* In a blue/green deployment, all the new Pods are started up and made ready without receiving any user traffic. 
  Then all traffic is switched over to the new Pods in one go, before retiring the old Pods.
* Rainbow deployments are similar to blue/green deployments, but with more than two versions in service simultaneously.
* You can implement blue/green and rainbow deployments in Kubernetes by adjusting the labels on your Pods and changing the selector on the frontend
Service to direct traffic to the appropriate set of Pods. 
* Helm hooks provide a way to apply certain Kubernetes resources (usually Jobs) at a particular stage of a deployment, 
 fe. to run a database migration. Hooks can define the order in which resources should be applied during a deployment, and cause the deployment to halt if something does not succeed.
* Knative and OpenFaaS allow to run “serverless” functions on clusters


# Best practice

# Annotated resources

* [skaffold](https://github.com/GoogleContainerTools/skaffold)
* [openfaas](https://www.openfaas.com)