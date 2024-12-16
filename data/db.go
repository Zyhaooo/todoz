package data

import (
	"database/sql"
	"fmt"
	"os"
	"todoz/model"

	_ "modernc.org/sqlite"
)

var db *sql.DB

var err error

var (
	rows *sql.Rows
	row  *sql.Row
)

func init() {
	db, err = sql.Open("sqlite", "./todoz_db")
	if err != nil {
		fmt.Printf("sql open err !\n error : %+v", err)
		os.Exit(1)
	}

	_, err = db.Exec(todos)
	if err != nil {
		fmt.Printf("db create todos err !\n error : %+v", err)
		os.Exit(1)
	}

}

const insert = `
INSERT INTO [todos] ([title], [description], [level], [create_time], [expiration_time], [finish]) VALUES (?,?,?,?,?,?);
`

func AddTodo(todo model.Todo) {

	if _, err = db.Exec(insert,
		todo.Title,
		todo.Description,
		todo.Level,
		todo.CreatedTime,
		todo.ExpirationTime,
		todo.Finish,
	); err != nil {
		fmt.Printf("db insert err ! \n error : %+v", err)
		os.Exit(1)
	}

}

const finish = `
UPDATE [todos]
SET [finish]='1'
WHERE ([todos].[id] = ?);
`

func FinishTodo(id uint64) {

	if _, err = db.Exec(finish,
		id,
	); err != nil {
		fmt.Printf("db finish todo err ! \nerror : %+v", err)
		os.Exit(1)
	}
}

const getCurrentTodo = `
SELECT * FROM [todos]
WHERE [expiration_time] >= datetime('now', 'localtime') AND [finish] = false
ORDER BY level desc LIMIT 1;
`

func GetCurrentTodo() (todo model.Todo) {

	row = db.QueryRow(getCurrentTodo)
	return scanTodo(row)
}

const getTodoList = `
SELECT * FROM [todos]
ORDER BY level desc
LIMIT ?
OFFSET ?;
`

func GetTodoList(page int, size int) (todos []model.Todo) {
	if rows, err = db.Query(getTodoList, size, (page-1)*size); err != nil {
		fmt.Printf("db query todo list err !\nerr : %+v\n", err)
		os.Exit(1)
	}

	todos = make([]model.Todo, 0)

	for rows.Next() {
		var todo model.Todo
		if err = rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Level, &todo.CreatedTime, &todo.ExpirationTime, &todo.Finish); err != nil {
			fmt.Printf("row scan err !\nerr : %+v\n", err)
			os.Exit(1)
		}
		todos = append(todos, todo)
	}

	return
}

func scanTodo(row *sql.Row) (todo model.Todo) {
	if err = row.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Level, &todo.CreatedTime, &todo.ExpirationTime, &todo.Finish); err != nil {
		fmt.Printf("row scan err !\nerr : %+v", err)
		os.Exit(1)
	}
	return todo
}
