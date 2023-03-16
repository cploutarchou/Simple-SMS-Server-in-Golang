# Building a Twilio SMS Sender in Go: A Practical Guide

This code package contains a simple HTTP server that can receive POST requests and send SMS messages using Twilio API.

### Setup
To use this code package, you need to set your Twilio Account SID and Auth Token as environment variables. You can do this by running the following commands in your terminal:

```bash
export TWILIO_ACCOUNT_SID=your_account_sid
export TWILIO_AUTH_TOKEN=your_auth_token
```
Replace `your_account_sid` and `your_auth_token` with your actual Twilio credentials.

### Usage
Once you have set your Twilio credentials, you can start the server by running the following command:

```bash
go run main.go
```
This will start the server on port 8080. You can then send a POST request to http://localhost:8080/sms with the following JSON payload:

```json
{
"to": "+1234567890",
"body": "Hello from Twilio!"
}
```
Replace `+1234567890` with the phone number you want to send the SMS to, and replace Hello from Twilio! with the message you want to send.

Alternatively, you can use the following curl command to send the POST request:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"to": "+1234567890", "body": "Hello, Golang!"}' http://localhost:8080/sms
```
The server will then use the Twilio API to send the SMS, and return a JSON response with the message ID:

```json
{
"message_id": "SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}
```
### Dependencies
This code package uses the following third-party packages:

* github.com/twilio/twilio-go
* github.com/twilio/twilio-go/rest/api/v2010

You can install these packages using the following command:
```bash
go get github.com/twilio/twilio-go github.com/twilio/twilio-go/rest/api/v2010
```
### License
This code package is distributed under the MIT license. See LICENSE for more information.
