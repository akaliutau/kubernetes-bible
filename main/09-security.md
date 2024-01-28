# Security and RBAC

Tools to scan clusters against any issues:

## Gatekeeper/OPA

De-facto is a set of CRDs to define and set extra rules/constraints to make cluster more stable and secure, fe.:

```shell
apiVersion: constraints.gatekeeper.sh/v1beta1
kind: K8sDisallowedTags
metadata:
  name: container-image-must-not-have-latest-tag
spec:
  match:
    kinds:
      - apiGroups: [""]
        kinds: ["Pod"]
    namespaces:
      - "my-namespace"
  parameters:
    tags: ["latest"]
```

## kube-bench

This is a tool for auditing K8s cluster against a set of benchmarks produced by the Center for Internet Security (CIS). 
In effect, it verifies that the cluster is set up according to security best practices.

# Key notions

* RBAC gives you fine-grained management of permissions in Kubernetes. Make sure it’s enabled, 
  and use RBAC roles to grant specific users and apps only the minimum privileges they need to do their jobs.
* Containers are not exempt from security and malware problems. Use a scanning tool to check any containers that you run in production.
* Using Kubernetes does not mean that you don’t need backups: back up data and the state of the cluster.
  (can be used for moving things between clusters as a spin-off/bonus)
* kubectl is a powerful tool for inspecting and reporting on all aspects of cluster and its workloads. 
* Kubernetes provides web console and `kube-ops-view` for a graphical overview of what’s going on in cluster

# Best practice

* Don’t run containers from untrusted sources. Run a scanning tool like Clair or Synk over all containers, 
  especially those you build yourself, to make sure there are no known vulnerabilities in any of the base images or deps.
* Setup a nightly backup (for restoring cluster state and persistent data) with velero. Run a restore test at least monthly.

# Annotated resources

* [gatekeeper](https://github.com/open-policy-agent/gatekeeper)
* [kube-bench](https://github.com/aquasecurity/kube-bench)
* [velero](https://velero.io)