package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "mark as done",
	Long:  `Mark the task with the given ID as done. Example: task-manager done 3`,
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Errorf("error: %v\n", err)
		}
		return service.Done(id)
	},
}
