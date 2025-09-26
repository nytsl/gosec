package services

import (
	"awesomeProject/pkg/logger"

	"github.com/likexian/whois"
)

// WhoisService Whois查询服务
type WhoisService struct{}

// NewWhoisService 创建Whois查询服务
func NewWhoisService() *WhoisService {
	return &WhoisService{}
}

// Query 查询Whois信息（Whois不支持代理）
func (s *WhoisService) Query(domain string) error {
	logger.Log.InfoMsgf("正在查询域名 %s 的Whois信息", domain)

	result, err := whois.Whois(domain)
	if err != nil {
		return logger.Log.ErrorMsgf("Whois查询错误:%s", err)
	}

	logger.Log.InfoMsgf("Whois查询结果：\n%s", result)
	return nil
}
