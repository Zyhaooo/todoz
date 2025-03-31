package main

import (
	"fmt"

	"todoz/cmd/todoz/commands"

	"github.com/spf13/cobra"
)

func main() {
	// 初始化数据库
	if err := commands.InitDB(); err != nil {
		fmt.Printf("初始化数据库失败: %v\n", err)
		return
	}
	defer commands.CloseDB()

	rootCmd := &cobra.Command{
		Use:   "todoz",
		Short: "一个简单的终端待办事项管理工具",
		Long: `TodoZ 是一个使用 Go 语言开发的终端待办事项管理工具。
它支持创建、查看、编辑和删除待办事项，并提供简洁的 TUI 界面。`,
	}

	// 注册所有命令
	commands.RegisterAddCmd(rootCmd)
	commands.RegisterListCmd(rootCmd)
	commands.RegisterShowCmd(rootCmd)
	commands.RegisterEditCmd(rootCmd)
	commands.RegisterCompleteCmd(rootCmd)
	commands.RegisterDeleteCmd(rootCmd)
	commands.RegisterCurrentCmd(rootCmd)
	commands.RegisterSetCurrentCmd(rootCmd)
	commands.RegisterAutoCurrentCmd(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
