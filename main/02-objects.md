# K8s base resources and management

```shell
kubectl create deployment goapp --image=akaliutau/microservice
kubectl get deployments
NAME         READY   UP-TO-DATE   AVAILABLE   AGE
goapp      1/1     1            1           2s
```

Describing deployments:

```shell
kubectl describe deployments/goapp
Name:                   goapp
Namespace:              default
CreationTimestamp:      Sun, 21 Jan 2024 22:54:46 +0000
Labels:                 app=goapp
Annotations:            deployment.kubernetes.io/revision: 1
Selector:               app=goapp
Replicas:               1 desired | 1 updated | 1 total | 1 available | 0 unavailable
StrategyType:           RollingUpdate
MinReadySeconds:        0
RollingUpdateStrategy:  25% max unavailable, 25% max surge
```

```shell
kubectl delete all --selector app=goapp
```

```shell
kubectl apply -f minimal/deployment.yaml
kubectl apply -f minimal/service.yaml
kubectl get pods --selector app=goapp
kubectl port-forward service/goapp 9999:8888
NAME                    READY   STATUS    RESTARTS   AGE
goapp-bdd7db547-4r9mj   1/1     Running   0          18s
```

```shell
kubectl get nodes
kubectl describe pod/goapp-bdd7db547-4r9mj
Name:             goapp-bdd7db547-4r9mj
Namespace:        default
Priority:         0
Service Account:  default
Node:             minikube/192.168.49.2
Start Time:       Sun, 21 Jan 2024 23:04:09 +0000
Labels:           app=goapp
                  pod-template-hash=bdd7db547
```

Installing demo goapp using helm (recommended approach due to rich configuration options):

```shell
helm upgrade --install goapp ./helm-goapp/
kubectl get deployment
kubectl get service
helm list
NAME 	NAMESPACE	REVISION	UPDATED                                	STATUS  	CHART      	APP VERSION
goapp	default  	1       	2024-01-21 23:18:16.592529516 +0000 UTC	deployed	goapp-1.0.0	
```

helm helps to keep configuration separately from k8s templates, which leads to better and simpler config management
(fe. for DEV/UAT/PROD envs)

# Key notions

* The _Pod_ is the fundamental unit of work in k8s == single container or group of communicating containers that are scheduled together.
* A _Deployment_ is a high-level Kubernetes resource that declaratively manages Pods: deploying, scheduling, updating.
* A _Service_ is the Kubernetes equivalent of a load balancer or proxy, routing traffic to its matching Pods via a single, well-known, durable IP address or DNS name.
* The _Kubernetes scheduler_ watches for a Pod that isn’t yet running on any node, finds a suitable node for it, and instructs the kubelet on that node to run the Pod.
* Deployments == records in Kubernetes’s internal database (etcd). Externally == manifest in YAML format
* _kubectl_ is the main tool for interacting with Kubernetes, allowing to apply manifests, query resources, make changes etc
* _helm_ is a Kubernetes package manager. It simplifies configuring and deploying Kubernetes applications

# Annotated resources

* [Helm official documentation](https://helm.sh/docs/intro/install/)