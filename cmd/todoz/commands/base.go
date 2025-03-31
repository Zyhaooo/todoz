package commands

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"todoz/internal/storage"
)

var (
	dbPath string
	db     storage.Storage
)

// InitDB 初始化数据库连接
func InitDB() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("无法获取用户主目录: %v", err)
	}

	defaultDBPath := filepath.Join(homeDir, ".todoz", "todos.db")
	flag.StringVar(&dbPath, "db", defaultDBPath, "数据库文件路径")

	// 确保数据库目录存在
	dbDir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		return fmt.Errorf("无法创建数据库目录: %v", err)
	}

	// 初始化数据库连接
	var dbErr error
	db, dbErr = storage.NewSQLiteStorage(dbPath)
	if dbErr != nil {
		return fmt.Errorf("无法初始化数据库: %v", dbErr)
	}

	return nil
}

// CloseDB 关闭数据库连接
func CloseDB() error {
	if err := db.Close(); err != nil {
		return fmt.Errorf("关闭数据库连接时出错: %v", err)
	}
	return nil
}
