package main

import (
	"context"
	"example/cloud-app/store/registry"
	"example/cloud-app/store/router"
	repository "example/cloud-app/store/usecase/repository/local"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
)

var PORT = "PORT"

func main() {
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		signal := <-c
		log.Printf("System call:%+v", signal)
		cancel()
	}()

	if err := serve(ctx); err != nil && err != http.ErrServerClosed {
		log.Printf("Failed to serve:%+v\n", err)
	}
}

func serve(ctx context.Context) error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %+v\n", err)
	}

	portEnv := os.Getenv(PORT)

	datastore := repository.NewKVStoreLocal()
	registry := registry.NewRegistry(datastore)
	router := router.Start(registry.NewAppController())

	srv := &http.Server{
		// Good practice to set timeouts to avoid Slowloris attacks.
		Addr:         ":" + portEnv,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	go func() {
		// Run our server in a goroutine so that it doesn't block.
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()

	log.Printf("server started")
	<-ctx.Done()
	log.Printf("server stopped")

	// Create a deadline to wait for current requests to complete
	ctxShutdown, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := srv.Shutdown(ctxShutdown); err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
	}
	log.Printf("server exited properly")

	return err
}
