package grpc

import (
	"fmt"
	"net"
	"todos-grpc/internal/app"
	"todos-grpc/internal/app/grpc/handler"
	"todos-grpc/pkg/proto"

	"google.golang.org/grpc"
)

func Start() error {

	system, err := app.New()
	if err != nil {
		return err
	}

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return err
	}

	server := grpc.NewServer()

	todosHandler := handler.NewTodo(
		system.Domain.Todo,
	)

	proto.RegisterTodosServer(server, todosHandler)

	if err = server.Serve(listener); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
