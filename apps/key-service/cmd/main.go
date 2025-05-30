package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adii1203/ttoken/apps/key-service/internal/api"
	"github.com/adii1203/ttoken/apps/key-service/internal/db"
	"github.com/adii1203/ttoken/apps/key-service/internal/db/repository"
	"github.com/adii1203/ttoken/apps/key-service/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	app := &http.Server{
		Addr:    "0.0.0.0:5000",
		Handler: server(),
	}

	appCtx, appStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig
		log.Println("shutdown signal received...")

		shutdownCtx, shutdownCancel := context.WithTimeout(appCtx, 30*time.Second)
		defer shutdownCancel()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out... forcing exit.")
			}
		}()

		err := app.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		appStopCtx()
	}()

	log.Println("server starting on", app.Addr)

	err := app.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-appCtx.Done()
	log.Println("server shutdown complete")
}

func server() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	db := db.InitDB()
	repo := repository.New(db)

	service := service.NewKeyService(repo)
	keyHandler := api.NewKeyHandler(service)

	// validator := validator.InitValidator()

	keyRouter := chi.NewRouter()
	keyRouter.Group(func(r chi.Router) {
		keyRouter.Post("/key.create", keyHandler.CreateKeyHandler)
	})

	r.Mount("/v1", keyRouter)
	return r
}
