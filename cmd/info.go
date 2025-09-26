package cmd

import (
	"awesomeProject/internal/usecase"
	"awesomeProject/pkg/logger"

	"github.com/spf13/cobra"
)

var (
	icpDomains       []string
	ip138Addresses   []string
	whoisDomains     []string
	subfinderDomains []string
	proxy            string
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "查询域名相关信息",
	Long:  `查询域名的ICP备案信息、IP反查域名信息或Whois信息`,
	Run: func(cmd *cobra.Command, args []string) {
		infoService := usecase.NewQueryManager()
		if cmd.Flags().NFlag() == 0 {
			cmd.Help()
			return
		}
		// 设置代理
		if proxy != "" {
			infoService.SetProxy(proxy)
		}

		// 处理ICP查询
		if len(icpDomains) > 0 {
			logger.Log.InfoMsgf("开始处理 ICP 信息查询")
			if err := infoService.QueryICP(icpDomains); err != nil {
				logger.Log.ErrorMsgf("ICP查询错误: %s", err)
			}
		}

		// 处理IP反查
		if len(ip138Addresses) > 0 {
			logger.Log.InfoMsgf("开始处理 IP 反查")
			if err := infoService.QueryIP(ip138Addresses); err != nil {
				logger.Log.ErrorMsgf("IP查询错误: %s", err)
			}
		}

		// 处理Whois查询
		if len(whoisDomains) > 0 {
			logger.Log.InfoMsgf("开始处理 Whois 查询")
			// Whois 不支持代理，但为了接口一致性，也可以传递代理参数（会被忽略）
			if err := infoService.QueryWhois(whoisDomains); err != nil {
				logger.Log.ErrorMsgf("Whois查询错误: %s", err)
			}
		}

		if len(subfinderDomains) > 0 {
			logger.Log.InfoMsgf("开始处理 Subfinder 子域名查询")
			if err := infoService.QuerySubdomains(subfinderDomains); err != nil {
				logger.Log.ErrorMsgf("Subfinder查询错误: %s", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	// 添加命令行标志
	infoCmd.Flags().StringSliceVarP(&icpDomains, "icp", "i", []string{}, "要查询ICP备案的域名列表，多个用逗号分隔")
	infoCmd.Flags().StringSliceVarP(&ip138Addresses, "ip", "n", []string{}, "要进行反查的IP地址列表，多个用逗号分隔")
	infoCmd.Flags().StringSliceVarP(&whoisDomains, "whois", "w", []string{}, "要查询Whois的域名列表，多个用逗号分隔")
	infoCmd.Flags().StringVarP(&proxy, "proxy", "p", "", "设置 HTTP 代理，例如：http://127.0.0.1:8080")
}
