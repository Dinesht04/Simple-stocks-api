package middleware

import (
	"log"
)

func Check(err error) {
	if err != nil {
		log.Fatal("error %w", err)
	}
}
