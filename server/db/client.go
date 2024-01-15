package db

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBCollections struct {
	User      string
	Form      string
	FormaData string
}

var Collections = DBCollections{
	User:      "User",
	Form:      "Form",
	FormaData: "FormData",
}

func InitDb() *mongo.Database {
	// load dotenv
	if err := godotenv.Load(); err != nil {
		panic("Error: Failed to load dotenv")
	}

	db_url := os.Getenv("MONGODB_URL")
	db_name := os.Getenv("MONGODB_NAME")

	if db_url == "" {
		panic("DB url cannot be empty")
	}

	if db_name == "" {
		panic("DB name cannot be empty")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db_url))

	if err != nil {
		panic(err)
	}

	// handle disconnect
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	return client.Database(db_name)
}
