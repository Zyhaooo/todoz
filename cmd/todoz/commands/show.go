package commands

import (
	"fmt"
	"strconv"
	"strings"

	"todoz/internal/display"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show [id]",
	Short: "显示特定待办事项的详细信息",
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

		// 使用 display 包渲染表格
		fmt.Println(display.RenderTodoTable(todo, display.DefaultConfig()))

		// 如果有详细描述，使用 glamour 渲染 Markdown
		if todo.Description != "" {
			fmt.Println("\n详细描述:")
			fmt.Println(strings.Repeat("-", 80))

			// 创建 glamour 渲染器
			r, err := glamour.NewTermRenderer(
				glamour.WithAutoStyle(),
				glamour.WithWordWrap(80),
			)
			if err != nil {
				fmt.Printf("创建渲染器失败: %v\n", err)
				return
			}

			// 渲染 Markdown
			out, err := r.Render(todo.Description)
			if err != nil {
				fmt.Printf("渲染 Markdown 失败: %v\n", err)
				return
			}

			fmt.Println(out)
			fmt.Println(strings.Repeat("-", 80))
		}
	},
}

// RegisterShowCmd 注册显示命令
func RegisterShowCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(showCmd)
}
