package main

import (
	api "live/kitex_gen/api/live"
	"log"
)

func main() {
	svr := api.NewServer(NewLiveImpl())

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
