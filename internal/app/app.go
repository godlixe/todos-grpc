package app

import (
	"context"
	"todos-grpc/internal/repository/db/postgre"
	"todos-grpc/internal/service"
	"todos-grpc/pkg/proto"
)

type TodosHandler struct {
	proto.UnimplementedTodosServer
}

type System struct {
	Domain *service.Domain
	DB     *postgre.DB
}

func New() (*System, error) {

	postgreDB, err := postgre.New(context.Background())
	if err != nil {
		return nil, err
	}

	domain := service.New(
		postgreDB,
	)

	res := &System{
		Domain: domain,
	}

	return res, nil
}
