package handler

import (
	"context"
	"todos-grpc/internal"
	"todos-grpc/internal/entity"
	"todos-grpc/pkg/proto"
)

type Todo struct {
	todoService internal.TodoSvcItf
	proto.UnimplementedTodosServer
}

func NewTodo(
	todoService internal.TodoDBItf,
) *Todo {
	res := &Todo{
		todoService: todoService,
	}

	return res
}

func (t *Todo) GetTodo(ctx context.Context, req *proto.GetTodoRequest) (*proto.GetTodoResponse, error) {
	res, err := t.todoService.GetTodoByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &proto.GetTodoResponse{
		Todo: &proto.Todo{
			Id:          res.ID,
			Title:       res.Title,
			Description: res.Description,
			IsDone:      res.IsDone,
			UserId:      res.UserID,
			TimeStamp:   res.TimeStamp.String(),
		},
	}, nil
}

func (t *Todo) GetAllTodos(ctx context.Context, req *proto.GetTodosRequest) (*proto.GetTodosResponse, error) {
	res, err := t.todoService.GetAllTodos(ctx)
	if err != nil {
		return nil, err
	}

	var todosRes []*proto.Todo

	for _, data := range res {
		todosRes = append(todosRes, &proto.Todo{
			Id:          data.ID,
			Title:       data.Title,
			Description: data.Description,
			IsDone:      data.IsDone,
			UserId:      data.UserID,
			TimeStamp:   data.TimeStamp.String(),
		})
	}

	return &proto.GetTodosResponse{
		Todos: todosRes,
	}, nil
}
func (t *Todo) CreateTodo(ctx context.Context, req *proto.CreateTodoRequest) (*proto.GetTodoResponse, error) {
	var todoReq = &entity.Todo{
		Title:       req.Title,
		Description: req.Description,
		IsDone:      req.IsDone,
		UserID:      req.UserId,
	}

	res, err := t.todoService.CreateTodo(ctx, todoReq)
	if err != nil {
		return nil, err
	}

	return &proto.GetTodoResponse{
		Todo: &proto.Todo{
			Id:          res.ID,
			Title:       res.Title,
			Description: res.Description,
			IsDone:      res.IsDone,
			UserId:      res.UserID,
			TimeStamp:   res.TimeStamp.String(),
		},
	}, nil
}

func (t *Todo) UpdateTodo(ctx context.Context, req *proto.UpdateTodoResponse) (*proto.GetTodoResponse, error) {
	var todoReq = &entity.Todo{
		ID:          req.Id,
		Title:       req.Title,
		Description: req.Description,
		IsDone:      req.IsDone,
	}

	res, err := t.todoService.UpdateTodo(ctx, todoReq)
	if err != nil {
		return nil, err
	}

	return &proto.GetTodoResponse{
		Todo: &proto.Todo{
			Id:          res.ID,
			Title:       res.Title,
			Description: res.Description,
			IsDone:      res.IsDone,
			UserId:      res.UserID,
			TimeStamp:   res.TimeStamp.String(),
		},
	}, nil
}
