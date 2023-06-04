package email

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/go-gomail/gomail"
	"github.com/yusuf/mailapp/model"
)

func MailServer(mailChan model.Mail) {
	d := gomail.NewDialer("smtp.gmail.com", 465, os.Getenv("GMAIL_ACC"), os.Getenv("APP_PASSWORD"))
	s, err := d.Dial()
	if err != nil {
		panic(err)
	}
	msg := gomail.NewMessage()

	msg.SetHeader("From", mailChan.Source)
	msg.SetAddressHeader("To", mailChan.Destination, mailChan.Name)
	msg.SetHeader("Subject", mailChan.Subject)
	msg.SetBody("text/html", fmt.Sprintf("%s", mailChan.Message))

	if err := gomail.Send(s, msg); err != nil {
		log.Printf("could not send mail to %v: %v", mailChan.Destination, err)
	}
	msg.Reset()
}

func MailDelivery(mailChan chan model.Mail) {
	worker := 5

	var wg sync.WaitGroup
	for x := 0; x < worker; x += 1 {
		wg.Add(1)
		defer wg.Done()
		for {
			MailServer(<-mailChan)
		}
	}
	wg.Wait()
}
