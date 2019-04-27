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

## Dependency Management
```
https://github.com/golang/dep
```

## API call on terminal
```
$ curl -X GET "http://localhost:8002/basket/list"

$ curl  -X POST http://localhost:8002/basket/save -d '{"name":"test1", "surname":"test2"}' -H "Content-Type: application/json"
```
