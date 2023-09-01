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

// func MongoDatabase() string {
// 	if os.Getenv("MONGOURI") == "" {
// 		fmt.Printf("MONGO_DB env is not should be empty")
// 		os.Exit(-1)
// 	}

// return os.Getenv("MONGOURI")
// }

func KafkaHost() []string {

	if os.Getenv("KAFKA_BROKERCONNECT") == "" {
		fmt.Printf("KAFKA_BROKERCONNECT env is not should be empty")
		os.Exit(-1)
	}

	return []string{os.Getenv("KAFKA_BROKERCONNECT")}
}

func WebhookName() string {

	if os.Getenv("WEBHOOK_NAME") == "" {
		fmt.Printf("WEBHOOK_NAME env is not should be empty")
		os.Exit(-1)
	}

	return os.Getenv("WEBHOOK_NAME")

}
