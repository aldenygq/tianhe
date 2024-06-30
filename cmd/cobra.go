package cmd

import (
	"errors"
	"fmt"
	"os"
	
	"oncall/cmd/api"
	
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:               "oncall schedule tool",
	Short:             "-v",
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	Long:              `oncall`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("require param need least one arg")
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		usageStr := `欢迎使用oncall，可以使用 -h 查看命令`
		fmt.Printf("%s\n", usageStr)
	},
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
}

//Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
