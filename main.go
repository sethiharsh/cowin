package main

import (
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	c.AddFunc("*/1 * * * *", CheckCenters)
	c.Start()

	for {
		time.Sleep(time.Second)
	}
}
