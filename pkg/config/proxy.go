package config

// GlobalProxy 全局代理配置
var GlobalProxy string

// SetGlobalProxy 设置全局代理
func SetGlobalProxy(proxy string) {
	GlobalProxy = proxy
}

// GetGlobalProxy 获取全局代理
func GetGlobalProxy() string {
	return GlobalProxy
}

// HasGlobalProxy 检查是否设置了代理
func HasGlobalProxy() bool {
	return GlobalProxy != ""
}
