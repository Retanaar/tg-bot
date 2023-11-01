package main

import (
	"flag"
	"log"
)

func main() {
	t := mustToken()

	//tgClient := telegram.New(token)

	//fetcher := fetcher.New()

	//processor := processor.New()

	//consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String(
		"tg-token",
		"",
		"token for telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatalf("token is not specified")
	}

	return *token
}
