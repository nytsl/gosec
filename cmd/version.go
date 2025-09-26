package cmd

import (
	"awesomeProject/pkg/config"
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "显示版本信息",
	Long:  `显示 awesomeProject 的版本信息`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("awesomeProject 版本: %s\n", config.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
