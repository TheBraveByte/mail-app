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
	dbCtx, dbCancelCtx := context.WithTimeout(context.Background(), 100*time.Second)
	defer dbCancelCtx()

	// connecting to the database using the URI string
	client, err := mongo.Connect(dbCtx, options.Client().ApplyURI(uri))
	//.SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1)))
	if err != nil {
		log.Panicln("Error while connecting to database: ", err)
	}

	// pinging the database
	if err := client.Ping(dbCtx, nil); err != nil {
		log.Fatalln("cannot ping the database: ", err)
	}

	return client, nil
}

// OpenConnect : This help to keep the database connection open if issue occur while trying to connect
func OpenConnect() *mongo.Client {
	uri := os.Getenv("URI")
	count := 0
	log.Println("....... Setting up Connection to MongoDB .......")
	for {

		client, err := SetConnect(uri)
		if err != nil {
			log.Println("Mail App Database not Connected")
			count++
		} else {
			log.Println("Mail App Database is Connected")
			return client
		}

		if count >= 5 {
			log.Println(err)
			return nil
		}

		log.Println("Wait:.... Mail App Database Retrying to Connect .....")
		time.Sleep(10 * time.Second)
		continue
	}
}
