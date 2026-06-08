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
	"myservice/internal/middleware"
	"myservice/internal/repositories"
	"myservice/internal/services"
	"myservice/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	// Сначала загружаем .env чтобы переменные окружения были доступны
	if err := godotenv.Load(); err != nil {
		// не фатальная ошибка — .env может отсутствовать в production
		log.Printf("Warning: could not load .env file: %v", err)
	}

	// Инициализация логгера (читает ENV из окружения)
	if err := logger.Init(); err != nil {
		log.Fatalf("failed to init logger: %v", err)
	}
	defer logger.Sync()

	logger.Log.Infof("logger initialized (ENV=%s)", os.Getenv("ENV"))

	conn, err := database.DBconnect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Log.Fatalf("database connect error: %v", err)
	}

	repo, _ := repositories.NewDatabaseConn(conn)

	// Инициализация таблицы users если нужно
	if err := repo.InitUserTable(ctx); err != nil {
		logger.Log.Fatalf("failed to init users table: %v", err)
	}

	service, err := services.NewUserService(repo)
	if err != nil {
		logger.Log.Fatalf("service init error: %v", err)
	}

	handler := user_handler.NewUserHandler(service)

	e := echo.New()
	// Request ID middleware (X-Request-ID)
	e.Use(middleware.RequestID)
	// Логируем метод и путь запроса (может использовать request_id из контекста)
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			method := c.Request().Method
			path := c.Request().URL.Path
			rid, _ := c.Get("request_id").(string)
			if rid != "" {
				logger.Log.Infof("Новый запрос: %s %s request_id=%s", method, path, rid)
			} else {
				logger.Log.Infof("Новый запрос: %s %s", method, path)
			}
			return next(c)
		}
	})

	api.RegisterHandlers(
		e,
		api.NewStrictHandler(handler, nil),
	)

	// Start server in goroutine
	go func() {
		if err := e.Start(":8081"); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatalf("server start error: %v", err)
		}
	}()

	// Use NotifyContext to listen for interrupt/terminate signals and support graceful shutdown
	sigCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Block until a signal is received
	<-sigCtx.Done()
	logger.Log.Infof("shutdown signal received")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := e.Shutdown(shutdownCtx); err != nil {
		logger.Log.Errorf("error during server shutdown: %v", err)
	}

	// Close DB connection
	if conn != nil {
		if err := conn.Close(shutdownCtx); err != nil {
			logger.Log.Warnf("error closing db connection: %v", err)
		}
	}

}
