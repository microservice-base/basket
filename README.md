# basket

https://microservice-base.github.io/

cd basket

go test


go build BasketApi.go 

./BasketApi

## dependecny management like gradle maven 
```bash
brew install dep
```
projemizi dep ile yönetmeye nasıl başlarız

projede hangi klasörde init komutu verirsen orada oluşuyor.
```go
dep init
```
toml ve lock dosyaları oluşacak

bir yeni dependency eklemek

https://github.com/naoina/go-stringutil

http ve https olmadan yazmalısın
```go
dep ensure -add github.com/naoina/go-stringutil
```

veya https://glide.sh/ kullan

## API
```
curl -X GET "http://localhost:8002/basket/list"

curl  -X POST http://localhost:8002/basket/save -d '{"name":"test1", "surname":"test2"}' -H "Content-Type: application/json"
```
