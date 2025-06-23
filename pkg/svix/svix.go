package svix

import (
	"log"

	svix "github.com/svix/svix-webhooks/go"
)

func InitSvix() *svix.Webhook {
	secret := ""

	wh, err := svix.NewWebhook(secret)
	if err != nil {
		log.Fatal("Error initlisizing svix", err)
	}

	return wh
}
