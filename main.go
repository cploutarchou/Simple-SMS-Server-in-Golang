package main

import (
	"encoding/json"
	"fmt"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"log"
	"net/http"
	"os"
)

var client *twilio.RestClient

func sendSMS(client *twilio.RestClient, to, body *string) (*twilioApi.ApiV2010Message, error) {
	from := "your number" // Replace with your Twilio phone number
	var message *twilioApi.ApiV2010Message
	// client params for the message
	msg := &twilioApi.CreateMessageParams{
		To:   to,
		From: &from,
		Body: body,
	}
	message, err := client.Api.CreateMessage(msg)
	log.Printf("Message sent to: %v\n", *msg.To)
	return message, err
}
func smsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		To   string `json:"to"`
		Body string `json:"body"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	message, err := sendSMS(client, &data.To, &data.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(struct {
		MessageID string `json:"message_id"`
	}{
		MessageID: *message.Sid,
	})
	if err != nil {
		return
	}
}

func main() {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	params := twilio.ClientParams{
		Username:   accountSid,
		Password:   authToken,
		AccountSid: accountSid,
	}
	client = twilio.NewRestClientWithParams(params)

	http.HandleFunc("/sms", func(w http.ResponseWriter, r *http.Request) {
		smsHandler(w, r)
	})

	port := "8080"
	fmt.Printf("SMS server is listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
