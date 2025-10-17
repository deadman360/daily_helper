package task_repository

import (
	"fmt"

	"github.com/deadman360/daily_helper/internal/domain"
)

func (r *repository) CreateTask(task *domain.Task) {
	fmt.Println(task.Name)
}
