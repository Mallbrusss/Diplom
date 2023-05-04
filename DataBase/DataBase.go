package database

import (
	"database/sql"
	"fmt"
	"time"

	tok "programm/Telegram/Token"
)

type User struct {
	ID           int
	TelegramID   int
	FirstName    string
	LastName     string
	UserName     string
	CreatedAt    time.Time
	LastSeenAt   time.Time
	Balance      float64
	IsAuthorized bool
}

const (
	dbConnStr = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
)

func OpenDb() (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf(dbConnStr, tok.Host, tok.Port, tok.User, tok.Password, tok.Dbname))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func AddUserBalance(telegramID int, amount float64) error {
	db, err := OpenDb()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Query("INSERT INTO users (telegram_id, balance) VALUES ($1, $2)", telegramID, amount)
	if err != nil {
		return err
	}

	return nil
}

func GetUserBalance(telegramID int) (float64, error) {
	var balance float64
	db, err := OpenDb()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	err = db.QueryRow("SELECT balance FROM users WHERE telegram_id = $1", telegramID).Scan(&balance)
	if err != nil {
		return 0, err
	}

	return balance, nil
}

func GetUser(telegramID int) (*User, error) {
	db, err := OpenDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row := db.QueryRow("SELECT * FROM users WHERE telegram_id = $1", telegramID)

	user := &User{}
	err = row.Scan(&user.ID, &user.TelegramID, &user.FirstName, &user.LastName, &user.UserName, &user.CreatedAt,
		&user.LastSeenAt, &user.Balance, &user.IsAuthorized)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func AddUser(user *User) error {
	
	db, err := OpenDb()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Query("INSERT INTO users (telegramid,id, firstname, lastname, username, createdat, lastseenat, balance, isauthorized) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		user.TelegramID, user.ID, user.FirstName, user.LastName, user.UserName, user.CreatedAt, user.LastSeenAt, user.Balance, user.IsAuthorized)
	if err != nil {
		return err
	}

	return nil
}