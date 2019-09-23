# basket app

 https://microservice-base.github.io

Ana projenin sepet modulü projesidir.

## Contributing

[Contributing guide](CONTRIBUTING.md).

```
$ git clone https://github.com/microservice-base/basket.git

$ cd basket/basket

$ dep ensure

$ go test
$ go run basketApplication.go

or

$ go build basketApplication.go 

$ ./basketApplication
```

## docker
```
$  docker pull keramiozsoy/image-basket
$  docker run -it --rm --name app-basket -p 8002:8002 image-basket // just run

or

$  git clone https://github.com/microservice-base/basket.git
$  cd basket
$  docker build -t image-basket  -f container/docker/Dockerfile . 
$  docker run -d --name app-basket -p 8002:8002 image-basket

$  docker rmi $(docker images | grep "<none>" | awk '{print $3}')
$  docker inspect -f '{{.Name}} - {{.NetworkSettings.IPAddress }}' $(docker ps -aq)
```

## kubernetes ( on minikube )
```
$  kubectl create deployment image-basket-deployment --image=keramiozsoy/image-basket
$  kubectl expose deployment image-basket-deployment --type=LoadBalancer --port=8002

$  minikube service image-basket-deployment

$  kubectl get deployment,rs,po  -l app=image-basket-deployment
$  kubectl scale deployment image-basket-deployment --replicas=3

$  kubectl get deployment,rs,po  -l app=image-basket-deployment
$  kubectl describe deployment image-basket-deployment

$  kubectl get deployment,rs,po  -l app=image-basket-deployment
$  kubectl rollout history deployment image-basket-deployment
$  kubectl rollout history deployment image-basket-deployment --revision=1
$  kubectl rollout history deployment image-basket-deployment --revision=2 ( Doesnt have different replicaset , different image version)

$  kubectl set image deployment image-basket-deployment --image=keramiozsoy/image-basket:v1
$  kubectl rollout history deployment image-basket-deployment --revision=1 
$  kubectl rollout history deployment image-basket-deployment --revision=2 (revisions have differend dockerized version)
$  kubectl get deployment,rs,po  -l app=image-basket-deployment (old revision replicaset now scaled down to zero :) )

$  kubectl rollout undo deploymnet image-basket-deployment --to-revision=1 ( undo deployment)
$  kubectl get deployment,rs,po  -l app=image-basket-deployment
$  kubectl rollout history deployment image-basket-deployment --revision=2
$  kubectl rollout history deployment image-basket-deployment --revision=3 (current state but check same replicaset name  same revision 1 )


$  kubectl delete service image-basket-deployment
$  kubectl delete deployment image-basket-deployment

```

## api call on terminal
```
$ curl -X GET "http://localhost:8002/basket/list"

$ curl -X POST http://localhost:8002/basket/save -d '{"name":"test1", "surname":"test2"}' -H "Content-Type: application/json"
```
