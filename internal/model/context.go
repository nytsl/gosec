package model

// QueryContext 查询上下文 - 包含临时配置
type QueryContext struct {
	Proxy   string // HTTP代理地址
	Timeout int    // 超时时间（秒）
	// 其他临时配置可以在这里添加
}

// NewQueryContext 创建查询上下文
func NewQueryContext() *QueryContext {
	return &QueryContext{}
}

// WithProxy 设置代理
func (ctx *QueryContext) WithProxy(proxy string) *QueryContext {
	ctx.Proxy = proxy
	return ctx
}

// WithTimeout 设置超时
func (ctx *QueryContext) WithTimeout(timeout int) *QueryContext {
	ctx.Timeout = timeout
	return ctx
}

// HasProxy 检查是否有代理配置
func (ctx *QueryContext) HasProxy() bool {
	return ctx.Proxy != ""
}
