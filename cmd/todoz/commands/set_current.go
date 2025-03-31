package commands

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var setCurrentCmd = &cobra.Command{
	Use:   "set-current [id]",
	Short: "设置指定待办事项为进行中",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("请提供待办事项 ID")
			return
		}

		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println("无效的 ID")
			return
		}

		// 检查待办事项是否存在
		todo, err := db.Get(id)
		if err != nil {
			fmt.Printf("获取待办事项失败: %v\n", err)
			return
		}

		if todo.Completed {
			fmt.Println("无法将已完成的待办事项设置为进行中")
			return
		}

		if err := db.SetInProgress(id); err != nil {
			fmt.Printf("设置进行中状态失败: %v\n", err)
			return
		}

		fmt.Printf("成功设置待办事项为进行中 (ID: %d)\n", id)
	},
}

// RegisterSetCurrentCmd 注册设置当前命令
func RegisterSetCurrentCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(setCurrentCmd)
}
