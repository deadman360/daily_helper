package task_repository

import (
	"database/sql"

	"github.com/deadman360/daily_helper/internal/domain"
)

type repository struct {
	db *sql.DB
}

type ITaskRepository interface {
	CreateTask(*domain.Task)
}

func NewTaskRepository(db *sql.DB) ITaskRepository {
	return &repository{
		db: db,
	}
}
