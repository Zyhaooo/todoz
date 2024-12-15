package main

import (
	"todoz/args"
	"todoz/model"
	"todoz/print"
)

var todos = []model.Todo{
	{
		Id:             1,
		Title:          "this now todo",
		Description:    "now todo is a goods todo",
		Level:          99,
		CreatedTime:    "2024-12-15",
		ExpirationTime: "2024-12-16",
		Finish:         false,
	},
	{
		Id:             2,
		Title:          "this not now todo",
		Description:    "not now todo is a goods todo",
		Level:          1,
		CreatedTime:    "2024-12-15",
		ExpirationTime: "2024-12-16",
		Finish:         true,
	},
	{
		Id:             3,
		Title:          "这是一段中文",
		Description:    "中文 todo",
		Level:          1,
		CreatedTime:    "2024-12-15",
		ExpirationTime: "2024-12-16",
		Finish:         false,
	},
}

func main() {
	var (
		as args.Args
	)

	as = args.Parse()

	switch as.Action {
	case args.ADD:
	case args.FINISH:
	case args.LIST:
		print.PrintList(todos)
	default:
		print.PrintSingle(todos[0])
	}
}
