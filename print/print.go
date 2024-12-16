package print

import (
	"fmt"
	"os"
	"todoz/model"

	"github.com/olekukonko/tablewriter"
)

var table = tablewriter.NewWriter(os.Stdout)

func init() {
	table.SetHeader([]string{"F", "id", "title", "description", "level", "ExpirationTime", "CreatedTime"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
}

func PrintSingle(m model.Todo) {
	table.Append([]string{
		printFinish(m.Finish),
		fmt.Sprint(m.Id),
		m.Title,
		m.Description,
		fmt.Sprint(m.Level),
		m.ExpirationTime,
		m.CreatedTime,
	})

	table.Render()
}

func PrintList(ms []model.Todo) {
	for _, m := range ms {
		table.Append([]string{
			printFinish(m.Finish),
			fmt.Sprint(m.Id),
			m.Title,
			m.Description,
			fmt.Sprint(m.Level),
			m.ExpirationTime,
			m.CreatedTime,
		})
	}

	table.Render()
}

func printFinish(b bool) string {
	if b {
		return "x"
	}
	return " "
}
