package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"todo-list/database"
	"todo-list/handlers"
)

func main() {

	database.Connect()
	database.Migrations()

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!!"))
	})

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		
		switch r.Method {
			case http.MethodGet:
				handlers.GetTodosHandler(w, r)
			
			case http.MethodPost:
				handlers.CreateTodoHandler(w, r)

			case http.MethodDelete: 
				handlers.DeleteTodoHandler(w, r)
			
			case http.MethodPut: 
				handlers.UpdateTodoHandler(w, r)
			default: 
				w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Servvidor rodando na porta 🚀", port)
	log.Fatal(http.ListenAndServe(":"+port, enableCORS(http.DefaultServeMux)))


}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}