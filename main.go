package main

import (
	"context"
	"log"
	"net/http"

	"github.com/akinbyte/mailapp/db"
	"github.com/akinbyte/mailapp/email"
	"github.com/akinbyte/mailapp/handlers"
	"github.com/akinbyte/mailapp/model"
	"github.com/joho/godotenv"
)

var (
	MailChan   chan model.Mail
	BufferSize int
	Worker     int
)

func main() {
	MailChan = make(chan model.Mail, BufferSize)
	Worker = 5


	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Starting the Mail App Server")

	log.Println("Preparing Database Connection")

	// connecting to mongodb database
	client := db.OpenConnect()
	defer func(ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			return
		}
	}(context.TODO())

	go email.MailDelivery(MailChan, Worker)
	
	defer close(MailChan)


	// getting access to the handlers
	app := handlers.NewMailApp(client, MailChan)

	handle := Routes(app)

	srv := http.Server{
		Addr:    ":8080",
		Handler: handle,
	}

	// helps to handle the incoming requests
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Shutting Down the Mail App Server ")
	}
}
