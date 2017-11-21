package main

import (
	"fmt"
	"golang.org/x/net/context"
	"github.com/scottrangerio/grpc/messages"
	"net"
	"google.golang.org/grpc"
)
const port = ":6789"

type Service struct {}

func (service Service) Add(ctx context.Context, req *messages.AddRequest) (*messages.AddResponse, error) {
    result := &messages.AddResponse{}
    result.Result = req.X + req.Y
    return result,nil
}

func main() {
	lis,_ := net.Listen("tcp", port)
	s := grpc.NewServer()
        messages.RegisterAddServiceServer(s, new(Service))
        fmt.Println("Listening")
	s.Serve(lis)
}
