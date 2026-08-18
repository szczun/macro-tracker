package models

import (
	"context"
	"database/sql"
	"time"
)

type User struct {
	ID            int
	Name          string
	Age           int
	Weight        float64
	Height        int
	ActivityLevel uint8
	Created       time.Time
	TDEE          int
}

type UserModel struct {
	DB *sql.DB
}

func NewUserModel(db *sql.DB) *UserModel {
	return &UserModel{DB: db}
}

type UserRepository interface {
	Insert(ctx context.Context, u *User) (int64, error)
}

func (m *UserModel) Insert(ctx context.Context, u *User) (int64, error) {
	query := `INSERT INTO users (name, age, weight, height, activity_level, created, tdee)
	VALUES(?, ?, ?, ?, ?, UTC_TIMESTAMP(), ?)`

	result, err := m.DB.ExecContext(ctx, query, u)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
