package email

import (
	_ "crypto/tls"
	_ "fmt"
	"log"
	"os"
	_ "sync"

	"github.com/akinbyte/mailapp/model"
	"gopkg.in/gomail.v2"
)

func MailServer(mailChan model.Mail) {
	d := gomail.NewDialer("smtp.gmail.com", 465, os.Getenv("GMAIL_ACC"), os.Getenv("APP_PASSWORD"))

	s, err := d.Dial()
	if err != nil {
		log.Panicf("Error connecting to the Mail Server: ", err)
	}

	msg := gomail.NewMessage()

	msg.SetAddressHeader("From", mailChan.Source, os.Getenv("USERNAME"))
	msg.SetHeader("To", mailChan.Destination)
	msg.SetHeader("Subject", mailChan.Subject)
	msg.SetBody("text/html", mailChan.Message)

	if err := gomail.Send(s, msg); err != nil {
		log.Printf("Mail Sever : %s %v\n", mailChan.Destination, err)
	}
	msg.Reset()
}

func MailDelivery(mailChan chan model.Mail, worker int) {
	// Buffered channel for completion signals
	completionChan := make(chan bool, worker)

	for x := 0; x < worker; x += 1 {
		go func() {
			// Signal completion
			defer func() {
				completionChan <- true
			}()

			for m := range mailChan {
				MailServer(m)
			}
		}()
	}

	// Wait for all goroutines to complete
	for x := 0; x < worker; x += 1 {
		<-completionChan
	}
}


// func MailDelivery(mailChan chan model.Mail, worker int) {
// 	var wg sync.WaitGroup
// 	for x := 0; x < worker; x += 1 {
// 		wg.Add(1)
// 			for m := range mailChan {
// 				MailServer(m)
// 			}
// 			defer wg.Done()
// 	}
// 	wg.Wait()
// }