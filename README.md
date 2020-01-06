# go-kit Rest API (http) and gRPC 

## Run
- Proto3 install
- protoc product.proto --go_out=plugins=grpc:product-service/pb
- go run product-service/main.go

contributor: huylqbk

## Test GRPC:
- install BloomRPC](https://github.com/uw-labs/bloomrpc)
## Test Http Server 
```
curl -X POST \
  http://localhost:8011/api/v1/product \
  -d '{
	"name":"hello"
}'
```
