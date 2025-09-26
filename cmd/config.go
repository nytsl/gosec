package cmd

import (
	"awesomeProject/pkg/config"
	"awesomeProject/pkg/logger"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	setBeiAnCookie string
	showConfig     bool
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "配置管理",
	Long:  `管理应用程序的配置信息，如Cookie等`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			logger.Log.ErrorMsgf("加载配置失败: %v", err)
			return
		}

		// 设置备案Cookie
		if setBeiAnCookie != "" {
			cfg.Cookies.BeiAn = setBeiAnCookie
			if err := config.SaveConfig(cfg); err != nil {
				logger.Log.ErrorMsgf("保存配置失败: %v", err)
				return
			}
			logger.Log.InfoMsgf("备案Cookie已更新")
		}

		// 显示当前配置
		if showConfig {
			fmt.Printf("当前配置:\n")
			fmt.Printf("  备案Cookie: %s\n", cfg.Cookies.BeiAn)
		}

		// 如果没有任何参数，显示帮助
		if setBeiAnCookie == "" && !showConfig {
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// 添加命令行标志
	configCmd.Flags().StringVar(&setBeiAnCookie, "set-beian-cookie", "", "设置ICP备案查询所需 Cookie（machine_str）")
	configCmd.Flags().BoolVar(&showConfig, "show", false, "显示当前已保存的配置")
}
