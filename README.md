# basket

https://microservice-base.github.io/


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
## Docker
```
$  git clone https://github.com/microservice-base/basket.git
$  cd basket
$  docker build -t image-basket  -f container/docker/Dockerfile . 
$  docker run -d --name basket-app -p 8002:8002 image-basket

$  docker rmi $(docker images | grep "<none>" | awk '{print $3}')
```

## Dependency Management
```
https://github.com/golang/dep
```

## API call on terminal
```
$ curl -X GET "http://localhost:8002/basket/list"

$ curl  -X POST http://localhost:8002/basket/save -d '{"name":"test1", "surname":"test2"}' -H "Content-Type: application/json"
```
