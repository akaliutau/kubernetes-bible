# kubectl

# Key notions

* `kubectl` includes complete and exhaustive documentation on itself, available with `kubectl -h`
* The --dry-run=client option to `kubectl`, combined with -o YAML to get YAML output, can be used
  to generate Kubernetes manifests.
* You can turn existing resources into YAML manifests, too, using the -o flag to `kubectl get`
* `kubectl diff` will tell you what would change if you applied a manifest, without actually changing it.
* You can see the output and error messages from any container with `kubectl logs` , stream them continuously 
  with the `--follow` flag, or do more sophisticated multi-Pod log tailing with Stern.
* To troubleshoot problem containers, you can attach to them with `kubectl attach` 
  or get a shell on the container with `kubectl exec -it ... -- /bin/sh`
* `Click` is a powerful Kubernetes shell that gives all the functionality of `kubectl`, but with added state: 
  it remembers the currently selected object from one command to the next, so you don’t have to specify it every time.
* `Lens` is a standalone application one can use to manage Kubernetes clusters. `VSCode` also has a Kubernetes extension
* Kubernetes is designed to be automated and controlled by code, the Kubernetes client-go library gives complete 
  control over every aspect of your cluster using Go code.

# Best practice

## Use Short Flags

For example, use `-n` vs `--namespace`:

```shell
`kubectl` get pods -n kube-system
`kubectl` get po
`kubectl` get deploy
`kubectl` get svc
`kubectl` get ns
```

Use `jq` for json data extraction:

```shell
`kubectl` get pods -n kube-system -o json | jq '.items[].metadata.name'
```

## Generate Resource Manifests

Use `--dry-run=client` flag to not actually create the resource, but to print the metadata descriptor

```shell
`kubectl` create deployment goapp --image=akaliutau/microservice --dry-run=client -o yaml > deployment.yaml
```

## Forward a Container Port

```shell
`kubectl` port-forward goapp-54f4458547-vm88z 9090:8888
```

# Annotated resources