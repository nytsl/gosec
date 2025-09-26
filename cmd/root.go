package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "awesomeProject",
	Short: "一个功能强大的域名信息查询工具",
	Long: `awesomeProject 是一个用于查询域名相关信息的命令行工具。
支持查询ICP备案信息、IP反查域名、Whois信息等功能。`,
}

// Execute 执行根命令
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// 禁用默认自动补全命令（暂时不需要）
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
