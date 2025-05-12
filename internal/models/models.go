package models

type User struct {
	ID           int    `db:"id"`
	TelegramID   string `db:"telegram_id"` // Можно не использовать, но оставим как уникальный логин
	Username     string `db:"username"`
	PasswordHash string `db:"password_hash"`
}

type Category struct {
	ID     int    `db:"id"`
	UserID int    `db:"user_id"`
	Name   string `db:"name"`
}

type Card struct {
	ID          int    `db:"id"`
	UserID      int    `db:"user_id"`
	CategoryID  int    `db:"category_id"`
	Word        string `db:"word"`
	Translation string `db:"translation"`
}

type Attempt struct {
	ID      int  `db:"id"`
	UserID  int  `db:"user_id"`
	CardID  int  `db:"card_id"`
	Correct bool `db:"correct"`
}
