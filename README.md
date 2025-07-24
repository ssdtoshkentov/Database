# 📦 Todo API with Go + PostgreSQL

Welcome! This project is a simple yet powerful RESTful API for managing todo items.  
Built with **Go** and **PostgreSQL**, it supports full CRUD functionality and is ready to be extended.

---

## 🚀 Features

- ⚡ Built using Go (Golang)
- 🐘 PostgreSQL database integration
- 🔁 Full CRUD operations (Create, Read, Update, Delete)
- 🛠 Modular and clean code structure
- 🌐 RESTful API using Gorilla Mux

---

## 📂 Project Structure

```
Database/
├── main.go       // Entry point of the application
├── db.go         // Database connection logic
├── model.go      // Data model (Todo struct)
├── go.mod        // Go module definitions
├── go.sum        // Dependency checksums
└── README.md     // Project documentation
```

---

## 📌 Endpoints

| Method | Endpoint        | Description             |
|--------|------------------|-------------------------|
| GET    | `/todos`         | Get all todo items      |
| POST   | `/todos`         | Create a new todo       |
| PATCH  | `/todos/{id}`    | Update an existing todo |
| DELETE | `/todos/{id}`    | Delete a todo           |

---

## ⚙️ Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/ssdtoshkentov/Database.git
cd Database
```

### 2. Setup PostgreSQL

Make sure your PostgreSQL server is running and the database `todo_db` exists.  
You can create it like this:

```sql
CREATE DATABASE todo_db;

CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    done BOOLEAN DEFAULT FALSE
);
```

Update the connection string in `main.go` if needed.

---

### 3. Run the server

```bash
go run main.go
```

Server will run at: [http://localhost:8080](http://localhost:8080)

---

## 📄 Example JSON Payloads

### ✅ POST /todos

```json
{
  "title": "Finish Go backend project",
  "done": false
}
```

### ✅ PATCH /todos/{id}

```json
{
  "title": "Update Go backend project",
  "done": true
}
```

---

## 🧠 Why this project?

This project was created to:

- Learn how to build REST APIs with Go
- Practice PostgreSQL integration
- Understand router handling with Gorilla Mux
- Improve clean and modular code organization

---

## 🧑‍💻 Author

**Abduvali Toshkentov**  
🔗 GitHub: [@ssdtoshkentov](https://github.com/ssdtoshkentov)

---

## ✅ Future Improvements

- ✅ Modular file structure
- ⏳ Add unit tests
- ⏳ Add JWT authentication
- ⏳ Use .env file for DB credentials
- ⏳ Add Docker support
- ⏳ Deploy to Railway or Render

---

## 📬 Feedback

Feel free to open issues or pull requests.  
Thanks for checking out the project! 🙌
