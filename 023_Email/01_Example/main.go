// using SendGrid's Go Library
// https://github.com/sendgrid/sendgrid-go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func testSendMail1() {
	fromName := "Tokoin"
	fromEmail := "care@tokoin.io"
	from := mail.NewEmail(fromName, fromEmail)

	toName := "Akagi"
	toEmail := "phuong.nguyen@tokoin.io"
	to := mail.NewEmail(toName, toEmail)

	subject := "Sending with SendGrid is Fun"

	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func testSendMail2() {
	fromName := "Akagi"
	fromEmail := "phuong.nguyen@tokoin.io"
	from := mail.NewEmail(fromName, fromEmail)

	toName := "Nguyen Hoai Phuong"
	toEmail := "hoaiphuong.nguyen.vn@gmail.com"
	to := mail.NewEmail(toName, toEmail)

	subject := "Sending email by sendgrid is so fun"

	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func main() {
	// testSendMail1()
	testSendMail2()
}
