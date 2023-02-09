package handler

import (
	"go.tienngay/mailer/database"
	"github.com/gofiber/fiber/v2"
	"go.tienngay/pkg/mysql/repositories/mail"
	"go.tienngay/pkg/mysql/entities"
	"github.com/sendgrid/sendgrid-go"
    mailSG "github.com/sendgrid/sendgrid-go/helpers/mail"
    "go.tienngay/mailer/config"
	"fmt"
	"sync"
)

func MailService() mail.Service {
	db := database.DB.Table("mails")
	repo := mail.NewRepo(db)
	service := mail.NewService(repo)
	return service;
}

func SendEmail(c *fiber.Ctx) error {
	fmt.Println("SendEmail begin")
	var mailService = MailService()
	var mails *[]entities.Mail
	mails = mailService.GetWaitingMails()

	for _, mail := range (*mails) {
		var mailService3 = MailService()
		mailService3.UpdateStatusSending(mail.ID)
	}
	chanMails := make(chan entities.Mail, len(*mails))
	var wg sync.WaitGroup
	for _, mail := range (*mails) {
		wg.Add(1)
		go func(mail entities.Mail, chanMails chan entities.Mail) {
			defer wg.Done()
			fmt.Println("Begin Sent email")
			fmt.Println(mail.ID)
			// Initialise the required mail message variables
			from := mailSG.NewEmail(mail.NameFrom, mail.From)
			subject := mail.Subject
			to := mailSG.NewEmail("", mail.To)
			plainTextContent := ""
	    	htmlContent := mail.Message
	    	message := mailSG.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	    	// Attempt to send the email
		    client := sendgrid.NewSendClient(config.Config("SENDGRID_API_KEY"))
		    response, err := client.Send(message)
		    if err != nil {
		        fmt.Println("Unable to send your email")
		        mail.Errors = "Unable to send your email"
		    }
		    // Check if it was sent
		    statusCode := response.StatusCode
		    if statusCode == 200 || statusCode == 201 || statusCode == 202 {
		    	mail.Errors = "Success"
		        fmt.Println("Email sent!")
		    }
		    chanMails <- mail
		}(mail, chanMails)
	}
	wg.Wait()
	close(chanMails)
    fmt.Println("Start insert log")
	for chanMail := range chanMails {
		fmt.Println(chanMail.Errors)
		var mailService2 = MailService()
		if (chanMail.Errors == "Success") {
			mailService2.UpdateStatusSuccess(chanMail)
		} else {
			mailService2.UpdateStatusErrors(chanMail)
		}
	}
	fmt.Println("End insert log")

	return c.JSON(fiber.Map{"status": "success", "message": "Success", "data": mails})
}
