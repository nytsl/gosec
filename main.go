package main

import (
	"awesomeProject/cmd"
	"awesomeProject/pkg/logger"
	"log/slog"
)

func main() {
	// 初始化日志
	logger.Init(
		logger.WithLevel(slog.LevelDebug),
		logger.WithTimeFormat("2006-01-02 15:04:05"),
		logger.WithOutputJson(false),
		logger.WithUseColor(true),
	)

	// 执行根命令
	cmd.Execute()
}
