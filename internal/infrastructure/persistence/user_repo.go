package persistence

import (
	"database/sql"
	"errors"
	"my-go-app/internal/domain/user"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteUserRepo struct {
	db *sql.DB
}

func NewSqliteUserRepo(dbPath string) (*SqliteUserRepo, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	// Create table if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id TEXT PRIMARY KEY,
        name TEXT NOT NULL,
        email TEXT NOT NULL UNIQUE,
        password_hash TEXT NOT NULL,
        created_at DATETIME NOT NULL
    )`)
	if err != nil {
		return nil, err
	}
	return &SqliteUserRepo{db: db}, nil
}

func (r *SqliteUserRepo) Save(u *user.User) error {
	_, err := r.db.Exec(`INSERT OR REPLACE INTO users (id, name, email, password_hash, created_at) VALUES (?, ?, ?, ?, ?)`,
		u.ID, u.Name, u.Email, u.PasswordHash, u.CreatedAt.Format(time.RFC3339))
	return err
}

func (r *SqliteUserRepo) FindByID(id user.UserID) (*user.User, error) {
	row := r.db.QueryRow(`SELECT id, name, email, password_hash, created_at FROM users WHERE id = ?`, id)
	return scanUser(row)
}

func (r *SqliteUserRepo) FindAll() ([]*user.User, error) {
	rows, err := r.db.Query(`SELECT id, name, email, password_hash, created_at FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*user.User
	for rows.Next() {
		u, err := scanUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *SqliteUserRepo) FindByEmail(email string) (*user.User, error) {
	row := r.db.QueryRow(`SELECT id, name, email, password_hash, created_at FROM users WHERE email = ?`, email)
	return scanUser(row)
}

func scanUser(scanner interface {
	Scan(dest ...any) error
}) (*user.User, error) {
	var u user.User
	var createdAt string
	err := scanner.Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash, &createdAt)
	if err != nil {
		return nil, errors.New("user not found")
	}
	t, err := time.Parse(time.RFC3339, createdAt)
	if err != nil {
		return nil, err
	}
	u.CreatedAt = t
	return &u, nil
}
