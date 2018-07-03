package main

import (
	"fmt"
	"github.com/robfig/cron"
)

func main() {
	i := 0
	c := cron.New()
	spec := "0 */1 * * * *"
	c.AddFunc(spec, func() {
		i++
		fmt.Println("execute per second", i)
	})
	c.Start()
	select {}
}

/*

go get -u github.com/robfig/cron

*/
