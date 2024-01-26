# Cluster setup

# Key notions

* Managing your own clusters requires investment of time, effort, and expertise. Managed services like GKE, EKS, AKS, 
and other similar offerings do all the heavy lifting for you, at much lower cost than self-hosting.
* If you have to host your own cluster, `kops` and `Kubespray` are mature and widely used tools that can provision 
and manage production-grade clusters on AWS / Google Cloud.
* `VMware Tanzu` and `Google Anthos` make it possible to centrally manage Kubernetes clusters running in 
multiple clouds and on-premise infrastructure.

# Best practice

## Kubernetes Installers

_kops_  - is a command-line tool for automated provisioning of Kubernetes clusters, supports building 
high-availability clusters, which makes it suitable for production k8s deployments. 

_kubespray_ (aka Kargo) - is a tool for easily deploying production-ready clusters. 
It offers lots of options, including high availability, and support for multiple platforms.

_kubeadm_ is part of the Kubernetes distribution, and it aims to help you install and maintain a Kubernetes 
cluster according to best practices.


# Annotated resources

* [kops](https://kops.sigs.k8s.io/)
* [kubespray](https://github.com/kubernetes-sigs/kubespray)
* [kubeadm](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/)