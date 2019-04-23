# basket

https://microservice-base.github.io/

cd basket

go test

go build BasketApi.go 

./BasketApi


## API
```
curl -X GET "http://localhost:8002/basket/list"

curl  -X POST http://localhost:8002/basket/save -d '{"name":"test1", "surname":"test2"}' -H "Content-Type: application/json"
```
