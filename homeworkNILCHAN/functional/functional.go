package functional

import (
	"database/sql"
	"fmt"
	"time"
)

type TodoItem struct {
	ID          int
	Title       string
	Description string
	CreatedAt   time.Time
	IsDone      bool
	CompletedAt *time.Time
}

func NewItem(db *sql.DB, title, description string) error {
	query := `
		INSERT INTO todo_items (title, description, created_at, is_done)
		VALUES ($1, $2, $3, $4)
	`

	_, err := db.Exec(query, title, description, time.Now(), false)
	if err != nil {
		return fmt.Errorf("ошибка добавления задачи: %w", err)
	}

	fmt.Println("Задача успешно добавлена в список!")
	return nil
}

func ShowFullList(db *sql.DB) {
	query := `
		SELECT id, title, description, created_at, is_done, completed_at
		FROM todo_items
		ORDER BY created_at DESC;
	`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	defer rows.Close()

	fmt.Println("Список всех заданий:")
	for rows.Next() {
		var id int
		var title, description string
		var createdAt time.Time
		var isDone bool
		var completedAt *time.Time

		rows.Scan(&id, &title, &description, &createdAt, &isDone, &completedAt)

		status := "Не выполнено"
		if isDone {
			status = "Выполнено"
		}

		createdStr := createdAt.Format("02.01.2006 15:04:05") // читаемый формат
		var completedStr string
		if completedAt != nil {
			completedStr = completedAt.Format("02.01.2006 15:04:05")
		} else {
			completedStr = "—"
		}

		result := fmt.Sprintf(
			"ID: %d\nНазвание: %s\nОписание: %s\nСоздано: %s\nСтатус: %s\nВыполнено: %s\n",
			id, title, description, createdStr, status, completedStr)

		fmt.Println(result)
	}
}

func ShowFullListEvents(db *sql.DB) {
	query := `
		SELECT id, event, created_at
		FROM events
		ORDER BY created_at DESC;
	`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	defer rows.Close()

	fmt.Println("Список всех событий:")
	for rows.Next() {
		var Id int
		var Event string
		var Created_at time.Time

		rows.Scan(&Id, &Event, &Created_at)

		result := fmt.Sprintf(
			"ID: %d\nСобытие: %s\nСоздано: %s",
			Id, Event, Created_at.Format("02.01.2006 15:04:05"))

		fmt.Println(result)
	}
}

func DeleteItem(db *sql.DB, title string) string {
	query := `
		DELETE FROM todo_items
		WHERE title = $1;
	`
	_, err := db.Exec(query, title)
	if err != nil {
		return "Ошбика удаления записи из БД"
	}
	result := fmt.Sprintf("Задание %s успешно удалено", title)
	return result
}

func Help() {
	fmt.Println("\nДоступные команды:")
	fmt.Println("  help                    — Показать это сообщение")
	fmt.Println("  add                     — Добавить новую задачу")
	fmt.Println("  list                    — Показать все задачи")
	fmt.Println("  del {название}          — Удалить задачу по названию")
	fmt.Println("  done {название}         — Пометить задачу как выполненную")
	fmt.Println("  events                  — Показать список событий")
	fmt.Println("  exit                    — Выйти из программы")
}

func MarkAsDoneByTitle(db *sql.DB, title string) error {
	query := `
		UPDATE todo_items
		SET is_done = true,
		completed_at = $1
		WHERE title = $2;
	`

	_, err := db.Exec(query, time.Now(), title)
	if err != nil {
		return fmt.Errorf("не удалось обновить задачу: %w", err)
	}

	fmt.Printf("Задача '%s' помечена как выполненная!\n", title)
	return nil
}
