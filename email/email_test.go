package email

import "testing"

func TestSendWithoutAuthentication(t *testing.T) {
	// E-mail configs.
	configs := Configs{
		Server: "aspmx.l.google.com",
		Port:   25,
		From:   "deliverynotify@movvo.com",
		To:     []string{"jorge.goncalves@movvo.com"}}

	SendWithoutAuthentication(configs, "Email test.", "This is body of the email.")
}
