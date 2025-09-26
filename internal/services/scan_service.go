package services

import (
	"awesomeProject/pkg/logger"
	"fmt"
	"net"
	"time"
)

// ScanService 网络扫描服务
type ScanService struct {
	timeout time.Duration
}

// NewScanService 创建新的扫描服务
func NewScanService() *ScanService {
	return &ScanService{
		timeout: 3 * time.Second,
	}
}

// SetTimeout 设置扫描超时时间
func (s *ScanService) SetTimeout(timeout time.Duration) {
	s.timeout = timeout
}

// ScanPort 扫描单个端口
func (s *ScanService) ScanPort(host string, port int) error {
	logger.Log.InfoMsgf("正在扫描 %s:%d", host, port)

	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, s.timeout)
	if err != nil {
		logger.Log.WarnMsgf("端口 %d 关闭或不可达", port)
		return err
	}
	defer conn.Close()

	logger.Log.InfoMsgf("端口 %d 开放", port)
	return nil
}

// ScanPortRange 扫描端口范围
func (s *ScanService) ScanPortRange(host string, startPort, endPort int) error {
	logger.Log.InfoMsgf("正在扫描 %s 的端口范围 %d-%d", host, startPort, endPort)

	var openPorts []int
	for port := startPort; port <= endPort; port++ {
		if err := s.ScanPort(host, port); err == nil {
			openPorts = append(openPorts, port)
		}
	}

	if len(openPorts) > 0 {
		logger.Log.InfoMsgf("发现开放端口: %v", openPorts)
	} else {
		logger.Log.InfoMsgf("在指定范围内未发现开放端口")
	}

	return nil
}
