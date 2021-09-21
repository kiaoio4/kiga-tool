package cmd

import (
	"fmt"
	"kiga-tool/pkg/mysql"
	"os"
	"runtime/debug"

	"github.com/spf13/cobra"
)

// SetupCmd setup db for task, and could called from platform-cli
var SetupCmd = &cobra.Command{
	Use:   "run",
	Short: "setup db for task in the system",
	Run: func(cmd *cobra.Command, args []string) {
		// recover
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered", r)
				debug.PrintStack()
			}
		}()

		mysqlConf := mysql.GetMysqlConfig()
		err := mysql.CreateDatabase(*mysqlConf)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
}
