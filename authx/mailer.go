package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/damonkeys/covid-checkin/monkeys/tracing"
	"github.com/mailgun/mailgun-go/v4"
)

// Your available domain names can be found here:
// (https://app.mailgun.com/app/domains)
const yourDomain = "[your-mailgun-domain]" // TODO: replace with env variables

// You can find the Private API Key in your Account Menu, under "Settings":
// (https://app.mailgun.com/app/account/security)
const privateAPIKey = "[your-mailgun-api-key]" // TODO: replace with env variable

// CTAMailContext contains all data that is need for the typical mail with a single call to cation (CTA)
// thus should be used with the cta template (cta-tpl) of our mailgun accountä-ö.
type CTAMailContext struct {
	templatename string
	recipient    string
	sender       string
	subject      string
	cta          string
	body         string
	ctalink      string
	linktext     string
}

// MGMessagePreparer prepares a mailgun message to be send. The actual how it creates the message
// depends on the struct binding and its implementation
type MGMessagePreparer interface {
	prepareMessage(mailgun mailgun.Mailgun) *mailgun.Message
}

func (mailContext CTAMailContext) prepareMessage(mailgun mailgun.Mailgun) *mailgun.Message {
	message := mailgun.NewMessage(mailContext.sender, mailContext.subject, fmt.Sprintf("%s %s: %s", mailContext.body, mailContext.linktext, mailContext.ctalink), mailContext.recipient)
	message.SetTemplate(mailContext.templatename)
	message.AddVariable("cta", mailContext.cta)
	message.AddVariable("body", mailContext.body)
	message.AddVariable("ctalink", mailContext.ctalink)
	message.AddVariable("linktext", mailContext.linktext)

	return message
}

func sendMail(preparer MGMessagePreparer) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	span := tracing.EnterWithContext(ctx)
	defer span.Finish()

	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)

	tracing.LogStruct(span, "mail-data", mg)

	message := preparer.prepareMessage(mg)

	// Send the message	with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)
	if err != nil {
		tracing.LogErrorMsg(span, err, "Mail not sent")
	}
	tracing.LogString(span, "response from mailgun", resp)
	tracing.LogString(span, "id from mailgun", id)
}

func sendTestMail() {
	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)

	sender := "sender@example.com"
	subject := "Fancy subject!"
	body := "Hello from Mailgun Go!"
	recipient := "support@chckr.de"

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message	with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Test mail with: ID %s and Resp: %s\n", id, resp)
}
