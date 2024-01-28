# Observability and Metrics

Some definitions

A _metric_ is a numerical measure of some characteristic of the system. Depending on the application, relevant metrics might include:
* The number of requests currently being processed
* The number of requests handled per minute (or per sec/hour)
* The number of errors encountered when handling requests
* The average time it took to serve requests (or the peak time, or the 99th percentile)

Infrastructure-related metrics:
* The CPU usage of individual processes or containers
* The disk I/O activity of nodes and servers
* The inbound and outbound network traffic of machines, clusters or load balancers

_Tracing_ follows a single user request through its whole life cycle

The _observability_ of system is a measure of how well instrumented it is, and how easily one can find out what’s going on inside it.

## Services metrics: The RED Pattern

* Requests _received_ per second
* Percentage of requests that returned an _error_
* _Duration_ of requests (also known as latency)

## Resources metrics: The USE Pattern

* Utilization == he average time that the resource was busy serving requests, or the amount of resource capacity that’s currently in use. 
* Saturation == the extent to which the resource is overloaded, or the length of the queue of requests waiting for this resource to become available. 
* Errors (aka Error Rate) == the number of times an operation on that resource failed. 

# Key notions

* Closed-box monitoring checks observe the external behavior of a system to detect predictable failures.
* Distributed systems expose the limitations of traditional monitoring because they’re not in either up or down states: 
  they exist in a constant state of partially degraded service. 
* Logs can be useful for post-incident troubleshooting, but are not scalable. Loki and ELK are popular centralized logging tools.
* Metrics open up a new dimension beyond simply working/not working, and give a continuous numerical time-series data on hundreds or thousands of aspects
of system. Prometheus is a popular option for aggregating application metrics.
* Metrics can help you answer the “why?” question, as well as identify problematic trends before they lead to outages.
* Tracing records events with precise timing through the life cycle of an individual request to help you debug performance problems. 
  Jaeger, Zipkin, and Pixie are some examples of tools that offer application tracing.
* Observability is the union of traditional monitoring, logging, metrics, and tracing to get a 360 overview of the system.
* Observability also represents a shift toward a team culture of engineering based on facts and feedback.
* Useful Kubernetes metrics include, at the cluster level, the number of nodes, Pods per node, and resource usage of nodes.
* At the deployment level, track deployments and replicas, especially unavailable replicas, which might indicate a capacity problem.
* At the container level, track resource usage per container, liveness/readiness states, restarts, network traffic, and network errors.

# Best practice

* Focus on the key metrics for each service: requests, errors, and duration (RED). For each resource: utilization, saturation, and errors (USE).
* Instrument apps to expose custom metrics, both for internal observability and for business KPIs.
* Build a dashboard for each service, using a standard layout and a primary information radiator that reports the vital signs of the whole system.

# Annotated resources

* [zipkin](https://zipkin.io/)
* [jaeger](https://www.jaegertracing.io/)