// main.go
package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"myservice/internal/database"
	"myservice/internal/repositories"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file %v\n", err)
	}

	conn := database.DBconnect(ctx, os.Getenv("DATABASE_URL"))

	if err := repositories.Initusertable(ctx, *conn); err != nil {
		log.Fatal(err)
	}

	//if err := repositories.Inserteuser(ctx, *conn, "фафыа"); err != nil {
	//	fmt.Print(fmt.Errorf("%w", err))
	//}

	//if err := repositories.Updateuser(ctx, *conn, 2, "aaaa"); err != nil {
	//	fmt.Print(fmt.Errorf("%w", err))
	//}

	//repositories.Deleteuser(ctx, *conn, 1)

	fmt.Print(repositories.GetUser(ctx, *conn, 2))
}
