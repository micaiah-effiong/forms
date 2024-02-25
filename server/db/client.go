package db

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBCollections struct {
	User     string
	Form     string
	FormData string
}

type FormData struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	FormId    primitive.ObjectID `bson:"formId" json:"form_id"`
	Data      primitive.M        `bson:"data" json:"data"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

var Collections = DBCollections{
	User:     "User",
	Form:     "Form",
	FormData: "FormData",
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
