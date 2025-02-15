package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/sam-berry/ecfr-analyzer/server/api"
	"github.com/sam-berry/ecfr-analyzer/server/config"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	masterCtx, masterCancel := context.WithCancel(context.Background())
	defer masterCancel()

	var sigs = make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	db := config.ConnectToDatabase("ecfr-service")
	defer db.Close()
	config.ConfigureDB(db)

	app := config.InitHTTPApp()

	router := app.Group("/ecfr-service")

	// Unauthenticated APIs
	registerAPIs(
		[]api.API{
			&api.ECFRImport{
				Router: router,
			},
		},
	)

	router.Use(config.TokenHandler)

	// Authenticated APIs
	registerAPIs(
		[]api.API{},
	)

	go func() {
		startApp(app)
	}()

	go func() {
		sig := <-sigs
		log.Printf("Received signal %s, initiating graceful shutdown...", sig)
		masterCancel()
	}()

	<-masterCtx.Done()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := app.ShutdownWithContext(shutdownCtx); err != nil {
		log.Printf("HTTP server Shutdown: %v", err)
	}

	if err := db.Close(); err != nil {
		log.Printf("Error closing database: %v", err)
	}

	log.Println("Graceful shutdown complete.")
}

func startApp(router *fiber.App) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}
	log.Println(fmt.Sprintf("Starting app on :%v", port))
	err := router.Listen(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}

func registerAPIs(apis []api.API) {
	for _, a := range apis {
		a.Register()
	}
}
