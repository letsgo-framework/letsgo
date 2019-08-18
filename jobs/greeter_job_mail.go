package jobs

import (
	"github.com/letsgo-framework/letsgo/mail"
)

func GreetingMail()  {
	mail.SendMail([]string{"greet@letsgo.com"}, "Greetings", "greeter-template.html", struct {
		App string
	}{
		App: "LetsGO",
	})
}
