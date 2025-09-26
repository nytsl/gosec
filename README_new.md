# awesomeProject

一个功能强大的域名信息查询工具，支持ICP备案信息、IP反查域名、Whois信息等查询功能。

## 🚀 快速开始

### 安装依赖
```bash
go mod tidy
```

### 编译运行
```bash
go build -o awesomeProject.exe
.\awesomeProject.exe --help
```

### 使用示例
```bash
# ICP备案查询
.\awesomeProject.exe info --icp baidu.com,taobao.com

# IP反查域名
.\awesomeProject.exe info --ip 8.8.8.8,1.1.1.1

# Whois查询
.\awesomeProject.exe info --whois example.com

# 使用代理（仅ICP和IP查询支持）
.\awesomeProject.exe info --icp baidu.com --proxy http://127.0.0.1:8080
```

## 📁 项目结构

```
awesomeProject/
├── cmd/                    # 命令行接口层
│   ├── root.go            # 根命令定义
│   ├── info.go            # 信息查询命令
│   ├── version.go         # 版本命令
│   └── config.go          # 配置管理命令
├── internal/              # 内部业务逻辑
│   ├── model/             # 数据模型
│   │   └── query.go       # 查询相关的数据结构
│   ├── services/          # 原子服务层
│   │   ├── icp_service.go # ICP备案查询服务
│   │   ├── ip_service.go  # IP反查服务
│   │   ├── whois_service.go # Whois查询服务
│   │   └── scan_service.go # 网络扫描服务
│   └── usecase/           # 用例编排层
│       └── query_manager.go # 服务管理器
├── pkg/                   # 公共库
│   ├── config/            # 配置管理
│   ├── logger/            # 日志系统
│   ├── table/             # 表格渲染
│   └── utils/             # 工具函数
├── main.go                # 程序入口点
├── config.yaml            # 应用配置文件
├── go.mod                 # Go模块定义
└── go.sum                 # 依赖版本锁定
```

## 🔧 开发指南

### 添加新的查询服务

1. **创建服务实现** (`internal/services/`)
   ```go
   type NewService struct {}

   // 实现 QueryService 接口
   func (s *NewService) Query(target string) error { ... }
   func (s *NewService) BatchQuery(targets []string) error { ... }

   // 如果需要代理支持，实现 Proxiable 接口
   func (s *NewService) SetProxy(proxy string) { ... }
   ```

2. **注册到 QueryManager** (`internal/usecase/query_manager.go`)
   - 添加服务字段
   - 在 `NewQueryManager` 中初始化
   - 添加对应的查询方法

3. **添加命令行参数** (`cmd/info.go`)
   - 在 `infoCmd.Flags()` 中添加新参数
   - 在 `RunE` 函数中处理新参数

### 架构原则

- **服务隔离**: 每个查询功能独立成服务，互不干扰
- **接口设计**: 使用 `QueryService` 基础接口 + `Proxiable` 可选能力接口
- **错误处理**: 批处理场景下记录错误但不中断其他查询
- **可扩展性**: 新增功能只需添加新服务，无需修改现有代码

### 依赖说明

- **CLI框架**: [Cobra](https://github.com/spf13/cobra) - 现代命令行应用框架
- **HTTP客户端**: [req](https://github.com/imroc/req) - 简洁的HTTP请求库  
- **HTML解析**: [goquery](https://github.com/PuerkitoBio/goquery) - jQuery风格的HTML解析
- **表格渲染**: [go-pretty](https://github.com/jedib0t/go-pretty) - 美观的表格输出
- **Whois查询**: [whois](https://github.com/likexian/whois) - Whois信息查询库

## 📝 贡献

欢迎提交 Issue 和 Pull Request 来改进项目！

---

> 该项目采用 Go 标准项目布局和分层架构设计，便于维护和扩展。
