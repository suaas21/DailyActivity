package signatures

import (
	"context"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type DB struct {
	Ctx      context.Context
	MongoDB  *mongo.Database
	DBClient *mongo.Client
}

var Config string

func init() {
	flag.StringVar(&Config, "config", ".env", "config file path")
}

func InitDB() (*DB, error) {
	flag.Parse()
	err := godotenv.Load(Config)
	if err != nil {
		log.Println(".env not found")
		return nil, err
	}

	// Set client options and connect to MongoDB
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	database := &DB{}
	database.DBClient = client
	database.MongoDB = client.Database(os.Getenv("DB_NAME"))
	if (os.Getenv("ENABLE_E2E_TEST")) == "true" {
		database.MongoDB = client.Database(os.Getenv("TEST_DB_NAME"))
	}
	database.Ctx = context.Background()

	fmt.Println("Connected to MongoDB!")
	return database, nil
}
