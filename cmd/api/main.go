package main

import (
	"Auth-Server/internal/data/database"
	"Auth-Server/internal/server"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	DbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	defer DbConn.Close()

	db := database.New(DbConn)

	err = server.NewServer(db).ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
