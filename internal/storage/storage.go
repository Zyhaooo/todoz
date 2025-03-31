package storage

import "todoz/internal/model"

// Storage 定义了存储接口
type Storage interface {
	// 基本 CRUD 操作
	Create(todo *model.Todo) error
	Get(id int64) (*model.Todo, error)
	List() ([]*model.Todo, error)
	Update(todo *model.Todo) error
	Delete(id int64) error
	Close() error

	// 进行中状态相关操作
	GetInProgress() (*model.Todo, error)
	SetInProgress(id int64) error
	AutoSetInProgress() error
}
