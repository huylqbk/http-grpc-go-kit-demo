# go-kit Rest API (http) and gRPC 

## Run
go run product-service/main.go

contributor: huylqbk

## Test GRPC:
- install BloomRPC
## Test Http Server 
```
curl -X POST \
  http://localhost:8011/api/v1/product \
  -d '{
	"name":"hello"
}'
```