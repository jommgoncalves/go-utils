package email

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"time"
)

// Configs E-mail configurations.
type Configs struct {
	Server   string
	Port     int
	Username string
	Password string
	From     string
	To       []string
}

// SendWithoutAuthentication Sends E-mail without authentication.
func SendWithoutAuthentication(configs Configs, subject string, bodyHTML string) {
	// Convert [string,string,...] to "string,string,..""
	var toString string
	for i, e := range configs.To {
		if i == len(configs.To)-1 {
			toString += e
		} else {
			toString += e + ","
		}
	}

	// Set headers.
	headerMap := map[string]string{}
	headerMap["Content-Type"] = "text/html"
	headerMap["Charset"] = "UTF-8"
	headerMap["Subject"] = subject
	headerMap["Date"] = time.Now().String()
	headerMap["From"] = "deliverynotify@movvo.com"
	headerMap["To"] = toString

	// Create headers.
	header := ""
	for k, v := range headerMap {
		header += fmt.Sprintf("%s:%s\r\n", k, v)
	}

	// Connect to the remote SMTP server.
	address := fmt.Sprintf("%v:%v", configs.Server, configs.Port)
	c, err := smtp.Dial(address)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Set the sender and recipient.
	c.Mail(configs.From)
	for _, r := range configs.To {
		c.Rcpt(r)
	}

	// Send the E-mail body.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close()

	// Write headers.
	wc.Write([]byte(header))

	buf := bytes.NewBufferString(bodyHTML)
	if _, err = buf.WriteTo(wc); err != nil {
		log.Fatal(err)
	}
}
