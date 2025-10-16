package task_repository

import "database/sql"

type repository struct {
	db *sql.DB
}

type ITaskRepository interface {
	CreateTask(string)
}

func NewTaskRepository(db *sql.DB) ITaskRepository {
	return &repository{
		db: db,
	}
}
