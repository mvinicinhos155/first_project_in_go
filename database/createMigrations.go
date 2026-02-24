package database

import (
	"log"
	"todo-list/models"
)

// Criar tabela (PostgreSQL)
func CreateTable() {

	query := `
		CREATE TABLE IF NOT EXISTS todos (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			done BOOLEAN DEFAULT false
		);
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Error while trying to create table:", err)
	}
}

func InsertTodo(todo *models.Todo) error {

	query := "INSERT INTO todos (title, done) VALUES ($1, $2)"

	_, err := DB.Exec(query, todo.Title, false)
	if err != nil {
		return err
	}

	return nil
}

func GetTodos() ([]models.Todo, error) {

	query := "SELECT id, title, done FROM todos"

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Done)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func DellTodo(id int) error {

	query := "DELETE FROM todos WHERE id = $1"

	_, err := DB.Exec(query, id)
	return err
}

func UpdateTodo(id int, todo models.Todo) error {

	query := "UPDATE todos SET title = $1, done = $2 WHERE id = $3"

	_, err := DB.Exec(query, todo.Title, todo.Done, id)
	return err
}