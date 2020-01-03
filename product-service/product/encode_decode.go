package product

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"http-grpc-go-kit-demo/product-service/pb"
)

func EncodeProductResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(ProductResponse)
	return &pb.ProductResponse{
		Id:   res.ID,
		Name: res.Name,
	}, nil
}

func EncodeHTTPResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func DecodeGetProductRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetProductRequest)
	return GetProductRequest{req.Id}, nil
}

func DecodeCreateProductRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CreateProductRequest)
	return CreateProductRequest{req.Name}, nil
}

func DecodeHttpGetProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetProductRequest
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	request.ID = int32(id)
	return request, nil
}

func DecodeHttpCreateProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
