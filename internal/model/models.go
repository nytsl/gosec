package model

// ICPResult ICP备案查询结果
type ICPResult struct {
	OrganizationName  string `json:"organization_name"`  // 主办单位名称
	OrganizationType  string `json:"organization_type"`  // 主办单位性质
	ICPNumber         string `json:"icp_number"`         // 备案号
	WebsiteName       string `json:"website_name"`       // 网站名称
	HomepageURL       string `json:"homepage_url"`       // 网站首页网址
	ReviewDate        string `json:"review_date"`        // 审核时间
	AccessRestriction string `json:"access_restriction"` // 限制访问情况
}

// IPResult IP反查结果
type IPResult struct {
	Address   string   `json:"address"`    // IP地址
	BindTimes []string `json:"bind_times"` // 绑定时间
	BindSites []string `json:"bind_sites"` // 绑定的网站
}

// WhoisResult Whois查询结果
type WhoisResult struct {
	Domain     string `json:"domain"`      // 域名
	Registrar  string `json:"registrar"`   // 注册商
	CreateDate string `json:"create_date"` // 创建时间
	UpdateDate string `json:"update_date"` // 更新时间
	ExpireDate string `json:"expire_date"` // 过期时间
	Status     string `json:"status"`      // 状态
	NameServer string `json:"name_server"` // 名称服务器
}

// QueryRequest 查询请求
type QueryRequest struct {
	Target string `json:"target"` // 查询目标（域名或IP）
	Type   string `json:"type"`   // 查询类型（icp, ip, whois）
	Proxy  string `json:"proxy"`  // 代理地址
}

// ScanResult 扫描结果
type ScanResult struct {
	Host    string `json:"host"`    // 目标主机
	Port    int    `json:"port"`    // 端口号
	IsOpen  bool   `json:"is_open"` // 是否开放
	Service string `json:"service"` // 服务类型
	Banner  string `json:"banner"`  // 服务横幅
}
