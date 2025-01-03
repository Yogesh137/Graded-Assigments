package main

import (
	db "crud/config"
	"crud/controller"
	"crud/repository"
	"crud/service"
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	_ "modernc.org/sqlite"
)

func InitializeAuthDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./users.db")
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to the database successfully.")

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL,
	password TEXT NOT NULL
	);`
	if _, err := db.Exec(createTableSQL); err != nil {
		return nil, fmt.Errorf("failed to create table: %v", err)
	}

	insertSQL := `
INSERT INTO users (username, password)
VALUES (?,?);`
	if _, err := db.Exec(insertSQL, "admin", "admin"); err != nil {
		return nil, err
	}

	return db, nil
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("Started %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		fmt.Printf("Completed %s in %v\n", r.URL.Path, time.Since(start))
	})
}

func Authmiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Basic ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		payload, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(authHeader, "Basic "))
		if err != nil {
			http.Error(w, "Invalid Authorization Header", http.StatusUnauthorized)
			return
		}
		credentials := strings.SplitN(string(payload), ":", 2)
		if len(credentials) != 2 {
			http.Error(w, "Invalid credentials format", http.StatusUnauthorized)
			return
		}

		username := credentials[0]
		password := credentials[1]
		db, err := sql.Open("sqlite", "users.db")
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		var dbPassword string
		query := "SELECT password FROM users WHERE username = ?"
		err = db.QueryRow(query, username).Scan(&dbPassword)
		if err != nil || dbPassword != password {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Initialize the database
	auth_db, err := InitializeAuthDatabase()
	if err != nil {
		fmt.Println("Database initialization failed:", err)
		return
	}
	defer auth_db.Close()

	db.InitializeDatabase()

	blogRepo := repository.NewBlogRepository(db.GetDB())
	blogService := service.NewBlogService(blogRepo)
	blogController := controller.NewBlogController(blogService)

	//Handlers
	homeHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Home page!")
	})
	http.Handle("/", Authmiddleware(LoggingMiddleware(homeHandler)))

	//Routes
	blogsHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			blogController.GetAllBlogs(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	http.Handle("/blogs", Authmiddleware(LoggingMiddleware(blogsHandler)))

	blogHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			fmt.Println("Create Blog")
			blogController.CreateBlog(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	http.Handle("/blog", Authmiddleware(LoggingMiddleware(blogHandler)))

	blogHandler2 := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fmt.Println("Get Blog")
			blogController.GetBlog(w, r)
		case http.MethodPut:
			fmt.Println("Update Blog")
			blogController.UpdateBlog(w, r)
		case http.MethodDelete:
			fmt.Println("Delete Blog")
			blogController.DeleteBlog(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	http.Handle("/blog/", Authmiddleware(LoggingMiddleware(blogHandler2)))

	// Start server
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error Starting Server:", err)
	}

}
