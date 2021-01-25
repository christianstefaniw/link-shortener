package storage

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var Collection *mongo.Collection

func LoadStorage() *mongo.Collection{
	clientOptions := options.Client().ApplyURI(os.Getenv("CONNECTION_STRING"))

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil{
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil{
		log.Fatal(err)
	}

	return client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("COLLNAME"))
}

func LoadEnv(){
	err := godotenv.Load(os.ExpandEnv("C:/Users/cpste/Desktop/Projects/url-shortener/server/.env"))
	if err != nil{
		log.Fatal("Error loading .env file")
	}
}