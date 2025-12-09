package cmd

import (
	"fmt"
	"os"
	"task-manager/intenal/core"
	"task-manager/intenal/storage"

	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var service *core.TaskService
var rootCmd = &cobra.Command{
	Use:           "task-manager",
	Short:         "task-manager cli application",
	SilenceUsage:  true,
	SilenceErrors: true,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			cmd.Help()
			return
		}
		runInteractive()
	},
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
	rootCmd.AddCommand(addCmd, ListCmd, doneCmd, deleteCmd, clearCmd)
}
func runInteractive() {
	fmt.Println(color.CyanString("welcome to task-manager"))
	fmt.Println(color.YellowString("write exit or quit for exiting the application"))
	fmt.Println(color.MagentaString("help for see all the options"))
	fmt.Println()
	p := prompt.New(
		executor,
		completer,
		prompt.OptionTitle("task-manager"),
		prompt.OptionPrefix("tasks >>> "),
		prompt.OptionPrefixTextColor(prompt.Cyan),
		prompt.OptionPreviewSuggestionTextColor(prompt.Yellow),
		prompt.OptionPreviewSuggestionBGColor(prompt.DarkBlue),
		prompt.OptionDescriptionTextColor(prompt.Black),
	)
	p.Run()
}
func Execute() error {
	return rootCmd.Execute()
}
