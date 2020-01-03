package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	"http-grpc-go-kit-demo/product-service/database"
	"http-grpc-go-kit-demo/product-service/pb"
	"http-grpc-go-kit-demo/product-service/product"
)

func main() {
	ctx := context.Background()
	productService := &product.ProductService{database.NewDBProductRepository()}
	errors := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errors <- fmt.Errorf("%s", <-c)
	}()

	endpoint := product.Endpoints{
		GetProductEndpoint:    product.MakeGetProductEndpoint(productService),
		GetProduct2Endpoint:   product.MakeGetProductEndpoint(productService),
		CreateProductEndpoint: product.MakeCreateProductEndpoint(productService),
	}

	//grpc
	go func() {
		listener, err := net.Listen("tcp", ":9090")
		if err != nil {
			errors <- err
			return
		}

		gRPCServer := grpc.NewServer()
		pb.RegisterProductServer(gRPCServer, product.NewGRPCServer(ctx, endpoint))

		fmt.Println("gRPC listen on 9090")
		errors <- gRPCServer.Serve(listener)
	}()

	//http
	httpAddr := fmt.Sprintf(":8011")
	r := mux.NewRouter()

	r.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"status": true})
	}).Methods("GET")

	r = product.MakeHTTPHandler(r, endpoint)

	go func() {
		fmt.Println("http listen on 8011")
		errors <- http.ListenAndServe(httpAddr, r)
	}()

	fmt.Println(<-errors)
}
