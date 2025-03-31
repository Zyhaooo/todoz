package commands

import (
	"fmt"
	"os"
	"time"

	"todoz/internal/model"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "添加新的待办事项",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("请提供待办事项标题")
			return
		}

		title := args[0]
		description, _ := cmd.Flags().GetString("description")
		descriptionFile, _ := cmd.Flags().GetString("description-file")
		topic, _ := cmd.Flags().GetString("topic")
		priority, _ := cmd.Flags().GetInt("priority")
		dueDateStr, _ := cmd.Flags().GetString("due-date")

		// 如果提供了描述文件，则从文件读取描述
		if descriptionFile != "" {
			content, err := os.ReadFile(descriptionFile)
			if err != nil {
				fmt.Printf("读取描述文件失败: %v\n", err)
				return
			}
			description = string(content)
		}

		// 解析截止日期
		var dueDate time.Time
		var err error
		if dueDateStr != "" {
			dueDate, err = time.Parse("2006-01-02", dueDateStr)
			if err != nil {
				fmt.Printf("无效的日期格式: %v\n", err)
				return
			}
		}

		// 创建待办事项
		todo := model.NewTodo(title, description, topic, model.Priority(priority), dueDate)
		if err := todo.Validate(); err != nil {
			fmt.Printf("无效的待办事项: %v\n", err)
			return
		}

		// 保存到数据库
		if err := db.Create(todo); err != nil {
			fmt.Printf("创建待办事项失败: %v\n", err)
			return
		}

		fmt.Printf("成功创建待办事项 (ID: %d)\n", todo.ID)
		fmt.Printf("优先级: %d, 截止日期: %s\n", todo.Priority, todo.DueDate.Format("2006-01-02 15:04:05"))
	},
}

// RegisterAddCmd 注册添加命令
func RegisterAddCmd(rootCmd *cobra.Command) {
	// 添加标志
	addCmd.Flags().StringP("description", "d", "", "待办事项的详细描述")
	addCmd.Flags().String("description-file", "", "包含待办事项详细描述的文件路径（支持 Markdown 格式）")
	addCmd.Flags().StringP("topic", "t", "", "待办事项的主题")
	addCmd.Flags().IntP("priority", "p", 3, "待办事项的优先级 (1-5)")
	addCmd.Flags().String("due-date", "", "待办事项的截止日期 (YYYY-MM-DD)")

	rootCmd.AddCommand(addCmd)
}
