# Containers

# Key notions

* A Linux container is an isolated set of processes, with ring-fenced resources. From inside a container, it looks like a Linux instance
* Containers are _not virtual machines_. Each container should run one primary process.
* A Pod = one container that runs a primary application + optional helper containers.
* Container image specifications can include a registry hostname, a repository namespace, an image repository, and a tag
* Programs in containers should not run as the root user. Instead, assign them an ordinary user.
* Containers in the same Pod can share data by reading and writing a mounted Volume. 
  The simplest Volume is of type _emptyDir_, which starts out empty and preserves its contents only as long as the Pod is running.
* A _PersistentVolume_ preserves its contents as long as needed. Pods can dynamically provision 
  new _PersistentVolumes_ using _PersistentVolumeClaims_
* _Init containers_ can be useful for doing initial setup before your application is started in a Pod (fe. sync Git repo)

# Best practice

* Container Security
   * run containers as non-root users, and block root containers from running, using the `runAsNonRoot: true` setting
   * block containers from running if they would run as the root user using `runAsNonRoot: true` setting
   * prevent the container from writing to its own filesystem using `readOnlyRootFilesystem: true` setting
   * disable privilege escalation using `allowPrivilegeEscalation: false` setting
   * set _capabilities_ for containers via `capabilities.drop` and `capabilities.add`
   * create a dedicated service account for the app, bind it to the required roles, and configure the Pod to use this SA


# Annotated resources