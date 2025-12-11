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

	// Repositories
	clubRepo := repository.NewClubRepo(dbPool)
	teamRepo := repository.NewTeamRepo(dbPool)
	coachRepo := repository.NewCoachRepo(dbPool)
	playerRepo := repository.NewPlayerRepo(dbPool)
	stadiumRepo := repository.NewStadiumRepo(dbPool)
	staffRepo := repository.NewStaffRepo(dbPool)
	gameRepo := repository.NewGameRepo(dbPool)

	// Usecase
	clubUC := usecase.NewClubUsecase(clubRepo)
	teamUC := usecase.NewTeamUsecase(teamRepo)
	coachUC := usecase.NewCoachUsecase(coachRepo)
	playerUC := usecase.NewPlayerUsecase(playerRepo)
	stadiumUC := usecase.NewStadiumUsecase(stadiumRepo)
	staffUC := usecase.NewStaffUsecase(staffRepo)
	gameUC := usecase.NewGameUsecase(gameRepo)

	// Handler
	clubHandler := delivery.NewClubHandler(clubUC)
	teamHandler := delivery.NewTeamHandler(teamUC)
	coachHandler := delivery.NewCoachHandler(coachUC)
	playerHandler := delivery.NewPlayerHandler(playerUC)
	stadiumHandler := delivery.NewStadiumHandler(stadiumUC)
	staffHandler := delivery.NewStaffHandler(staffUC)
	gameHandler := delivery.NewGameHandler(gameUC)

	// Router
	router := delivery.NewRouter(
		clubHandler,
		teamHandler,
		coachHandler,
		playerHandler,
		stadiumHandler,
		staffHandler,
		gameHandler,
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
