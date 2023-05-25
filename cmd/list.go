package cmd

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/labadiejoaco/todo-cli/todo"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks in the todo list",
	Long:  "List all tasks in the todo list.",
	Run: func(cmd *cobra.Command, args []string) {
		todos := todo.GetTodos()

		if len(todos) <= 0 {
			fmt.Println("Your todo list is currently empty.")
		} else {
			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"Id", "Name", "Description"})

			for i, todo := range todos {
				t.AppendRow([]interface{}{i + 1, todo.Name, todo.Description})
				t.AppendSeparator()
			}

			t.Render()
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
