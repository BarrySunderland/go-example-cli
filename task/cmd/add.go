package cmd

import (
	"fmt"
	"os"
	"strings"
	"task/db"
	"time"

	"github.com/spf13/cobra"
)

var days int

func init() {
	RootCmd.AddCommand(addTask)
	addTask.Flags().IntVarP(&days, "days", "d", 1, "Number of days to complete the task e.g. 1 for a task that's due tomorrow")
}

var addTask = &cobra.Command{
	Use:   "add",
	Short: "add a task to the list",
	Long: `add a task to the list
accepts a single string as an argument. example:
task add "eat cheese"`,
	Run: func(cmd *cobra.Command, args []string) {

		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("failed to create task: ", err.Error())
			os.Exit(1)
		}

		days, _ := cmd.Flags().GetInt("days")
		fmt.Println("days:", days)

		t := time.Now().Format("2006-Jan-01 15:04")
		msg := fmt.Sprintf("%s    added task: %s", t, task)

		msg += fmt.Sprintf(" due in: %v days", days)
		fmt.Println(msg)
	},
}
