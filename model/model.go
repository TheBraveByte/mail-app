package model

import "time"

// Subscriber : information or details from users
type Subscriber struct {
	FirstName string `bson:"first_name" json:"first_name"`
	LastName  string `bson:"last_name" json:"last_name"`
	Email     string `bson:"email" json:"email"`
	Interest  string `bson:"interest" json:"interest"`
}

// Mail: contains field of what a email entails
type Mail struct {
	Source      string
	Destination string
	Message     string
	Subject     string
	Name        string
}

// MailUpload - holds the upload contect be the admin
type MailUpload struct {
	DocxName    string    `bson:"docx_name" json:"docx_name"`
	DocxContent string    `bson:"docx" json:"docx"`
	Date        time.Time `bson:"date" json:"date"`
}
