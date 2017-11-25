package main

import (
    "github.com/scottrangerio/grpc/messages"
    "fmt"
    "log"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
)

var addr = "localhost:6789"

func main() {
    req := &messages.AddRequest{
        X: 1,
        Y: 2,
    }

    conn, err := grpc.Dial(addr, grpc.WithInsecure())

    if err != nil {
      log.Fatal(err)
    }

    defer conn.Close()

    client := messages.NewAddServiceClient(conn)

    res, err := client.Add(context.Background(), req)

    if err != nil {
      log.Fatal(err)
    }

    fmt.Printf("%v\n", res)
}
