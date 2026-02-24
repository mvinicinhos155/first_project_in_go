package database

import (
	"log"
	"todo-list/models"
)

//função que irá criar o banco de dados
func CreateDatabase() {

	query := "CREATE DATABASE IF NOT EXISTS first_project"

	_, err := DB.Exec(query)

	if err != nil {
		log.Fatal("Error while trying to create database.", err)
	}
}

//função que irá criar as tabelas
func CreateTable() {

	query := (`
		CREATE TABLE IF NOT EXISTS first_project (
		 id INT AUTO_INCREMENT PRIMARY KEY,
		 title VARCHAR(255),
		 done TINYINT(1) DEFAULT 0
		 )
		 `)


		 _, err := DB.Exec(query)


		  if err != nil {
			log.Fatal("Error while trying to create table.", err)
		  }
}


func InsertTodo(todo *models.Todo) error {

	query := "INSERT INTO first_project (title, done) VALUES (?, ?)"

	_, err := DB.Exec(query, todo.Title, false)
	
	if err != nil {
		log.Fatal("Error while trying to insert todo.", err)
	}

	return err
	
} 


func GetTodos() ([]models.Todo, error) {

	query := "SELECT id, title, done FROM first_project"

	rows, err := DB.Query(query)

	if err != nil {
		log.Fatal("Error while trying to get todos")
	}

	defer rows.Close()

	var todos []models.Todo


	for rows.Next() {
		var todo models.Todo
		rows.Scan(&todo.ID, &todo.Title,&todo.Done)
		todos = append(todos, todo)
	}

	return  todos, nil
}

func DellTodo(id int) error {

	query := "DELETE FROM first_project WHERE id = ?"

	_, err := DB.Exec(query, id)

	if err != nil {
		log.Fatal("Error while tryning to delete todo", err)
	}

	return err
} 

func UpdateTodo(id int, todo models.Todo) error {

	query := "UPDATE first_project SET title = ?, done = ? WHERE id = ?"

	_, err := DB.Exec(query, todo.Title, todo.Done, id)

	if err != nil {
		log.Fatal("Error while tryning to update todo", err)
		
	}

	return err

}