package logging

import (
	"database/sql"
	"fmt"
	"time"
)

func AddEventToDb(db *sql.DB, event string) {
	query := `
		INSERT INTO events (event, created_at)
		VALUES ($1, $2)
	`
	_, err := db.Exec(query, event, time.Now()) // передаём оба параметра!
	if err != nil {
		fmt.Println("Ошибка добавления события", err)
	}
	fmt.Println("Событие залогировано")
}
