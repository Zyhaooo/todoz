.PHONY: build run test clean

# 构建项目
build:
	go build -o todoz cmd/todoz/main.go

# 运行项目
run: build
	./todoz list

# 运行测试
test:
	go test ./...

# 清理构建文件
clean:
	rm -f todoz

# 安装依赖
deps:
	go mod tidy

# 开发模式运行
dev:
	go run cmd/todoz/main.go list 