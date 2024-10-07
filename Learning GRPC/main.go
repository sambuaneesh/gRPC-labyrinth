package main

import (
	"context"
	"grpc-intro/invoicer"
	"log"
	"net"

	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer // this is an interface provided by grpc so that if new methods are added to the server it doesnt break
}

func bussinessLogic(req *invoicer.CreateRequest) string {
	returnString := "Dihadi Details: " + req.Kisko + " ko " + req.Dihadi.Kaam + " " + req.Kisse + " karwana hai"
	return returnString
}

func (s myInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Message: bussinessLogic(req),
		Paisa:   req.Dihadi.Paisa,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &myInvoicerServer{} // pointer to the myinvoicerserver implements the invoicerserver interface

	invoicer.RegisterInvoicerServer(serverRegistrar, service)
	if err := serverRegistrar.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
