package services

import (
	"awesomeProject/pkg/config"
	"awesomeProject/pkg/logger"
	"awesomeProject/pkg/table"
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/imroc/req/v3"
	prettytable "github.com/jedib0t/go-pretty/v6/table"
)

// IPService IP反查服务
type IPService struct{}

// NewIPService 创建IP查询服务
func NewIPService() *IPService {
	return &IPService{}
}

// Query 查询IP反查信息
func (s *IPService) Query(ip string) error {
	logger.Log.InfoMsgf("正在进行IP %s 的反查", ip)

	client := req.C()
	// 使用全局代理配置
	if proxy := config.GetGlobalProxy(); proxy != "" {
		client.SetProxyURL(proxy)
	}

	request := client.R()
	resp, err := request.Get("https://ip138.com/" + ip)
	if err != nil {
		return logger.Log.ErrorMsgf("网络请求失败:%s", err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return logger.Log.ErrorMsgf("HTML解析失败：%s", err)
	}

	address := doc.Find("h3").First().Text()
	noResult := doc.Find("#list li").Last().Text()

	if noResult == "暂无结果" {
		address = strings.TrimSpace(address)
		s.renderTable(ip, address, []string{}, []string{})
	} else {
		var bindtimes []string
		var bindsites []string
		address = strings.TrimSpace(address)

		doc.Find("#list li").Each(func(i int, s *goquery.Selection) {
			if i < 2 {
				return
			}
			date := s.Find(".date").Text()
			site := s.Find("a").Text()
			if date != "" && site != "" {
				bindtimes = append(bindtimes, strings.TrimSpace(date))
				bindsites = append(bindsites, strings.TrimSpace(site))
			}
		})

		s.renderTable(ip, address, bindsites, bindtimes)
	}

	return nil
}

// renderTable 渲染IP查询结果表格
func (s *IPService) renderTable(ip string, address string, bindSites []string, bindTimes []string) {
	basicHeader := prettytable.Row{"项目", "值"}
	basicTable := table.Tables(os.Stdout, 200, fmt.Sprintf("IP %s 查询结果", ip), basicHeader)

	basicTable.AppendRow(table.ColorData(prettytable.Row{"IP地址", ip}))
	basicTable.AppendRow(table.ColorData(prettytable.Row{"IP归属地", address}))

	basicTable.Render()

	if len(bindSites) > 0 {
		bindHeader := prettytable.Row{"绑定网站", "绑定时间"}
		bindTable := table.Tables(os.Stdout, 200, "绑定信息", bindHeader)

		for i := range bindSites {
			bindTable.AppendRow(table.ColorData(prettytable.Row{bindSites[i], bindTimes[i]}))
		}

		bindTable.Render()
	} else {
		logger.Log.InfoMsgf("未查询到相关绑定信息！")
	}
}
