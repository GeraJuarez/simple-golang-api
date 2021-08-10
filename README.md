# Simple golang API

Golang API to practice Clean Architecture, containers and Cloud Design Patterns

Using [manato's article: Clean Architecture with GO](https://medium.com/manato/clean-architecture-with-go-bce409427d31).

## TODO

* v3 to practice concurrecy and channels
* implement Cloud Native Patterns
* utests and mocks
    * https://pkg.go.dev/github.com/gorilla/mux#readme-testing-handlers
* Add file logger
* implement more patterns from Cloud Native Go book
* gRPC
* env file best practices
* use oauth
* add https security
* add db logger
* add db keystore
* document API
* input validation in all layers
* DB seeder

## Usefull commands

### GOlang

```bash
go mod init
go mod tidy
```

### Docker

```bash
docker build --tag kvstore key-store-api/ -f .
docker run --detach --publish 8080:8080 kvstore
```

### Kubernetes minikube

```bash
kubectl create -f backend-deployment.yaml --namespace development
kubectl create -f backend-service.yaml --namespace development
```

```bash
minikube version
minikube start
kubectl version
kubectl get # list resources
kubectl describe # show detailed information about a resource
kubectl logs # print the logs from a container in a pod
kubectl exec # execute a command on a container in a pod
```

```bash
# Deploy app
kubectl get nodes
kubectl create deployment <my-name> --image=<path>:<version>
kubectl get deployments
kubectl proxy # access deployemnts 
kubectl get pods
kubectl describe pods
kubectl logs <pod_name>
kubectl exec -ti <pod_name> -- bash

# Expose app
kubectl get services
kubectl expose deployment/example-service --type="NodePort" --port 8080 # can use curl <minikube ip>:<node_port> to test
kubectl describe deployment
kubectl get pods -l app=exmple-service
kubectl get services -l app=exmple-service
kubectl label pods <pod_name> version=v1
kubectl get pods -l version=v1
kubectl delete services -l app=exmple-service

# Scale app
kubectl get rs # ReplicaSet created by the Deployment
kubectl scale deployments/example-app --replicas=4
kubectl get pods -o wide
kubectl describe services/example-app
export NODE_PORT=$(kubectl get services/example-app -o go-template='{{(index .spec.ports 0).nodePort'}})
curl $(minikube ip):$NODE_PORT # this should ping a different pod with every request

# Update app
kubectl set image deployments/example-app example-app=jocatalin/example-app:v2
kubectl rollout status deployments/example-app
kubectl rollout undo deployments/example-app
```
