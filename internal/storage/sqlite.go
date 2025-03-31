package storage

import (
	"database/sql"
	"time"

	"todoz/internal/model"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStorage struct {
	db *sql.DB
}

func NewSQLiteStorage(dbPath string) (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	storage := &SQLiteStorage{db: db}
	if err := storage.initDB(); err != nil {
		return nil, err
	}

	return storage, nil
}

func (s *SQLiteStorage) initDB() error {
	query := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		topic TEXT,
		priority INTEGER NOT NULL,
		due_date DATETIME NOT NULL,
		completed BOOLEAN NOT NULL DEFAULT 0,
		in_progress BOOLEAN NOT NULL DEFAULT 0,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL
	);
	`
	_, err := s.db.Exec(query)
	return err
}

func (s *SQLiteStorage) Close() error {
	return s.db.Close()
}

func (s *SQLiteStorage) Create(todo *model.Todo) error {
	if todo.InProgress {
		if err := s.clearInProgress(); err != nil {
			return err
		}
	}

	query := `
		INSERT INTO todos (
			title, description, topic, priority, due_date, 
			completed, in_progress, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	result, err := s.db.Exec(query,
		todo.Title, todo.Description, todo.Topic, todo.Priority,
		todo.DueDate, todo.Completed, todo.InProgress,
		todo.CreatedAt, todo.UpdatedAt,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	todo.ID = id
	return nil
}

func (s *SQLiteStorage) Get(id int64) (*model.Todo, error) {
	query := `
	SELECT id, title, description, topic, priority, due_date, completed, created_at, updated_at
	FROM todos WHERE id = ?;
	`
	todo := &model.Todo{}
	err := s.db.QueryRow(query, id).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Description,
		&todo.Topic,
		&todo.Priority,
		&todo.DueDate,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *SQLiteStorage) Update(todo *model.Todo) error {
	if todo.InProgress {
		if err := s.clearInProgress(); err != nil {
			return err
		}
	}

	query := `
		UPDATE todos SET
			title = ?, description = ?, topic = ?, priority = ?,
			due_date = ?, completed = ?, in_progress = ?, updated_at = ?
		WHERE id = ?
	`
	_, err := s.db.Exec(query,
		todo.Title, todo.Description, todo.Topic, todo.Priority,
		todo.DueDate, todo.Completed, todo.InProgress,
		time.Now(), todo.ID,
	)
	return err
}

func (s *SQLiteStorage) Delete(id int64) error {
	query := `DELETE FROM todos WHERE id = ?;`
	_, err := s.db.Exec(query, id)
	return err
}

func (s *SQLiteStorage) List() ([]*model.Todo, error) {
	query := `
	SELECT id, title, description, topic, priority, due_date, completed, created_at, updated_at
	FROM todos ORDER BY created_at DESC;
	`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*model.Todo
	for rows.Next() {
		todo := &model.Todo{}
		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Description,
			&todo.Topic,
			&todo.Priority,
			&todo.DueDate,
			&todo.Completed,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (s *SQLiteStorage) GetInProgress() (*model.Todo, error) {
	var todo model.Todo
	query := `
		SELECT id, title, description, topic, priority, due_date,
			completed, in_progress, created_at, updated_at
		FROM todos WHERE in_progress = 1
		LIMIT 1
	`
	err := s.db.QueryRow(query).Scan(
		&todo.ID, &todo.Title, &todo.Description, &todo.Topic,
		&todo.Priority, &todo.DueDate, &todo.Completed, &todo.InProgress,
		&todo.CreatedAt, &todo.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (s *SQLiteStorage) SetInProgress(id int64) error {
	if err := s.clearInProgress(); err != nil {
		return err
	}

	query := `UPDATE todos SET in_progress = 1 WHERE id = ?`
	_, err := s.db.Exec(query, id)
	return err
}

func (s *SQLiteStorage) clearInProgress() error {
	query := `UPDATE todos SET in_progress = 0`
	_, err := s.db.Exec(query)
	return err
}

func (s *SQLiteStorage) AutoSetInProgress() error {
	if err := s.clearInProgress(); err != nil {
		return err
	}

	query := `
		SELECT id FROM todos
		WHERE completed = 0
		ORDER BY 
			CASE 
				WHEN due_date < date('now') THEN 1
				ELSE 0
			END,
			priority DESC,
			due_date ASC
		LIMIT 1
	`
	var id int64
	err := s.db.QueryRow(query).Scan(&id)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		return err
	}

	return s.SetInProgress(id)
}
