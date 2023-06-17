package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/akinbyte/mailapp/db"
	"github.com/akinbyte/mailapp/model"
	"github.com/akinbyte/mailapp/tools"

	"go.mongodb.org/mongo-driver/mongo"
)

type MailApp struct {
	MailDB   db.DataStore
	MailChan chan model.Mail
}

func NewMailApp(client *mongo.Client, mailchan chan model.Mail) Logic {
	return &MailApp{
		MailDB:   db.NewMongo(client),
		MailChan: mailchan,
	}
}

func (ma *MailApp) Home() http.HandlerFunc {
	return func(wr http.ResponseWriter, rq *http.Request) {
		err := tools.HTMLRender(wr, rq, nil)
		if err != nil {
			log.Println(err)
		}
	}
}

// GetSubscriber: this will process the registration of the subscribers
func (ma *MailApp) GetSubscriber() http.HandlerFunc {
	return func(wr http.ResponseWriter, rq *http.Request) {
		var subs model.Subscriber
		subscriber, err := tools.JSONReader(wr, rq, subs)
		if err != nil {
			http.Error(wr, fmt.Sprintf("failed to read json : ", err), http.StatusBadRequest)
		}
		ok, msg, err := ma.MailDB.AddSubscriber(subscriber)
		if err != nil {
			http.Error(wr, msg, http.StatusInternalServerError)
		}
		switch ok {
		case msg == "":
			tools.JSONWriter(wr, "You have already registered", http.StatusOK)
		case msg != "":
			tools.JSONWriter(wr, msg, http.StatusOK)
		}
	}
}

// SendMail: this will send the uploaded mail to the subscribers and have it save in the database
func (ma *MailApp) SendMail() http.HandlerFunc {
	return func(wr http.ResponseWriter, rq *http.Request) {
		var mailUpload model.MailUpload
		upload, err := tools.ReadForm(wr, rq, mailUpload)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusBadRequest)
		}

		msg, err := ma.MailDB.AddMail(upload)
		if err != nil {
			http.Error(wr, msg, http.StatusInternalServerError)
		}
		log.Println(msg)
		log.Println("........ preparing to send mail to subscribers ........ ")
		time.Sleep(time.Millisecond)
		log.Println("........ Accessing the subscribers Database ........ ")

		res, err := ma.MailDB.FindSubscribers()
		if err != nil {
			http.Error(wr, fmt.Sprintf("SendMail: failed query: %v", err), http.StatusInternalServerError)
		}

		for _, s := range res {
			subEmail := s["email"].(string)
			firstName := s["first_name"].(string)
			lastName := s["last_name"].(string)

			subName := fmt.Sprintf("%s %s", firstName, lastName)
			mail := model.Mail{
				Source:      os.Getenv("GMAIL_ACC"),
				Destination: subEmail,
				Name:        subName,
				Message:     upload.DocxContent,
				Subject:     upload.DocxName,
			}
			ma.MailChan <- mail
		}
		err = tools.JSONWriter(wr, fmt.Sprintf("Mail Sent %v subscribers", len(res)), http.StatusOK)
		if err != nil {
			http.Error(wr, err.Error(), http.StatusInternalServerError)
		}
	}
}
