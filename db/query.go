package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/yusuf/mailapp/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	MailDB *mongo.Client
}

func NewMongo(client *mongo.Client) DataStore {
	return &Mongo{MailDB: client}
}

// AddSubscriber - this will process the subscriber details and have that added in the dataase
func (mg *Mongo) AddSubscriber(subs model.Subscriber) (bool, string, error) {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelCtx()
	var res bson.M
	filter := bson.D{{Key: "email", Value: subs.Email}}
	err := Default(mg.MailDB, "subscribers").FindOne(ctx, filter).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			_, err := Default(mg.MailDB, "subscribers").InsertOne(ctx, subs)
			if err != nil {
				return false, "", fmt.Errorf("AddSubscriber: cannot registered this account : %v", err)
			}
			return true, fmt.Sprintf("New Subscriber Added"), nil
		}
		log.Fatalln("AddSubscriber: cannot query database", err.Error())
	}
	return true, "", nil
}

// AddMail - this will allow the mai user/admin to upload new newletter mail to his/her subscriber
func (mg *Mongo) AddMail(mu model.MailUpload) (string, error) {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelCtx()
	// all the uploaded mail to the database
	_, err := Default(mg.MailDB, "mails").InsertOne(ctx, mu)
	if err != nil {
		return "", fmt.Errorf("AddMail: unable to add new mail, %v", err)
	}
	return fmt.Sprint("New mail successfully added"), nil
}

// FindSubscribers - this query the database to get all the mail to sent to subscribers
func (mg *Mongo) FindSubscribers() ([]primitive.M, error) {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelCtx()

	var res []bson.M
	// get all suscribers data
	cursor, err := Default(mg.MailDB, "subscribers").Find(ctx, bson.D{{}})
	if err != nil {
		return []bson.M{}, err
	}
	if err = cursor.All(ctx, &res); err != nil {
		return []bson.M{}, fmt.Errorf("FindMail: Cannot get all mail: %v", err)
	}
	defer cursor.Close(ctx)

	if err = cursor.Err(); err != nil {
		return []bson.M{}, fmt.Errorf("FindMail: Cursor Error : %v", err)
	}
	return res, nil
}
