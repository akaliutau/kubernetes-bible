# Kubernetes Operations

# Key notions

* Before provisioning prod-level K8s cluster, plan the size (how many nodes is needed and of what size)
* At least 3 control plane nodes and at least 2 (ideally 3) worker nodes are needed. 
* Kubernetes clusters can scale to many thousands of nodes and up to 150,000 containers.
* To scale beyond that, use multiple clusters (sometimes this is needed also due to security/compliance reasons). 
* You can scale your cluster up and down manually without too much trouble, and
* There’s a well-defined standard for Kubernetes vendors and products: the CNCF Certified Kubernetes mark.
* Chaos testing is a process of knocking out Pods at random and seeing if your application still works. 

# Best practice

* K8s cluster needs at least 3 nodes running the control plane components in order to be HA.
  Two worker nodes is the abs min to be a fault-tolerant to the failure of a single node
* Max size of cluster - K8s v.1.22 officially supports clusters of up to 5,000 nodes
* Use _federation_ for workloads to be replicated across clusters (is not used with newer ver of K8s)

# Annotated resources

* [Cluster federation](https://github.com/kubernetes-retired/kubefed)
* [Successor project]()