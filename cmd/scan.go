package cmd

import (
	"awesomeProject/internal/services"
	"awesomeProject/pkg/logger"
	"time"

	"github.com/spf13/cobra"
)

var (
	scanHost      string
	scanPort      int
	scanPortRange string
	scanTimeout   int
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "网络扫描功能",
	Long:  `执行网络扫描，包括端口扫描、服务识别等`,
	Run: func(cmd *cobra.Command, args []string) {
		scanService := services.NewScanService()

		// 设置超时时间
		if scanTimeout > 0 {
			scanService.SetTimeout(time.Duration(scanTimeout) * time.Second)
		}

		// 单端口扫描
		if scanHost != "" && scanPort > 0 {
			if err := scanService.ScanPort(scanHost, scanPort); err != nil {
				logger.Log.ErrorMsgf("端口扫描失败: %s", err)
			}
			return
		}

		// 端口范围扫描
		if scanHost != "" && scanPortRange != "" {
			// 解析端口范围 (例如: "80-443")
			// TODO: 实现端口范围解析逻辑
			logger.Log.InfoMsgf("端口范围扫描功能待实现: %s", scanPortRange)
			return
		}

		// 如果没有任何参数，显示帮助
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	// 添加命令行标志
	scanCmd.Flags().StringVar(&scanHost, "host", "", "目标主机地址")
	scanCmd.Flags().IntVar(&scanPort, "port", 0, "目标端口")
	scanCmd.Flags().StringVar(&scanPortRange, "port-range", "", "端口范围 (例如: 80-443)")
	scanCmd.Flags().IntVar(&scanTimeout, "timeout", 3, "扫描超时时间（秒）")
}
