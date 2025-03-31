package commands

import (
	"fmt"
	"strings"

	"todoz/internal/display"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "显示当前进行中的待办事项",
	Run: func(cmd *cobra.Command, args []string) {
		todo, err := db.GetInProgress()
		if err != nil {
			fmt.Printf("获取进行中的待办事项失败: %v\n", err)
			return
		}

		if todo == nil {
			fmt.Println("当前没有进行中的待办事项")
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

// RegisterCurrentCmd 注册当前命令
func RegisterCurrentCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(currentCmd)
}
