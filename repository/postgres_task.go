package repository

import (
	"ApiRest/model"
	"context"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
)

type postgresTaskRepository struct {
	Conn *sql.DB
}

type TaskRepository interface {
	GetTaskByID(ctx context.Context, id uint) ([]model.Task, error)
	GetAllTask(ctx context.Context) ([]model.Task, error)
	UpdateTask(ctx context.Context, id uint, p *model.Task) error
	DeleteTask(ctx context.Context, id uint) error
	AddTask(ctx context.Context, task *model.Task) error
}

func NewPostgresTaskRepository(conn *sql.DB) TaskRepository {
	return &postgresTaskRepository{conn}
}

func (p *postgresTaskRepository) GetTaskByID(ctx context.Context, id uint) (res []model.Task, err error) {
	query := `SELECT * FROM task WHERE id = $1`
	rows, err := p.Conn.QueryContext(ctx, query, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	res = make([]model.Task, 0)
	for rows.Next() {
		task := model.Task{}
		err = rows.Scan(
			&task.Id,
			&task.Title,
			&task.Description,
			&task.DueDate,
			&task.CreatedAt,
			&task.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		res = append(res, task)
	}

	return res, nil
}

func (p *postgresTaskRepository) GetAllTask(ctx context.Context) (res []model.Task, err error) {
	query := `SELECT * FROM task`
	rows, err := p.Conn.QueryContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	res = make([]model.Task, 0)
	for rows.Next() {
		task := model.Task{}
		err = rows.Scan(
			&task.Id,
			&task.Title,
			&task.Description,
			&task.DueDate,
			&task.CreatedAt,
			&task.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		res = append(res, task)
	}

	return res, nil

}

func (p *postgresTaskRepository) UpdateTask(ctx context.Context, id uint, task *model.Task) (err error) {
	query := `UPDATE task set title=$1, description=$2, due_date=$3 WHERE id = $4`

	_, err = p.Conn.QueryContext(ctx, query, task.Title, task.Description, task.DueDate, id)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return
}

func (p *postgresTaskRepository) DeleteTask(ctx context.Context, id uint) (err error) {
	query := `DELETE FROM task WHERE id = $1`

	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		logrus.Error(err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		logrus.Error(err)
		return err
	}

	if rowsAffected == 0 {
		err = fmt.Errorf("task with ID %d not found", id)
		return err
	} else if rowsAffected != 1 {
		err = fmt.Errorf("weird behavior. Total affected: %d", rowsAffected)
		return err
	}

	return nil
}

func (p *postgresTaskRepository) AddTask(ctx context.Context, task *model.Task) (err error) {
	query := `INSERT INTO task (title, description, due_date) VALUES ($1, $2, $3) RETURNING id`

	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	err = stmt.QueryRowContext(ctx, task.Title, task.Description, task.DueDate).Scan(&task.Id)
	if err != nil {
		return err
	}

	return
}
