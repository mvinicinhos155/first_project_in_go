package handlers

import (
	"encoding/json"
	"net/http"
	"todo-list/database"
	"todo-list/models"
	"strconv"
)


func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {


if r.Method == http.MethodOptions {
    return
}

	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
 	}

	err = database.InsertTodo(&todo)
	if  err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {

	todos, err := database.GetTodos()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(todos)
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	 if err != nil {
		http.Error(w, "Invalid id", 400)
		return
	 }

	err = database.DellTodo(id)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	
	json.NewEncoder(w).Encode("todo deleted successfully")
	w.WriteHeader(http.StatusNoContent)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {

	isSrt := r.URL.Query().Get("id")
	id, err := strconv.Atoi(isSrt)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var todo models.Todo

	err = json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = database.UpdateTodo(id, todo)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode("Todo updated successfully")
}