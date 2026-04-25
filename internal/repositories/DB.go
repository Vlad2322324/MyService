package repositories

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func DBconnect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %v\n", err)
	}
	fmt.Println(os.Getenv("DATABASE_URL"))

	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatalf("Error connecting to database%v\n", err)
	}
	defer db.Close(context.Background())

	_, err = db.Exec(context.Background(),
		"CREATE TABLE IF NOT EXISTS stub (id SERIAL PRIMARY KEY)")
	if err != nil {
		log.Fatalf("Cannot create table%v\n", err)
	}
}
