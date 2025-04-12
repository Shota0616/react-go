package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s?parseTime=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE")))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	// データベース接続テスト
	http.HandleFunc("/api/db-test", func(w http.ResponseWriter, r *http.Request) {
		db, err := connectDB()
		if err != nil {
			http.Error(w, fmt.Sprintf("DB connection failed: %v", err), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			http.Error(w, fmt.Sprintf("DB ping failed: %v", err), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Successfully connected to MySQL database!")
	})

	// 既存のハンドラ
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Go backend!")
	})

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
