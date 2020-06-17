package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/nexmo-community/nexmo-go"
// )

// func main() {
// 	googleAPIKey := os.Getenv("google_api_key")

// 	// Auth
// 	auth := nexmo.NewAuthSet()
// 	auth.SetAPISecret("API_KEY", googleAPIKey)

// 	// Init Nexmo
// 	client := nexmo.NewClient(http.DefaultClient, auth)

// 	// SMS
// 	smsContent := nexmo.SendSMSRequest{
// 		From: "+84938098955",
// 		To:   "+84938038621",
// 		Text: "This is a message sent from Go!",
// 	}

// 	smsResponse, _, err := client.SMS.SendSMS(smsContent)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Status:", smsResponse.Messages[0].Status)
// }

import (
	"fmt"

	"github.com/subosito/twilio"
)

func main() {
	// Common stuffs
	AccountSid := ""
	AuthToken := ""
	from := "+12058396624"
	to := "+84854337467"
	callbackURL := "http://tokoin.io/"

	// Initialize twilio client
	t := twilio.NewTwilio(AccountSid, AuthToken)

	// You can set custom Transport, eg: when you're using `appengine/urlfetch` on Google's appengine.
	// c := appengine.NewContext(r) // r is a *http.Request
	// t.Transport = urlfetch.Transport{Context: c}

	// Send SMS
	params := twilio.SMSParams{StatusCallback: callbackURL}
	s, err := t.SendSMS(from, to, "Hello Go!", params)

	// or, make a voice call
	// params := twilio.CallParams{Url: callbackURL}
	// s, err := t.MakeCall(from, to, params)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", s)
	return
}
