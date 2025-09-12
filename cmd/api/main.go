package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pedroandriotti/homelink-ddd/internal/bootstrap"
	"github.com/pedroandriotti/homelink-ddd/internal/container"
	"github.com/pedroandriotti/homelink-ddd/internal/database"
	"github.com/pedroandriotti/homelink-ddd/internal/messaging"
	"gorm.io/gorm"
)

func main() {
	app, err := Bootstrap()

	if err != nil {
		log.Fatal("Failed to bootstrap application", err)
	}

	port := getEnv("PORT", "8080")
	log.Printf("Starting server on port %s", port)

	go startWebServer(app, port)

	waitForShutdown(app)
}

func setupDatabase() *gorm.DB {
	log.Println("Setting up database...")
	config := database.NewConfigFromEnv()

	db, err := database.NewConnection(config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database setup completed")
	return db
}

func Bootstrap() (*bootstrap.App, error) {
	db := setupDatabase()

	publisher, err := messaging.NewEventPublisher()
	if err != nil {
		return nil, err
	}

	cont := container.NewContainer(db, publisher)

	app := bootstrap.NewWebApp(cont)

	if err := bootstrap.WireEvents(cont, publisher); err != nil {
		return nil, err
	}

	return app, nil
}

func startWebServer(app *bootstrap.App, port string) {
	log.Printf("Starting web server on port %s", port)

	if err := app.Start(port); err != nil {
		log.Printf("Web server stopped: %v", err)
	}
}

func waitForShutdown(app *bootstrap.App) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := app.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
