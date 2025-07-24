package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var db *sql.DB

func main() {
	var err error
	connStr := "user=postgres password=postgres dbname=todo_db sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Databasega ulanishda xatolik:", err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/todos", getTodos).Methods("GET")
	r.HandleFunc("/todos", createTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", updateTodo).Methods("PATCH")
	r.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")

	fmt.Println("Server http://localhost:8080 da ishga tushdi")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, done FROM todos")
	if err != nil {
		http.Error(w, "Ma'lumotlar olinmadi", 500)
		return
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Done); err != nil {
			http.Error(w, "Ma'lumotni o'qishda xatolik", 500)
			return
		}
		todos = append(todos, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var t Todo
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Yaroqsiz JSON", 400)
		return
	}

	err := db.QueryRow("INSERT INTO todos (title, done) VALUES ($1, $2) RETURNING id", t.Title, t.Done).Scan(&t.ID)
	if err != nil {
		http.Error(w, "Qo‘shishda xatolik", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var t Todo
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Yaroqsiz JSON", 400)
		return
	}

	result, err := db.Exec("UPDATE todos SET title=$1, done=$2 WHERE id=$3", t.Title, t.Done, id)
	if err != nil {
		http.Error(w, "Yangilashda xatolik", 500)
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		http.Error(w, "Bunday todo topilmadi", 404)
		return
	}

	fmt.Fprint(w, "Todo yangilandi ✅")
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	result, err := db.Exec("DELETE FROM todos WHERE id=$1", id)
	if err != nil {
		http.Error(w, "O‘chirishda xatolik", 500)
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		http.Error(w, "Bunday todo topilmadi", 404)
		return
	}

	fmt.Fprint(w, "Todo o‘chirildi 🗑️")
}
