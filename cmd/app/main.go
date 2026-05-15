// main.go
package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"myservice/internal/database"
	"myservice/internal/repositories"
	"myservice/internal/services"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file %v\n", err)
	}

	conn, err := database.DBconnect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	userRepo, _ := repositories.NewDatabaseConn(*conn)

	if err := userRepo.InitUserTable(ctx); err != nil {
		log.Fatal(err)
	}

	userService, _ := services.NewUserService(userRepo)

	userService.CreatetUser(ctx, "SDAASD")
	fmt.Println(userService.GetUser(ctx, 1))
	fmt.Println(userService.GetUser(ctx, 2))
	fmt.Println(userService.GetUser(ctx, 3))
	//if err := repositories.InsertUser(ctx, *conn, "фафыа"); err != nil {
	//	fmt.Print(fmt.Errorf("%w", err))
	//}

	//if err := repositories.UpdateUser(ctx, *conn, 2, "aaaa"); err != nil {
	//	fmt.Print(fmt.Errorf("%w", err))
	//}

	//repositories.DeleteUser(ctx, *conn, 1)

	//fmt.Print(repositories.GetUser(ctx, *conn, 2))
}
