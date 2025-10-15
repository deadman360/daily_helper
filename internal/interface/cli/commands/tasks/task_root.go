package interfaces

import (
	"github.com/deadman360/daily_helper/internal/infrastructure"
	usecases "github.com/deadman360/daily_helper/internal/usecases/task"
	"github.com/spf13/cobra"
)

func TaskCommand(repository infrastructure.IRepository) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "task",
		Short: "Manage tasks",
	}

	rootCmd.Flags().String("name", "", "Task name (required)")
	rootCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(createTaskCommand(usecases.NewTaskInterface(repository)))
	return rootCmd
}
