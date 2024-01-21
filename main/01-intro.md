# K8s fundamentals

# Key notions

* Kubernetes cluster  = control plane + worker nodes with workloads.
* Production clusters must be HA => failure of a control plane node won’t lose data or affect the operation of the cluster.
* Managed services like GKE, EKS, AKS, and other similar offerings do all the heavy DevOps lifting
* _kops_ and _Kubespray_ are mature and widely used tools that can provision and manage production-grade clusters on AWS and GCP

## Definitions

### The control plane:

**kube-apiserver** - the frontend server for the control plane, handling API requests.

**etcd** - the database where Kubernetes stores all its information: what nodes exist, what resources exist on the cluster, and so on.

**kube-scheduler** - decides where to run newly created Pods.

**kube-controller-manager** - used for running resource controllers, such as Deployments.

**cloud-controller-manager** - used with the cloud provider (in cloud-based clusters), managing resources such as load balancers and disk volumes.

The control-plane components in a production cluster need to provide high availability (HA)

### Worker nodes.

Each worker node in a Kubernetes cluster runs these components:

**kubelet** - responsible for driving the container runtime to start workloads that are scheduled on the node, and monitoring their status.

**kube-proxy** - to route requests between Pods on different nodes, and between Pods and the internet.

**Container runtime** - starts and stops containers and handles their communications.

Historically the most popular option has been Docker, but Kubernetes supports other container runtimes as well, such as containerd and CRI-O.

## Check-list for self-hosted cluster

* Is the control plane highly available? 
* if any node goes down or becomes unresponsive, does your cluster still work? 
* Will running applications still be fault-tolerant without the control plane?
* Is pool of worker nodes highly available? 
* If an outage should take down several worker nodes, or even a whole cloud availability zone, will your workloads stop running? 
* Will it be able to automatically provision new nodes to heal itself, or will it require manual intervention?
* Is cluster set up securely? Do its internal components communicate using TLS encryption and trusted certificates? 
* Do users and applications have minimal rights and permissions for cluster operations? 
* Are container security defaults set properly? 
* Do nodes have unnecessary access to control plane components? 
* Is access to the underlying _etcd_ database properly controlled and authenticated?
* Are cluster nodes fully config-managed, rather than being set up by imperative shell scripts and then left alone? 
* Does the operating system and kernel on each node needs to be updated, have security patches applied, etc?
* Is the data in cluster properly backed up, including any persistent storage?
* What is backup process? How often it takes place?
* How cluster is maintained over time? How are new nodes provisioned? 
* How is Kubernetes updates rolled out? 

## Kubernetes Installers

_kops_ - a command-line tool for automated provisioning of Kubernetes clusters (part of the Kubernetes project)

_kubespray_ - (aka Kargo) a project under the Kubernetes umbrella, is a tool for easily deploying production-ready clusters.

# Best practice

* Don’t self-host k8s cluster without sound business reasons

# Annotated resources

* [Containerd](https://containerd.io)
* [Kubernetes the Hard Way](https://github.com/kelseyhightower/kubernetes-the-hard-way)
* [kubespray](https://github.com/kubernetes-sigs/kubespray)