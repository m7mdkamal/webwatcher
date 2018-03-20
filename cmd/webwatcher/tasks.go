package cmd

import (
	"fmt"

	"github.com/m7mdkamal/webwatcher/model"

	"github.com/spf13/cobra"
)

func init() {
	addTaskCmd.Flags().StringP("name", "n", "", "Name help you identify the task")
	addTaskCmd.Flags().StringP("filter", "f", "", "Filter to get the results you want")
	addTaskCmd.Flags().IntP("interval", "i", 60*24, "Run every X mins")

	tasksCmd.AddCommand(listTasksCmd)
	tasksCmd.AddCommand(clearTasksCmd)
	tasksCmd.AddCommand(addTaskCmd)
	rootCmd.AddCommand(tasksCmd)
}

var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var listTasksCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		listTasks()
	},
}

var addTaskCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new tasks",
	Long:  "Add new tasks",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		filter, _ := cmd.Flags().GetString("filter")
		interval, _ := cmd.Flags().GetInt("interval")
		task := model.Task{
			Name:     name,
			Filter:   filter,
			Interval: interval,
		}
		DB.CreateTask(&task)
	},
}

var clearTasksCmd = &cobra.Command{
	Use:   "clear",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CLEAR tasks HERE")
	},
}

func listTasks() {
	tasks, err := DB.GetTasks()
	if err != nil {
		panic(err)
	}
	if len(tasks) == 0 {
		fmt.Println("No Tasks Found")
	} else {
		for _, task := range tasks {
			fmt.Println(task.Name, task.Filter, task.Interval, task.CreatedAt)
		}
	}
}
