package commands

import (
	"fmt"
	"strconv"
	"time"

	"todoz/internal/model"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [id]",
	Short: "编辑待办事项",
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

		// 获取要更新的字段
		if title, _ := cmd.Flags().GetString("title"); title != "" {
			todo.Title = title
		}
		if description, _ := cmd.Flags().GetString("description"); description != "" {
			todo.Description = description
		}
		if topic, _ := cmd.Flags().GetString("topic"); topic != "" {
			todo.Topic = topic
		}
		if priority, _ := cmd.Flags().GetInt("priority"); priority > 0 {
			todo.Priority = model.Priority(priority)
		}
		if dueDateStr, _ := cmd.Flags().GetString("due-date"); dueDateStr != "" {
			dueDate, err := time.Parse("2006-01-02", dueDateStr)
			if err != nil {
				fmt.Printf("无效的日期格式: %v\n", err)
				return
			}
			todo.DueDate = dueDate
		}

		if err := todo.Validate(); err != nil {
			fmt.Printf("无效的待办事项: %v\n", err)
			return
		}

		if err := db.Update(todo); err != nil {
			fmt.Printf("更新待办事项失败: %v\n", err)
			return
		}

		fmt.Printf("成功更新待办事项 (ID: %d)\n", todo.ID)
	},
}

// RegisterEditCmd 注册编辑命令
func RegisterEditCmd(rootCmd *cobra.Command) {
	// 添加标志
	editCmd.Flags().String("title", "", "新的标题")
	editCmd.Flags().StringP("description", "d", "", "新的描述")
	editCmd.Flags().StringP("topic", "t", "", "新的主题")
	editCmd.Flags().IntP("priority", "p", 0, "新的优先级 (1-5)")
	editCmd.Flags().String("due-date", "", "新的截止日期 (YYYY-MM-DD)")

	rootCmd.AddCommand(editCmd)
}
