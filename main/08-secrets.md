# Secrets and Environments

Example:

```shell
apiVersion: v1
kind: ConfigMap
metadata:
  name: demo-config
data:
  config.yaml: |
    autoSaveInterval: 60
    batchSize: 128
    protocols:
      - http
      - https
```
Deploying app with env config vars:

```shell
kubectl apply -f samples/configmaps/
kubectl describe configmap/goapp-config
```
# Key notions

* Separate your configuration data from application code and deploy it using `ConfigMaps` and `Secrets`. 
  That way, you don’t need to redeploy your app every time you change a password.
* One can get data into ConfigMaps by writing it directly in your Kubernetes manifest file, or use `kubectl` to convert 
  an existing YAML file into a ConfigMap spec.
* Once data is in a ConfigMap, one can insert it into a container’s environment, or into the command-line arguments of its entrypoint. 
  Alternatively, you can write the data to a file that is mounted on the container.
* Secrets work just like ConfigMaps, except that the data is encrypted at rest, and obfuscated in kubectl output.
* A simple way to manage secrets is to store them directly in your source code repo, but encrypt them using `Sops` or 
  another text-based encryption tool.
* Dedicated secret management tools like Vault, or hosted cloud KMS tools offer better auditing and flexibility for securing secrets.
* `Sops` is an encryption tool that works with key-value files like YAML and JSON. It can get its encryption key from a local keyring, 
  or cloud key management services like Amazon KMS and Google Cloud KMS.
* Sealed Secrets makes it easy to store encrypted secrets in source control and securely pass them to your applications from within a Kubernetes cluster.

# Best practice

Updating Pods on a Config Change

Add this annotation to your Deployment spec:
```shell
checksum/config: {{ include (print $.Template.BasePath "/configmap.yaml") . | sha256sum }}
```

Deployment template now includes a hash sum of the config settings. When you run `helm upgrade`, Helm will detect 
that the Deployment spec has changed, and restart all the Pods.

# Annotated resources