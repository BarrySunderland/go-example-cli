package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var days int

func init() {
	rootCmd.AddCommand(addTask)
	addTask.Flags().IntVarP(&days, "days", "d", 1, "Number of days to complete the task e.g. 1 for a task that's due tomorrow")
}

var addTask = &cobra.Command{
	Use:   "add",
	Short: "add a task to the list",
	Long: `add a task to the list
accepts a single string as an argument. example:
task add "eat cheese"`,
	Run: func(cmd *cobra.Command, args []string) {

		task := args[0]
		days, _ := cmd.Flags().GetInt("days")
		fmt.Println("days:", days)

		t := time.Now().Format("2006-Jan-01 15:04")
		msg := fmt.Sprintf("%s    added task: %s", t, task)

		fmt.Sprintf("due in: %v days", days)
		fmt.Println(msg)
	},
}
