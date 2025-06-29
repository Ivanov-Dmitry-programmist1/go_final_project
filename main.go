package main

import (
	"fmt"
	"go_final_project/go_final_project/pkg/db"
	"log"
	"os"

	//"go1f/pkg/server"
)

func main() {
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = server.DefaultPort
	}

	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "scheduler.db"
	}

	err := db.Init(dbFile)
	if err != nil {
		log.Fatal("Ошибка инициализации базы данных:", err)
	}
	defer func() {
		if err := db.CloseDB(); err != nil {
			log.Println("Ошибка закрытия базы данных:", err)
		}
	}()

  
	fmt.Printf("Starting server on port %s\n", port)
	err = server.StartServer(port, server.WebDir)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
