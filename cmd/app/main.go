// main.go
package main

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"myservice/internal/database"
	"myservice/internal/repositories"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %v\n", err)
	}

	conn := database.DBconnect(ctx, os.Getenv("DATABASE_URL"))

	if err := repositories.Initusertable(ctx, *conn); err != nil {
		log.Fatal(err)
	}

}
