package config

import (
	"fmt"
	"os"
)

func MongoUri() string {
	if os.Getenv("MONGOURI") == "" {
		fmt.Printf("MONGOURI env is not should be empty")
		os.Exit(-1)
	}

	return os.Getenv("MONGOURI")
}
