# TodoZ - 终端待办事项管理工具

TodoZ 是一个简单但功能强大的终端待办事项管理工具，使用 Go 语言开发，支持 SQLite 持久化存储。

## 功能特点

- 创建、查看、编辑、删除待办事项
- 支持 Markdown 格式的详细描述
- 待办事项属性：
  - ID（自动生成）
  - 标题
  - 详细描述（支持 Markdown）
  - 主题分类
  - 重要程度（1-5）
  - 截止日期
  - 完成状态
- 支持按不同条件筛选和排序待办事项
- 简洁的 TUI 界面

## 安装

```bash
go install github.com/yourusername/todoz@latest
```

## 使用方法

### 基本命令

```bash
# 添加新的待办事项
todoz add "完成项目文档" -d "需要编写详细的项目文档，包括..." -t "工作" -p 3 -d 2024-04-01

# 列出所有待办事项
todoz list

# 查看特定待办事项
todoz show 1

# 编辑待办事项
todoz edit 1 -t "新标题"

# 标记待办事项为完成
todoz complete 1

# 删除待办事项
todoz delete 1
```

### 筛选和排序

```bash
# 按主题筛选
todoz list -t "工作"

# 按重要程度筛选
todoz list -p 3

# 按截止日期排序
todoz list --sort-by due

# 只显示未完成的待办事项
todoz list --incomplete
```

## 开发

```bash
# 克隆仓库
git clone https://github.com/yourusername/todoz.git

# 安装依赖
go mod tidy

# 运行测试
go test ./...

# 构建
go build
```

## 许可证

MIT License 