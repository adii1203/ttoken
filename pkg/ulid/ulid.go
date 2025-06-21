package ulid

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func GenerateULID() ulid.ULID {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))

	return ulid.MustNew(ulid.Timestamp(time.Now()), entropy)
}
