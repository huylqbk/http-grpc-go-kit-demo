package product

import (
	"http-grpc-go-kit-demo/product-service/httpjson"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHTTPHandler(r *mux.Router, endpoint Endpoints) *mux.Router {
	logger := log.NewLogfmtLogger(os.Stdout)
	options := httpjson.DefaultServerOptions(logger)
	v1 := r.PathPrefix("/api/v1").Subrouter()

	//public v1 API
	v1.Methods("GET").Path("/product").
		Handler(httptransport.NewServer(
			endpoint.GetProductEndpoint,
			DecodeHttpGetProductRequest,
			EncodeHTTPResponse,
			options...,
		))

	v1.Methods("POST").Path("/product").
		Handler(httptransport.NewServer(
			endpoint.CreateProductEndpoint,
			DecodeHttpCreateProductRequest,
			EncodeHTTPResponse,
			options...,
		))

	return r
}
