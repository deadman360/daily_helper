package infrastructure

import (
	"database/sql"

	"github.com/deadman360/daily_helper/internal/infrastructure/task_repository"
)

type repository struct {
	task_repository.ITaskRepository
}

type IRepository interface {
	task_repository.ITaskRepository
}

func NewRepository(db *sql.DB) IRepository {
	task_repo := task_repository.NewTaskRepository(db)

	return &repository{
		ITaskRepository: task_repo,
	}
}
