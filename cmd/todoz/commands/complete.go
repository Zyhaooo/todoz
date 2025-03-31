package commands

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete [id]",
	Short: "标记待办事项为已完成",
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

		todo, err := db.Get(id)
		if err != nil {
			fmt.Printf("获取待办事项失败: %v\n", err)
			return
		}

		todo.Completed = true
		if err := db.Update(todo); err != nil {
			fmt.Printf("更新待办事项失败: %v\n", err)
			return
		}

		fmt.Printf("成功标记待办事项为已完成 (ID: %d)\n", todo.ID)
	},
}

// RegisterCompleteCmd 注册完成命令
func RegisterCompleteCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(completeCmd)
}
