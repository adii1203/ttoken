package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.ngrok.com/ngrok/v2"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file", err)
	}
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

const address = "http://localhost:5000"

func run(ctx context.Context) error {
	agent, err := ngrok.NewAgent(ngrok.WithAuthtoken(os.Getenv("NGROK_AUTHTOKEN")))

	if err != nil {
		return err
	}

	ln, err := agent.Forward(ctx,
		ngrok.WithUpstream(address),
		ngrok.WithURL("darling-precious-sunbeam.ngrok-free.app"),
	)

	if err != nil {
		fmt.Println("Error", err)
		return err
	}

	fmt.Println("Endpoint online: forwarding from", ln.URL(), "to", address)

	<-ln.Done()

	return nil
}
