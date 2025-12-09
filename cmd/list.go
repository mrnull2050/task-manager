package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "list of task",
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := service.List()
		if err != nil {
			return err
		}
		for _, task := range tasks {
			status := " "
			if task.Done {
				status = "*"
			}

			fmt.Printf("%d , [%s] , %s , \n ", task.ID, status, task.Title)
		}
		return nil

	},
}
