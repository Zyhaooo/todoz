package commands

import (
	"fmt"

	"todoz/internal/display"

	"github.com/spf13/cobra"
)

var autoCurrentCmd = &cobra.Command{
	Use:   "auto-current",
	Short: "自动设置最紧要的待办事项为进行中",
	Run: func(cmd *cobra.Command, args []string) {
		if err := db.AutoSetInProgress(); err != nil {
			fmt.Printf("自动设置进行中状态失败: %v\n", err)
			return
		}

		// 显示新设置的进行中待办事项
		todo, err := db.GetInProgress()
		if err != nil {
			fmt.Printf("获取进行中的待办事项失败: %v\n", err)
			return
		}

		if todo == nil {
			fmt.Println("没有可用的待办事项")
			return
		}

		fmt.Printf("已自动设置最紧要的待办事项为进行中 (ID: %d)\n", todo.ID)
		fmt.Println(display.RenderTodoTable(todo, display.DefaultConfig()))
	},
}

// RegisterAutoCurrentCmd 注册自动设置当前命令
func RegisterAutoCurrentCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(autoCurrentCmd)
}
