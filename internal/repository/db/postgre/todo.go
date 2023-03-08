package postgre

import (
	"context"
	"todos-grpc/internal/entity"
	"todos-grpc/pkg/database"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
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

	query := `
		SELECT
			*
		FROM
			todos
		WHERE
			id = ($1)
	`

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

func (t *Todo) GetAllTodos(ctx context.Context) ([]*entity.Todo, error) {
	db, err := t.postgreClient.GetConn(ctx)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT
			*
		FROM
			todos
	`

	var todos []*entity.Todo

	var rows pgx.Rows

	rows, err = db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo entity.Todo
		var todoTime pgtype.Timestamptz
		if err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Description,
			&todo.IsDone,
			&todo.UserID,
			&todoTime,
		); err != nil {
			return nil, err
		}

		todo.TimeStamp = todoTime.Time

		todos = append(todos, &todo)
	}

	return todos, err
}

func (t *Todo) CreateTodo(ctx context.Context, todo *entity.Todo) (*entity.Todo, error) {
	db, err := t.postgreClient.GetConn(ctx)
	if err != nil {
		return nil, err
	}

	query := `
		INSERT INTO todos (
			title,
			description,
			is_done,
			user_id,
			time_stamp
		)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			CURRENT_TIMESTAMP
		)
	`

	_, err = db.Exec(
		context.Background(),
		query,
		todo.Title,
		todo.Description,
		todo.IsDone,
		todo.UserID,
	)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (t *Todo) UpdateTodo(ctx context.Context, todo *entity.Todo) (*entity.Todo, error) {
	db, err := t.postgreClient.GetConn(ctx)
	if err != nil {
		return nil, err
	}

	query := `
		UPDATE todos 
		SET
			title = $2,
			description = $3,
			is_done = $4,
			user_id = $5,
			time_stamp = CURRENT_TIMESTAMP
		WHERE
			id = $1
	`

	_, err = db.Exec(
		context.Background(),
		query,
		todo.ID,
		todo.Title,
		todo.Description,
		todo.IsDone,
		todo.UserID,
	)
	if err != nil {
		return nil, err
	}

	return todo, nil
}
