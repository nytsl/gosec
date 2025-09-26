package config

import (
	"gopkg.in/yaml.v3"
	"io/fs"
	"os"
	"path/filepath"
)

type Config struct {
	Cookies struct {
		BeiAn string `yaml:"BeiAn"`
	} `yaml:"Cookies"`
}

var (
	CfgName = "config.yaml"
	CfgYaml = `Cookies:
  BeiAn: ""`
	Version = "1.0.0"
)

// LoadConfig 加载配置文件
func LoadConfig() (*Config, error) {
	var cfg Config

	// 检查配置文件是否存在
	if _, err := os.Stat(CfgName); os.IsNotExist(err) {
		// 如果不存在，创建默认配置文件
		if err := createDefaultConfig(); err != nil {
			return nil, err
		}
	}

	// 读取配置文件
	data, err := os.ReadFile(CfgName)
	if err != nil {
		return nil, err
	}

	// 解析配置文件
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// createDefaultConfig 创建默认配置文件
func createDefaultConfig() error {
	// 确保目录存在
	dir := filepath.Dir(CfgName)
	if dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	// 创建配置文件
	return os.WriteFile(CfgName, []byte(CfgYaml), fs.FileMode(0644))
}

// SaveConfig 保存配置文件
func SaveConfig(cfg *Config) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(CfgName, data, fs.FileMode(0644))
}
