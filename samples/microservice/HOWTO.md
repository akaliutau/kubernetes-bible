# Building a simple Go microservice

```shell
sudo docker image build -t microservice .
```

After build the image should be in local cache:

```shell
sudo docker images
REPOSITORY                TAG                    IMAGE ID       CREATED          SIZE
microservice           latest                   94b3c3104f47  28 seconds ago    6.07MB
```

Running app:

```shell
sudo docker container run -p 9999:8888 --name minimal microservice:latest
```

After navigating to `http://localhost:9999/` the app should show the string "Hello, world"

Publish image to docker hub:

```shell
sudo docker tag microservice:latest akaliutau/microservice
sudo docker login
sudo docker image push akaliutau/microservice
```

Running app at local K8s cluster:

```shell
minikube start
kubectl run goapp --image=akaliutau/microservice --port=9999 --labels app=goapp
kubectl get pods --selector app=goapp
kubectl port-forward pod/goapp 9999:8888
```

Clean up:

```shell
kubectl delete pod goapp
```