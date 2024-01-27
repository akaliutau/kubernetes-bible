# Pods

# Key notions

* _Labels_ are key-value pairs that identify resources, and can be used with selectors to match a specified group of resources.
* Node _affinities_ attract or repel Pods to or from nodes with specified attributes (fe. to run a Pod on a node in a specified availability zone)
* _Hard node affinities_ = to block a Pod from running, _soft node affinities_ = suggestions to the scheduler. One can combine multiple soft affinities
with different weights.
* Pod affinities express a preference for Pods to be scheduled on the same node as other Pods
* Pod anti-affinities repel other Pods instead of attracting. For example, an anti-affinity to replicas of the same Pod 
  can help spread your replicas evenly across the cluster.
* _Taints_ are a way of tagging nodes with specific information, usually about node problems or failures (Pods won’t be scheduled on tainted nodes)
* _Tolerations_ allow a Pod to be scheduled on nodes with a specific taint. 
* _DaemonSets_ allow you to schedule one copy of a Pod on every node (fe, a logging agent).
* _StatefulSets_ start and stop Pod replicas in a specific numbered sequence, allowing you to address each by a predictable DNS name. 
* _Jobs_ run a Pod once (or a specified number of times) before completing (CronJobs run a Pod periodically at specified times)
* Horizontal Pod Autoscalers (HPAs) watch a set of Pods, trying to optimize a given metric (such as CPU utilization). 
  They increase or decrease the desired number of replicas to achieve the specified goal.
* Custom Resource Definitions (CRDs) allow to create your own custom Kubernetes objects, to store any data you wish. 
  Operators are Kubernetes client programs that can implement orchestration behavior for your specific application. 
* Ingress resources route requests to different services, depending on a set of rules, for example, matching parts of the request URL. 
  They can also terminate TLS connections for your application.
* _Istio_, _Linkerd_, and _Consul Connect_ are service mesh tools to provide networking features for microservice environments such as encryption, QoS,
metrics, logging, and more route strategies.

## Ingress Controllers comparison list

* _nginx-ingress_ - _nginx_ has long been a popular load balancer tool, even before Kubernetes came on the scene. 
The nginx-ingress project is maintained by the Kubernetes community.
* _Contour_ - is based on Envoy to proxy requests between clients and Pods.
* _Traefik_ is a lightweight proxy tool that can also automatically manage TLS certificates for your Ingress.
* _Kong_ hosts the Kong Plugin Hub with plugins that integrate with their Ingress controller to configure things like OAuth authentication, 
  LetsEncrypt certificates, IP restriction, metrics, and other useful features for load balancers.
* _HAProxy_ has been another popular tool for load balancers for several years, 
  and they also have their own Ingress controller for Kubernetes along with a Helm chart for installing it in your clusters.

## Service Mesh comparison list

Kubernetes Ingress and Services are for routing requests from clients to applications. A service mesh is responsible 
for managing more complex network operations such as rate-limiting and encrypting network traffic between microservices.

* _Istio_ - was one of the first tools associated with providing a service mesh. 
  It is available as an optional add-on component to many hosted Kubernetes clusters, including GKE. 
* _Linkerd_ offers many of the key service mesh features, but with a much lighter footprint and less complexity involved compared to Istio. 
* _Consul Connect_ - offered by HashiCorp wa, was focused on service discovery. 
  It handled application health checks and the automatic routing of requests to the right place in a distributed
compute environment. 

# Best practice


# Annotated resources

* [traefik](https://doc.traefik.io/traefik/providers/kubernetes-ingress/)
* [ingress-nginx](https://github.com/kubernetes/ingress-nginx)
* [contour](https://github.com/projectcontour/contour)
* [haproxy](https://haproxy-ingress.github.io/docs/)
* [kong](https://github.com/Kong/kubernetes-ingress-controller)
* [resource for searching community-built Operators](https://operatorHub.io)