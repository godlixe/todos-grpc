package service

import (
	"todos-grpc/internal"
	"todos-grpc/internal/repository/db/postgre"
)

type Domain struct {
	Todo internal.TodoSvcItf
}

func New(
	postgreDB *postgre.DB,
) *Domain {
	todoSvc := NewTodo(
		postgreDB.Todo,
	)

	res := &Domain{
		Todo: todoSvc,
	}

	return res
}
