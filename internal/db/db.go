package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func (c Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.DBName,
		c.SSLMode,
	)
}

func New(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DSN())
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	return db, nil
}

// migrationLockID scopes a Postgres advisory lock so that concurrent
// replicas migrating at startup serialize instead of racing on DDL.
const migrationLockID = 918273645

func Migrate(db *sql.DB) error {
	ctx := context.Background()

	conn, err := db.Conn(ctx)
	if err != nil {
		return fmt.Errorf("acquire connection: %w", err)
	}
	defer conn.Close()

	if _, err := conn.ExecContext(ctx, `SELECT pg_advisory_lock($1)`, migrationLockID); err != nil {
		return fmt.Errorf("acquire migration lock: %w", err)
	}
	defer conn.ExecContext(ctx, `SELECT pg_advisory_unlock($1)`, migrationLockID)

	query := `
		CREATE TABLE IF NOT EXISTS tasks (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			description TEXT,
			done BOOLEAN NOT NULL DEFAULT FALSE
		);
	`

	if _, err := conn.ExecContext(ctx, query); err != nil {
		return fmt.Errorf("create tasks table: %w", err)
	}

	return nil
}

type Task struct {
	ID          int64
	Title       string
	Description string
	Done        bool
}

func ListTasks(db *sql.DB) ([]Task, error) {
	rows, err := db.Query(`SELECT id, title, description, done FROM tasks ORDER BY id ASC`)
	if err != nil {
		return nil, fmt.Errorf("list tasks: %w", err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Done); err != nil {
			return nil, fmt.Errorf("scan task: %w", err)
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate tasks: %w", err)
	}

	return tasks, nil
}

func GetTaskByID(db *sql.DB, id int64) (*Task, error) {
	var task Task
	if err := db.QueryRow(`SELECT id, title, description, done FROM tasks WHERE id = $1`, id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Done,
	); err != nil {
		return nil, fmt.Errorf("get task by id: %w", err)
	}
	return &task, nil
}

func CreateTask(db *sql.DB, title, description string, done bool) (*Task, error) {
	var task Task
	if err := db.QueryRow(
		`INSERT INTO tasks (title, description, done) VALUES ($1, $2, $3) RETURNING id, title, description, done`,
		title,
		description,
		done,
	).Scan(&task.ID, &task.Title, &task.Description, &task.Done); err != nil {
		return nil, fmt.Errorf("create task: %w", err)
	}
	return &task, nil
}

func UpdateTask(db *sql.DB, id int64, title, description string, done bool) (*Task, error) {
	var task Task
	if err := db.QueryRow(
		`UPDATE tasks SET title = $1, description = $2, done = $3 WHERE id = $4 RETURNING id, title, description, done`,
		title,
		description,
		done,
		id,
	).Scan(&task.ID, &task.Title, &task.Description, &task.Done); err != nil {
		return nil, fmt.Errorf("update task: %w", err)
	}
	return &task, nil
}

func DeleteTask(db *sql.DB, id int64) (bool, error) {
	result, err := db.Exec(`DELETE FROM tasks WHERE id = $1`, id)
	if err != nil {
		return false, fmt.Errorf("delete task: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("rows affected: %w", err)
	}

	return rowsAffected > 0, nil
}
