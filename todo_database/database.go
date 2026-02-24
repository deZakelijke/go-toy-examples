package tododatabase

import (
	"database/sql"
	"fmt"
	_ "github.com/glebarez/go-sqlite"
)

type TodoItem struct {
	Id          int64  `json:"id" form:"id"`
	Description string `json:"description" form:"description"`
	Done        bool   `json:"done" form:"done"`
}

var DB *sql.DB

func GetOrCreateDB(path string) error {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := createTables(db); err != nil {
		return err
	}
	DB = db
	return nil
}

func createTables(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS todo (
		id INTEGER PRIMARY KEY,
		description TEXT,
		done BOOL DEFAULT false
	);`

	_, err := db.Exec(query)

	if err != nil {
		return err
	}
	return nil
}

func Get(db *sql.DB) ([]TodoItem, error) {
	var items []TodoItem

	query := "SELECT id, description, done FROM todo;"

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Get: %v", err)
	}

	defer func() {

		if err := rows.Close(); err != nil {
			fmt.Println("Error when closin read rows: ", err)
		}
	}()

	for rows.Next() {
		var item TodoItem
		if err := rows.Scan(&item.Id, &item.Description, &item.Done); err != nil {
			return nil, fmt.Errorf("Get: %v", err)
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Get: %v", err)
	}

	return items, nil
}

func Insert(db *sql.DB, newItem TodoItem) (int64, error) {
	query := "INSERT INTO todo (description) VALUES (?);"

	result, err := db.Exec(query, newItem.Description)
	if err != nil {
		return 0, fmt.Errorf("Insert: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Insert: %v", err)
	}

	return id, nil
}

func Update(db *sql.DB, id int64, done bool) (int64, error) {
	query := "UPDATE todo SET done = ? WHERE id = ?;"

	result, err := db.Exec(query, done, id)
	if err != nil {
		return 0, fmt.Errorf("Update: %v", err)
	}
	return result.RowsAffected()
}
