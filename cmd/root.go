package cmd

import (
	"os"
	"task-manager/intenal/core"
	"task-manager/intenal/storage"

	"github.com/spf13/cobra"
)

var service *core.TaskService
var rootCmd = &cobra.Command{
	Use:           "task-manager",
	Short:         "task-manager cli application",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	home, _ := os.UserHomeDir()
	st := storage.NewJsonStorage(home + "/task-manager.json")
	service = core.NewTaskService(st)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(ListCmd)
	rootCmd.AddCommand(clearCmd)
	rootCmd.AddCommand(doneCmd)
	rootCmd.AddCommand(deleteCmd)
}
func Execute() error {
	return rootCmd.Execute()
}
