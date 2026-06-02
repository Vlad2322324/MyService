// main.go
package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	api "myservice/internal/api/generated"
	"myservice/internal/database"
	"myservice/internal/handlers/user_handler"
	"myservice/internal/repositories"
	"myservice/internal/services"
	"os"
)

func main() {

	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//
	//if err := godotenv.Load(); err != nil {
	//	log.Fatalf("Error loading .env file %v\n", err)
	//}
	//
	//conn, err := database.DBconnect(ctx, os.Getenv("DATABASE_URL"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//userRepo, _ := repositories.NewDatabaseConn(*conn)
	//
	//if err := userRepo.InitUserTable(ctx); err != nil {
	//	log.Fatal(err)
	//}
	//
	//userService, _ := services.NewUserService(userRepo)
	//
	//userService.CreateUser(ctx, "SDAASD")
	//
	//fmt.Println(userService.GetUser(ctx, 1))
	//fmt.Println(userService.GetUser(ctx, 2))
	//fmt.Println(userService.GetUser(ctx, 3))
	//if err := repositories.InsertUser(ctx, *conn, "фафыа"); err != nil {
	//	fmt.Print(fmt.Errorf("%w", err))
	//}

	//if err := repositories.UpdateUser(ctx, *conn, 2, "aaaa"); err != nil {
	//	fmt.Print(fmt.Errorf("%w", err))
	//}

	//repositories.DeleteUser(ctx, *conn, 1)

	//fmt.Print(repositories.GetUser(ctx, *conn, 2))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file %v\n", err)
	}

	conn, err := database.DBconnect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	repo, _ := repositories.NewDatabaseConn(*conn)

	service, err := services.NewUserService(repo)
	if err != nil {
		log.Fatal(err)
	}

	handler := user_handler.NewUserHandler(service)

	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Логируем метод и путь запроса
			method := c.Request().Method
			path := c.Request().URL.Path
			println("Новый запрос:", method, path)
			return next(c) // Передаем управление следующему обработчику
		}
	})

	api.RegisterHandlers(
		e,
		api.NewStrictHandler(handler, nil),
	)

	e.Logger.Fatal(e.Start(":8080"))

}
