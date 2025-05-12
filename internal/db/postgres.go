package db

import (
	"fmt"
	"io/ioutil"
	"wordtrainer/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func Connect(cfg *config.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	var err error
	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Выполняем миграции
	if err := runMigrations(); err != nil {
		return nil, fmt.Errorf("ошибка при выполнении миграций: %v", err)
	}

	return db, nil
}

func GetDB() *sqlx.DB {
	return db
}

func runMigrations() error {
	// Читаем файл с миграциями
	migrations, err := ioutil.ReadFile("internal/migrations/001_init.up.sql")
	if err != nil {
		return fmt.Errorf("ошибка при чтении файла миграций: %v", err)
	}

	// Выполняем миграции
	_, err = db.Exec(string(migrations))
	if err != nil {
		return fmt.Errorf("ошибка при выполнении миграций: %v", err)
	}

	return nil
}
