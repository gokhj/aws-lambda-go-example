package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
	subject := "This is a test"
	to := [2]string{"Gokhan", "hi@gokhanarkan.com"}
	content := "Hello world"
	
	response := send_mail(subject, to, content)

	if response == 202 {
		fmt.Println("Successful")
	} else {
		fmt.Println("Nooooo")
	}
}

func send_mail(email_subject string, address [2]string, content string) int {
	from := mail.NewEmail(os.Getenv("FROM_NAME"), os.Getenv("FROM_EMAIL"))
	subject := email_subject
	to := mail.NewEmail(address[0], address[1])
	plainTextContent := content
	htmlContent := content
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return -1
	} else {
		// response.Body
		// response.Headers
		return response.StatusCode
	}
}