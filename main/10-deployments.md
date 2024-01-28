# Deployments

helm is the best tool for deployment and releasemanship  

For example, the following command deploys goapp to the cluster using staging env vars to generate manifests:

```shell
helm upgrade --install goapp-staging \
  --values=samples/helm-goapp/staging-values.yaml samples/helm-goapp
```

This will create a new release, with a new name ( demo-staging ), and the running container’s ENVIRONMENT 
variable will be set to staging instead of development. The extra values in file `--values` are combined with
those in the default values file (`samples/helm-goapp/values.yaml`). There’s only one variable -
environment - which will be merged.

Helm allows to setup a private repo for charts.
The charts need to be available over HTTP, and there’s a variety of ways you can do this: 

* cloud storage bucket
* in ChartMuseum server
* artifactory
* GitHub Pages
* existing web server

_Helmfile_ is a helm for helm charts: it enables to deploy everything that should be installed on your cluster,
with a single command:

```shell
helmfile sync
```

# Key notions

* A chart is a Helm package specification, including metadata about the package, some configuration values for it, 
and template Kubernetes objects that reference those values.
* Chart Installation == new Helm release. Release update with different config values == increment of release revision number.
* To customize a Helm chart for your own requirements, create a custom values file overriding just the settings 
* With Helmfile, one can declaratively specify a set of Helm charts and values to be applied to your cluster, and install or update
* Helm can be used along with Sops for handling secret configuration in charts. It can also use a function to automatically base64-encode your secrets,
which Kubernetes expects them to be.
* `Tanka` and `Kapitan` are alternative manifest management tools that use `Jsonnet`, a different templating language.
* Use `kubeval` to test and validate quickly manifests (will check for valid syntax and common errors in manifests)

# Best practice

* Use a single source of truth: don’t mix deploying individual charts manually with Helm, and declaratively managing all your charts
across the cluster with Helmfile or a GitOps tool.

# Annotated resources

* [chart-repository-guide](https://v2.helm.sh/docs/developing_charts/#the-chart-repository-guide)
* [kustomize](https://github.com/kubernetes-sigs/kustomize)