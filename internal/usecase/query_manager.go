package usecase

import (
	"awesomeProject/internal/services"
	"awesomeProject/pkg/config"
	"awesomeProject/pkg/logger"
)

// 基础查询接口
type QueryService interface {
	Query(string) error
}

// QueryManager 查询管理器 - 统一管理所有查询服务
// 职责：协调多个底层查询服务；不负责具体解析实现。
type QueryManager struct {
	icpService        *services.ICPService
	ipService         *services.IPService
	whoisService      *services.WhoisService
	subdomainsService *services.SubdomainsService
}

// NewQueryManager 创建查询管理器
func NewQueryManager() *QueryManager {
	return &QueryManager{
		icpService:        services.NewICPService(),
		ipService:         services.NewIPService(),
		whoisService:      services.NewWhoisService(),
		subdomainsService: services.NewSubdomainsService(),
	}
}

// SetProxy 设置全局代理
func (qm *QueryManager) SetProxy(proxy string) {
	config.SetGlobalProxy(proxy)
}

// batchQuery 统一的批量处理实现
func (qm *QueryManager) batchQuery(service QueryService, targets []string, serviceType string) error {
	for _, target := range targets {
		if err := service.Query(target); err != nil {
			logger.Log.ErrorMsgf("%s查询失败 %s: %v", serviceType, target, err)
		}
	}
	return nil
}

// QueryICP 查询ICP备案信息
func (qm *QueryManager) QueryICP(domains []string) error {
	if len(domains) == 1 {
		return qm.icpService.Query(domains[0])
	}
	return qm.batchQuery(qm.icpService, domains, "ICP")
}

// QueryIP 查询IP反查信息
func (qm *QueryManager) QueryIP(ips []string) error {
	if len(ips) == 1 {
		return qm.ipService.Query(ips[0])
	}
	return qm.batchQuery(qm.ipService, ips, "IP")
}

// QueryWhois 查询Whois信息
func (qm *QueryManager) QueryWhois(domains []string) error {
	if len(domains) == 1 {
		return qm.whoisService.Query(domains[0])
	}
	return qm.batchQuery(qm.whoisService, domains, "Whois")
}

func (qm *QueryManager) QuerySubdomains(domains []string) error {
	if len(domains) == 1 {
		return qm.subdomainsService.Query(domains[0])
	}
	return qm.batchQuery(qm.subdomainsService, domains, "Subdomains")
}
