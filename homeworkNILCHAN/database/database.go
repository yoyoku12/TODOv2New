package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env не найден")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Panicf(" Ошибка подключения: %v", err)
	}

	fmt.Println("Успешное подключение к БД!")

	createTablTodoList := `
	CREATE TABLE IF NOT EXISTS todo_items (
	id SERIAL PRIMARY KEY,
	title TEXT NOT NULL,
	description TEXT,
	created_at TIMESTAMP NOT NULL DEFAULT now(),
	is_done BOOLEAN NOT NULL DEFAULT false,
	completed_at TIMESTAMP
);`

	createTableEvents := `
	CREATE TABLE IF NOT EXISTS events (
	id SERIAL PRIMARY KEY,
	event TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT now()
);`

	_, err = db.Exec(createTablTodoList)
	if err != nil {
		log.Panicf("Не удалось создать таблицу todo_items: %v", err)
	}

	_, err = db.Exec(createTableEvents)
	if err != nil {
		log.Panicf("Не удалось создать таблицу events: %v", err)
	}

	return db
}
