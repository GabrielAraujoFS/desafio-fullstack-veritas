package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

var tasks []Task
var CurrentID int

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if tasks == nil {
		tasks = []Task{}
	}

	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	json.NewDecoder(r.Body).Decode(&newTask)
	CurrentID++
	newTask.ID = CurrentID
	newTask.Status = "todo"
	tasks = append(tasks, newTask)
	json.NewEncoder(w).Encode(newTask)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for i, task := range tasks {
		if task.ID == id {
			json.NewDecoder(r.Body).Decode(&tasks[i])
			tasks[i].ID = id
			json.NewEncoder(w).Encode(tasks[i])
			return
		}
	}
	http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Tarefa deletada com sucesso",
			})
			return
		}
	}
	http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
}
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

	// Permitir comunicação com o frontend React
	corsObj := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", handlers.CORS(corsObj, headers, methods)(router))
}
