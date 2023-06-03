package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetConnect(uri string) (*mongo.Client, error) {
	// setting up context timeout for the database connection
	dbCtx, dbCancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer dbCancelCtx()

	// connecting to the database using the URI string
	client, err := mongo.Connect(dbCtx, options.Client().ApplyURI(uri).SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1)))
	if err != nil {
		log.Panicln("Error while connecting to database: ", err)
	}

	// pinging the database
	if err := client.Ping(dbCtx, nil); err != nil {
		log.Fatalln("cannot ping the database: ", err)
	}

	return client, nil
}

// OpenConnect : This help to keep the database connection open if
// issue occur while trying to connect
func OpenConnect() *mongo.Client {
	count := 0
	for {
		
		client, err := SetConnect(os.Getenv("URI"))
		if err == nil {
			log.Println("Mail App Database is Connected")
			return client
		}
		if count >= 10 {
			log.Println("Mail App Database not Connected")
			return nil
		}

		log.Println("Mail App Database Trying to Connect")
		count += 1
		continue
	}
}
