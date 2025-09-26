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

// ICPService ICP备案查询服务
type ICPService struct {
	config *config.Config
}

// NewICPService 创建ICP查询服务
func NewICPService() *ICPService {
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Log.WarnMsgf("加载配置文件失败，使用默认配置: %v", err)
		cfg = &config.Config{}
	}

	return &ICPService{
		config: cfg,
	}
}

// Query 查询单个域名的ICP备案信息
func (s *ICPService) Query(domain string) error {
	logger.Log.InfoMsgf("正在查询域名 %s 的ICP备案信息", domain)

	client := req.C()

	// 使用全局代理配置
	if proxy := config.GetGlobalProxy(); proxy != "" {
		client.SetProxyURL(proxy)
	}

	request := client.R().SetHeader("Cookie", fmt.Sprintf("machine_str=%s", s.config.Cookies.BeiAn))
	resp, err := request.Get("https://www.beianx.cn/search/" + domain)
	if err != nil {
		return logger.Log.ErrorMsgf("请求失败：%s", err)
	}

	if !resp.IsSuccessState() {
		return logger.Log.ErrorMsgf("请求失败，状态码为：%s", resp.Status)
	}

	if resp.StatusCode == 401 {
		return logger.Log.ErrorMsgf("Cookie失效，请重新获取")
	}

	logger.Log.InfoMsgf("[+]请求成功")

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return logger.Log.ErrorMsgf("Html解析失败：%s", err)
	}

	var allRows [][]string
	doc.Find("table tbody tr").Each(func(index int, selection *goquery.Selection) {
		var row []string
		for i := 1; i < 7; i++ {
			row = append(row, strings.TrimSpace(selection.Find("td").Eq(i).Text()))
		}
		allRows = append(allRows, row)
	})

	s.renderTable(allRows)
	return nil
}

// renderTable 渲染查询结果表格
func (s *ICPService) renderTable(allrows [][]string) {
	if len(allrows) > 0 {
		Header := prettytable.Row{"主办单位名称", "主办单位性质", "ICP备案号", "网站名称", "网站首页地址", "审核通过日期", "是否限制接入"}
		t := table.Tables(os.Stdout, 300, "ICP查询结果", Header)
		for _, row := range allrows {
			tableRow := make(prettytable.Row, len(row))
			for i, v := range row {
				tableRow[i] = v
			}
			t.AppendRow(table.ColorData(tableRow))
		}
		t.Render()
	}
}
