package commands

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "删除待办事项",
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

		if err := db.Delete(id); err != nil {
			fmt.Printf("删除待办事项失败: %v\n", err)
			return
		}

		fmt.Printf("成功删除待办事项 (ID: %d)\n", id)
	},
}

// RegisterDeleteCmd 注册删除命令
func RegisterDeleteCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(deleteCmd)
}
