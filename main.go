package main

import (
	"context"
	"log"
	"net/http"
	_ "time"

	"github.com/joho/godotenv"
	"github.com/yusuf/mailapp/db"
	"github.com/yusuf/mailapp/email"
	"github.com/yusuf/mailapp/handlers"
	"github.com/yusuf/mailapp/model"
)

var (
	MailChan   chan model.Mail
	BufferSize int
)

func main() {
	MailChan = make(chan model.Mail, BufferSize)

	defer close(MailChan)

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

	go email.MailDelivery(MailChan)

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
