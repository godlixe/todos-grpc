package main

import (
	"fmt"
	"todos-grpc/internal/app/grpc"
)

func main() {
	err := grpc.Start()
	if err != nil {
		fmt.Println(err)
	}
}
