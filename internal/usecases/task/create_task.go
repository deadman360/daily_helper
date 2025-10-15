package usecases

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (task *task) CreateTask(cmd *cobra.Command, args []string) {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		fmt.Println("Error reading flag:", err)
		return
	}
	task.repository.Create(name)
}
