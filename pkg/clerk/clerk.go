package clerk

import (
	"log"

	"github.com/clerk/clerk-sdk-go/v2"
)

func InitClerk() {
	// key := os.Getenv("CLERK_API_KEY")
	key := ""
	if key == "" {
		log.Fatal("Invalid clerk key")
	}

	clerk.SetKey(key)
}
