package cmd

import (
	"github.com/labadiejoaco/todo-cli/todo"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to the todo list",
	Long:  "Add a new task to the todo list.",
	Run: func(cmd *cobra.Command, args []string) {
		todos := todo.GetTodos()

		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")

		todo.AddTodo(todos, name, description)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().String("name", "", "Todo name")
	addCmd.Flags().String("description", "", "Todo description")

	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("description")
}
