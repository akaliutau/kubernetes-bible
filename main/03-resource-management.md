# Resource management (Limits, Requests, QoS tuning)

# Key notions

* Kubernetes allocates CPU and memory resources to containers on the basis of requests and limits specified in manifests
* A container’s requests are the _minimum_ amounts of resources it needs to run; 
  the _limit_ specifies the maximum amount it’s allowed to use.
* Minimal container images are faster to build, push, deploy, and start (The smaller image - 
  the fewer the potential security vulnerabilities)
* Liveness probes tell Kubernetes whether the container is working properly
* Readiness probes tell Kubernetes that the container is ready and able to serve requests
  If the readiness probe fails, the container will be removed from any Services that reference it, disconnecting it from user traffic.
* Startup probes are like liveness probes, but are only used for determining if an application has finished starting.
* `PodDisruptionBudgets` let you limit the number of Pods that can be stopped at once during evictions, preserving HA for application.
* Namespaces are a way of logically partitioning your cluster (can create a namespace for each application, or group of related applications)
* To refer to a Service in another namespace, you can use a DNS address like this: `SERVICE.NAMESPACE.`
* `ResourceQuotas` let to set overall resource limits for a given namespace.
* `LimitRanges` specify default resource requests and limits for containers in a namespace.
* Don’t allocate more cloud storage than you need, and don’t provision high bandwidth storage unless it’s really needed.
* Set owner annotations on all your resources, and scan the cluster regularly for unowned resources.
* Find and clean up resources that aren’t being used
* Reserved instances can save you money in long run.
* _Preemptive_ instances can save you money right now, but be ready for them to vanish on short notice. 
  Use _node affinities_ to keep failure-sensitive Pods away from _preemptive_ nodes.

## Resource Requests
A resource request specifies the minimum amount of that resource that the Pod needs to run. 
For example, a resource request of 100m (100 millicpus) and 250Mi (250 MiB of memory) means the min amount of
resources needed to run the Pod (ie. it cannot be scheduled on a node with less than those resources available)

## Resource Limits
A resource limit specifies the maximum amount of resource that a Pod is allowed to use.

Example:

```shell
spec:
  containers:
  - name: goapp
    image: akaliutau/microservice
    ports:
    - containerPort: 8888
    resources:
      requests:
        memory: "10Mi"
        cpu: "100m"
```

## Quality of Service
Based on the requests and limits of a Pod, Kubernetes will classify it as one of the following Quality of Service (QoS) classes: 
* Guaranteed
* Burstable
* BestEffort

# Best practice

* Set Liveness and Readiness Probes
* Use namespaces to segregate and manage resource on the user by user or system basis
* Use Pod Disruption Budgets (`kind: PodDisruptionBudget`) to specify the max limit of failed pods in any moment of time
* Use `ResourceQuotas` in each namespace to enforce a limit on the number of Pods that can run in the namespace


# Annotated resources

