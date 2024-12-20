package args

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
	"todoz/model"
)

// action
const (
	_             = iota
	ADD    Action = iota // 1
	FINISH Action = iota // 2
	LIST   Action = iota // 3
)

type Action int

type Args struct {
	Action Action
	model.Todo
	Page int
	Size int
}

var ActionMap = map[string]Action{
	"add":    ADD,
	"finish": FINISH,
	"list":   LIST,
}

func Parse() Args {
	var (
		f        *flag.FlagSet
		args     = Args{}
		duration time.Duration
		//format   = "2006-01-02-15"
	)

	// 初始化自定义flag
	f = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	switch len(os.Args) {
	case 1:
		return args
	default:

		if action, ok := ActionMap[strings.ToLower(os.Args[1])]; ok {
			args.Action = action
		} else {
			fmt.Println("unkown action !")
			os.Exit(1)
		}

		switch args.Action {
		case ADD:
			f.StringVar(&args.Title, "title", "", "todo title")
			f.StringVar(&args.Description, "description", "", "todo description")
			f.IntVar(&args.Level, "level", 1, "todo level")
			f.DurationVar(&duration, "time", time.Hour, "todo expiration time")

			f.Parse(os.Args[2:])

		case FINISH:
			f.Uint64Var(&args.Id, "id", 0, "finish todo id ; if not provide , finish now todo !")
			f.Parse(os.Args[2:])
		case LIST:
			f.IntVar(&args.Page, "page", 1, "list page")
			f.IntVar(&args.Size, "size", 10, "list page size")
			f.Parse(os.Args[2:])
		}

	}

	if args.Action == ADD {
		args.CreatedTime = time.Now().Format(time.RFC3339)
		args.ExpirationTime = time.Now().Add(duration).Format(time.RFC3339)
	}

	return args
}
