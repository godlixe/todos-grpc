package postgre

import (
	"context"
	"todos-grpc/internal"
	"todos-grpc/pkg/database"
)

type DB struct {
	postgreClient database.PostgreItf
	Todo          internal.TodoDBItf
}

func New(
	ctx context.Context,
) (*DB, error) {

	postgreClient, err := database.NewPostgreClient()
	if err != nil {
		return nil, err
	}

	todoRepository := NewTodo(postgreClient)

	res := &DB{
		postgreClient: postgreClient,
		Todo:          todoRepository,
	}

	return res, nil
}
