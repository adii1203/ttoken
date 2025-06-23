package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	key_handler "github.com/adii1203/ttoken/internal/app/key/api"
	key_service "github.com/adii1203/ttoken/internal/app/key/service"
	project_handler "github.com/adii1203/ttoken/internal/app/project/api"
	project_service "github.com/adii1203/ttoken/internal/app/project/service"
	user_handler "github.com/adii1203/ttoken/internal/app/user/api"
	user_servise "github.com/adii1203/ttoken/internal/app/user/service"
	"github.com/adii1203/ttoken/internal/db"
	"github.com/adii1203/ttoken/internal/db/repository"
	"github.com/adii1203/ttoken/pkg/clerk"
	"github.com/adii1203/ttoken/pkg/validator"
	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("error loding .env")
	}
}

func main() {
	app := &http.Server{
		Addr:    os.Getenv("APP_ADDRESS"),
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

	clerk.InitClerk()

	db := db.InitDB()
	repo := repository.New(db)

	validator := validator.InitValidator()

	keyService := key_service.NewKeyService(repo)
	keyHandler := key_handler.NewKeyHandler(keyService, validator)

	projectService := project_service.NewProjectService(repo)
	projectHandler := project_handler.NewProjectHandler(projectService, validator)

	userService := user_servise.NewUserService(repo)
	userHandler := user_handler.NewUserHandler(userService)

	router := chi.NewRouter()

	r.Post("/webhook/clerk", userHandler.ClerkHandler)

	router.Group(func(r chi.Router) {
		r.Use(clerkhttp.RequireHeaderAuthorization())
		r.Post("/key.create", keyHandler.CreateKeyHandler)
	})

	router.Group(func(r chi.Router) {
		r.Post("/project.create", projectHandler.CreateProjectHandler)
	})

	r.Mount("/v1", router)
	return r
}
