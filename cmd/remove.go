package cmd

import (
	"github.com/labadiejoaco/todo-cli/todo"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a task from the todo list",
	Long:  "Remove a task from the todo list.",
	Run: func(cmd *cobra.Command, args []string) {
		todos := todo.GetTodos()

		name, _ := cmd.Flags().GetString("name")

		todo.RemoveTodo(todos, name)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().String("name", "", "Todo name")

	removeCmd.MarkFlagRequired("name")
}
