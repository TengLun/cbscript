package main

import (
	"cbs"
	"flag"

	"github.com/tenglun/logger"
)

func main() {

	api := flag.String("api_key", "", "Account Security Token to add items to blacklist")
	filename := flag.String("file", "", "CSV containing blacklist info to be sent to kochava")
	flag.Parse()

	logger := logger.CreateLogger("log")

	list, err := cbs.LoadBlacklist(logger, *filename)
	if err != nil {
		logger.Println(err)
	}

	err = cbs.SendList(logger, list, *api)
	if err != nil {
		logger.Println(err)
	}

	if err == nil {
		logger.Println("List sent correctly")
	}

}
