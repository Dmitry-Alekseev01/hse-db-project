package main

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"hse-football/config"
	_ "hse-football/docs"
	"hse-football/internal/delivery"
	"hse-football/internal/repository"
	"hse-football/internal/usecase"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title           Football Service API
// @version         1.0
// @description     REST API сервиса управления футбольным клубом.
// @termsOfService  http://swagger.io/terms/
// @contact.name   	API Support
// @contact.email  	support@example.com
// @license.name  	MIT License
// @license.url   	https://opensource.org/licenses/MIT
// @host      		localhost:8080
// @BasePath 		/api/v1
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.Load()
	dsn := cfg.BuildDSN()

	// Graceful shutdown handling
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		cancel()
	}()

	// Database setup
	dbPool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	defer dbPool.Close()

	if err := dbPool.Ping(ctx); err != nil {
		log.Fatal("DB ping failed:", err)
	}

	// Wire dependencies
	clubRepo := repository.NewClubRepo(dbPool)
	coachRepo := repository.NewCoachRepo(dbPool)
	clubUC := usecase.NewClubUsecase(clubRepo)
	coachUC := usecase.NewCoachUsecase(coachRepo)
	clubHandler := delivery.NewClubHandler(clubUC)
	coachHandler := delivery.NewCoachHandler(coachUC)

	router := delivery.NewRouter(
		clubHandler,
		coachHandler,
	)

	// HTTP Server with graceful shutdown
	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	go func() {
		log.Printf("Starting HTTP server on port %s", cfg.Port)
		log.Printf("Swagger available at http://localhost:%s/swagger/index.html", cfg.Port)

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server failed: %v", err)
		}
	}()

	<-ctx.Done()

	// Graceful shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	}

	log.Println("Service stopped gracefully")
}
