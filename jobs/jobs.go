package jobs

import (
	"github.com/robfig/cron"
)

func Run() *cron.Cron {
	c := cron.New()

	// Add Jobs here
	c.AddFunc("@every 1m", Greet)
	c.AddFunc("@every 1m", GreetingMail)

	c.Start()
	return c
}