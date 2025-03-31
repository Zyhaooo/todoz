package display

import (
	"bytes"
	"fmt"
	"strings"

	"todoz/internal/model"

	"github.com/olekukonko/tablewriter"
)

// TableConfig 定义表格的配置选项
type TableConfig struct {
	ShowBorders       bool
	ShowHeaders       bool
	AutoWrapText      bool
	AutoFormatHeaders bool
	HeaderAlignment   int
	Alignment         int
}

// DefaultConfig 返回默认的表格配置
func DefaultConfig() TableConfig {
	return TableConfig{
		ShowBorders:       true,
		ShowHeaders:       true,
		AutoWrapText:      false,
		AutoFormatHeaders: true,
		HeaderAlignment:   tablewriter.ALIGN_LEFT,
		Alignment:         tablewriter.ALIGN_LEFT,
	}
}

// RenderTodoTable 渲染单个待办事项的表格
func RenderTodoTable(todo *model.Todo, config TableConfig) string {
	var buf bytes.Buffer
	table := tablewriter.NewWriter(&buf)

	// 设置表格样式
	table.SetBorders(tablewriter.Border{
		Left:   config.ShowBorders,
		Top:    config.ShowBorders,
		Right:  config.ShowBorders,
		Bottom: config.ShowBorders,
	})
	table.SetCenterSeparator("|")
	table.SetColumnSeparator("|")
	table.SetRowSeparator("-")
	table.SetAutoWrapText(config.AutoWrapText)
	table.SetAutoFormatHeaders(config.AutoFormatHeaders)
	table.SetHeaderAlignment(config.HeaderAlignment)
	table.SetAlignment(config.Alignment)

	// 设置表头
	if config.ShowHeaders {
		table.SetHeader([]string{"字段", "内容"})
	}

	// 添加数据行
	table.Append([]string{"ID", fmt.Sprintf("%d", todo.ID)})
	table.Append([]string{"标题", todo.Title})
	table.Append([]string{"主题", todo.Topic})
	table.Append([]string{"优先级", strings.Repeat("*", int(todo.Priority))})
	table.Append([]string{"截止日期", todo.DueDate.Format("2006-01-02")})
	table.Append([]string{"状态", getStatus(todo.Completed)})
	table.Append([]string{"创建时间", todo.CreatedAt.Format("2006-01-02 15:04:05")})
	table.Append([]string{"更新时间", todo.UpdatedAt.Format("2006-01-02 15:04:05")})

	table.Render()
	return buf.String()
}

// RenderTodoListTable 渲染待办事项列表的表格
func RenderTodoListTable(todos []*model.Todo, config TableConfig) string {
	var buf bytes.Buffer
	table := tablewriter.NewWriter(&buf)

	// 设置表格样式
	table.SetBorders(tablewriter.Border{
		Left:   config.ShowBorders,
		Top:    config.ShowBorders,
		Right:  config.ShowBorders,
		Bottom: config.ShowBorders,
	})
	table.SetCenterSeparator("|")
	table.SetColumnSeparator("|")
	table.SetRowSeparator("-")
	table.SetAutoWrapText(config.AutoWrapText)
	table.SetAutoFormatHeaders(config.AutoFormatHeaders)
	table.SetHeaderAlignment(config.HeaderAlignment)
	table.SetAlignment(config.Alignment)

	// 设置表头
	if config.ShowHeaders {
		table.SetHeader([]string{"ID", "标题", "主题", "优先级", "截止日期", "状态"})
	}

	// 添加数据行
	for _, todo := range todos {
		table.Append([]string{
			fmt.Sprintf("%d", todo.ID),
			todo.Title,
			todo.Topic,
			strings.Repeat("*", int(todo.Priority)),
			todo.DueDate.Format("2006-01-02"),
			getStatus(todo.Completed),
		})
	}

	table.Render()
	return buf.String()
}

// getStatus 返回待办事项的状态显示文本
func getStatus(completed bool) string {
	if completed {
		return "已完成"
	}
	return "未完成"
}
