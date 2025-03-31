package model

import "errors"

var (
	ErrEmptyTitle      = errors.New("标题不能为空")
	ErrInvalidPriority = errors.New("优先级必须在1-5之间")
)
