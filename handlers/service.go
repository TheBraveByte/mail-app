package handlers

import "net/http"

type Logic interface {
	Home() http.HandlerFunc
	GetSubscriber() http.HandlerFunc
	SendMail() http.HandlerFunc
}
