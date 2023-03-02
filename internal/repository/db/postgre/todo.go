package postgre

import (
	"context"
	"fmt"
	"todos-grpc/internal/entity"
	"todos-grpc/pkg/database"

	"github.com/jackc/pgtype"
)

// Todo repository
type Todo struct {
	postgreClient database.PostgreItf
}

func NewTodo(
	postgreClient database.PostgreItf,
) *Todo {
	todoRepository := &Todo{
		postgreClient: postgreClient,
	}

	return todoRepository
}

func (t *Todo) GetTodoByID(ctx context.Context, id int64) (*entity.Todo, error) {
	db, err := t.postgreClient.GetConn(ctx)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprint(`
		SELECT
			*
		FROM
			todos
		WHERE
			id = ($1)
	`)

	var todo entity.Todo
	var todoTime pgtype.Timestamptz

	err = db.QueryRow(ctx, query, id).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Description,
		&todo.IsDone,
		&todo.UserID,
		&todoTime,
	)
	if err != nil {
		return nil, err
	}

	// set the time
	todo.TimeStamp = todoTime.Time

	return &todo, err
}

func (t *Todo) GetAllTodos(ctx context.Context) ([]entity.Todo, error) {
	return nil, nil
}

// 	db, err := t.postgreClient.GetConn(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	query := fmt.Sprint(`
// 		SELECT
// 			*
// 		FROM
// 			todos
// 		WHERE
// 			id = ($1)
// 	`)

// 	var todo entity.Todo
// 	err = db.QueryRow(ctx, query, id).Scan(
// 		&todo.ID,
// 		&todo.Title,
// 		&todo.Description,
// 		&todo.IsDone,
// 		&todo.UserID,
// 		&todo.TimeStamp,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &todo, err
// }

func (t *Todo) CreateTodo(ctx context.Context, todo *entity.Todo) error {
	return nil
}

// 	db, err := t.postgreClient.GetConn(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	query := fmt.Sprint(`
// 		SELECT
// 			*
// 		FROM
// 			todos
// 		WHERE
// 			id = ($1)
// 	`)

// 	var todo entity.Todo
// 	err = db.QueryRow(ctx, query, id).Scan(
// 		&todo.ID,
// 		&todo.Title,
// 		&todo.Description,
// 		&todo.IsDone,
// 		&todo.UserID,
// 		&todo.TimeStamp,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &todo, err
// }

func (t *Todo) UpdateTodo(ctx context.Context, todo *entity.Todo) error {
	return nil
}

// 	db, err := t.postgreClient.GetConn(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	query := fmt.Sprint(`
// 		SELECT
// 			*
// 		FROM
// 			todos
// 		WHERE
// 			id = ($1)
// 	`)

// 	var todo entity.Todo
// 	err = db.QueryRow(ctx, query, id).Scan(
// 		&todo.ID,
// 		&todo.Title,
// 		&todo.Description,
// 		&todo.IsDone,
// 		&todo.UserID,
// 		&todo.TimeStamp,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &todo, err
// }
