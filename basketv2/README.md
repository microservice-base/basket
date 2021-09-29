go mod init basketv2

go mod tidy
//go mod vendor

go build .

./basketv2

curl -X GET http://localhost:3000
