package commands

import (
	"fmt"

	"todoz/internal/display"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有待办事项",
	Run: func(cmd *cobra.Command, args []string) {
		// 获取所有待办事项
		todos, err := db.List()
		if err != nil {
			fmt.Printf("获取待办事项列表失败: %v\n", err)
			return
		}

		// 直接输出表格
		fmt.Println(display.RenderTodoListTable(todos, display.DefaultConfig()))
	},
}

// RegisterListCmd 注册列表命令
func RegisterListCmd(rootCmd *cobra.Command) {
	// 添加标志
	listCmd.Flags().StringP("topic", "t", "", "按主题筛选")
	listCmd.Flags().IntP("priority", "p", 0, "按优先级筛选")
	listCmd.Flags().String("sort-by", "created", "排序方式 (created, due, priority)")
	listCmd.Flags().Bool("incomplete", false, "只显示未完成的待办事项")

	rootCmd.AddCommand(listCmd)
}
