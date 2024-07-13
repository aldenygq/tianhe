package cmd

import (
	"errors"
	"fmt"
	"os"
	
	"tianhe/cmd/api"
	
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:               "./tianhe server -c ./conf/conf.yaml",
	Short:             "-v",
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	Long:              `Tianhe Operation and Maintenance Management Platform`,
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
		usageStr := `欢迎使用天和运维管理平台，可以使用 -h 查看命令`
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
