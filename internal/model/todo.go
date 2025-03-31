package model

import "time"

// Priority 表示待办事项的重要程度
type Priority int

const (
	PriorityLowest  Priority = 1
	PriorityLow     Priority = 2
	PriorityMedium  Priority = 3
	PriorityHigh    Priority = 4
	PriorityHighest Priority = 5
)

// GetDefaultDueDate 根据优先级返回默认的截止日期
func GetDefaultDueDate(priority Priority) time.Time {
	now := time.Now()
	switch priority {
	case PriorityHighest:
		return now.Add(2 * time.Hour) // 最高优先级：2小时
	case PriorityHigh:
		return now.Add(4 * time.Hour) // 高优先级：4小时
	case PriorityMedium:
		return now.Add(12 * time.Hour) // 中等优先级：12小时
	case PriorityLow:
		return now.Add(24 * time.Hour) // 低优先级：1天
	case PriorityLowest:
		return now.Add(7 * 24 * time.Hour) // 最低优先级：1周
	default:
		return now.Add(12 * time.Hour) // 默认：12小时
	}
}

// Todo 表示一个待办事项
type Todo struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Topic       string    `json:"topic"`
	Priority    Priority  `json:"priority"`
	DueDate     time.Time `json:"due_date"`
	Completed   bool      `json:"completed"`
	InProgress  bool      `json:"in_progress"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// NewTodo 创建一个新的待办事项
func NewTodo(title string, description string, topic string, priority Priority, dueDate time.Time) *Todo {
	now := time.Now()

	// 如果没有指定截止日期，根据优先级设置默认值
	if dueDate.IsZero() {
		dueDate = GetDefaultDueDate(priority)
	}

	return &Todo{
		Title:       title,
		Description: description,
		Topic:       topic,
		Priority:    priority,
		DueDate:     dueDate,
		Completed:   false,
		InProgress:  false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// Validate 验证待办事项的有效性
func (t *Todo) Validate() error {
	if t.Title == "" {
		return ErrEmptyTitle
	}
	if t.Priority < PriorityLowest || t.Priority > PriorityHighest {
		return ErrInvalidPriority
	}
	return nil
}
