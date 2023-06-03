package db

import (
	"github.com/yusuf/mailapp/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DataStore interface {
	AddSubscriber(subs *model.Subscriber) (bool, string, error)
	AddMail(mu model.MailUpload) (string, error)
	FindSubscribers() ([]primitive.M, error)
}
