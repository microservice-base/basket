# basket

https://microservice-base.github.io



```
$ git clone https://github.com/microservice-base/basket.git

$ dep ensure

$ cd basket/basket

$ go test

$ go run basketApplication.go

or

$ go build basketApplication.go 

$ ./basketApplication
```
## dependency management
```
https://github.com/golang/dep
```

## docker
```
$  docker pull keramiozsoy/image-basket
$  docker run -it --rm --name basket-app -p 8002:8002 image-basket // just run

or

$  git clone https://github.com/microservice-base/basket.git
$  cd basket
$  docker build -t image-basket  -f container/docker/Dockerfile . 
$  docker run -d --name basket-app -p 8002:8002 image-basket

$  docker rmi $(docker images | grep "<none>" | awk '{print $3}')
```

## kubernetes ( minikube )
```
$  kubectl create deployment image-basket-deployment --image=keramiozsoy/image-basket
$  kubectl expose deployment image-basket-deployment --type=LoadBalancer --port=8002

$  kubectl delete service image-basket-deployment
$  kubectl delete deployment image-basket-deployment

```

## api call on terminal
```
$ curl -X GET "http://localhost:8002/basket/list"

$ curl  -X POST http://localhost:8002/basket/save -d '{"name":"test1", "surname":"test2"}' -H "Content-Type: application/json"
```
