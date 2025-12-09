package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "list", Description: "all of the tasks list"},
		{Text: "add", Description: "add a new task"},
		{Text: "done", Description: "mark as done task"},
		{Text: "delete", Description: "delete a task"},
		{Text: "clear", Description: "clear a done task"},
		{Text: "help", Description: "show this help"},
		{Text: "exit", Description: "exit the program"},
		{Text: "quit", Description: "quit the program"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)

}
func executor(in string) {
	in = strings.TrimSpace(in)
	if in == "" {
		return
	}
	args := strings.Fields(in)
	command := args[0]
	switch command {
	case "list":
		listTasksInteractive()
	case "add":
		if len(args) < 2 {
			color.Red("you must specify a task name")
			return
		}
		title := strings.Join(args[1:], " ")
		if err := service.Add(title); err != nil {
			color.Red(err.Error())
			return
		} else {
			color.Green("task %s added", title)
		}
	case "done":
		if len(args) != 2 {
			color.Red("⚠️  استفاده: done <id>")
			return
		}
		id, _ := strconv.Atoi(args[1])
		if err := service.Done(id); err != nil {
			color.Red("خطا: %v", err)
		} else {
			color.Green("task %s done", args[1])
		}
	case "delete":
		if len(args) < 2 {
			color.Red("you must write id of a task")
			return
		}
		id, _ := strconv.Atoi(args[0])
		if err := service.Delete(id); err != nil {
			color.Red(err.Error())
			return
		}
		color.Green("task %s deleted", args[0])

	case "clear":
		if err := service.ClearDone(); err != nil {
			color.Red(err.Error())
		}
		color.Green("task %s cleared", args[0])
	case "help":
		fmt.Println(color.CyanString("avaliable tasks:\n"))
		fmt.Println("list ------------------ list of tasks")
		fmt.Println("add -------------- add a new task")
		fmt.Println("done -------------- add a task done")
		fmt.Println("delete -------------- delete a task done")
		fmt.Println("clear -------------- clear a task done")
	case "exit", "quit":
		color.Yellow("task-manager exiting...")
		os.Exit(0)
	default:
		fmt.Println(color.CyanString("invalid command"))
	}
}
func listTasksInteractive() {
	tasks, err := service.List()
	if err != nil {
		color.Red(err.Error())
		return
	}
	if len(tasks) == 0 {
		color.Red("no tasks found")
		return
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Created At"})

	for _, task := range tasks {
		status := "⬜"
		statusColor := color.New(color.FgWhite).SprintFunc()
		title := task.Title

		if task.Done {
			status = "✅"
			statusColor = color.New(color.FgGreen).SprintFunc()
			title = fmt.Sprintf("~~%s~~", title)
		}

		t.AppendRow(table.Row{
			color.CyanString("%d", task.ID),
			statusColor(status),
			title,
		})
	}

	t.SetStyle(table.StyleLight)
	t.Render()
}
