package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a done task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Errorf("error: %v\n", err)
		}
		if err := service.Delete(id); err != nil {
			fmt.Errorf("error: %v\n", err)
		}
		fmt.Println("Task deleted 🗑️")
		return nil
	},
}
