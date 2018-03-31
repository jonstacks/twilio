package main

import (
	"fmt"
	"os"

	"github.com/jonstacks/twilio"
)

// Twilio Variables that will be set during compilation. See README.md
var (
	TwilioAccountSID string
	TwilioAuthToken  string
	TwilioFromNumber string
)

func printHelp() {
	fmt.Printf(`notify-me. Send text Messages from the command line.

Usage:
  notify-me <number> <message>
	notify-me -h | --help

Options:
  -h  --help  Show this screen

Example:
  notify-me "+12345678910" "Hello from server 123"
`)
}

func main() {
	debug := os.Getenv("DEBUG") == "True"

	for _, arg := range os.Args {
		if arg == "-h" || arg == "--help" {
			printHelp()
			os.Exit(1)
		}
	}

	if len(os.Args) != 3 {
		fmt.Println("Unable to parse arguments.")
		printHelp()
		os.Exit(1)
	}

	to := os.Args[1]
	body := os.Args[2]

	client := twilio.NewClient(twilio.APIBase, TwilioAccountSID, TwilioAuthToken)
	msg := &twilio.Message{
		To:   to,
		From: TwilioFromNumber,
		Body: body,
	}

	if debug {
		fmt.Printf("Sending '%s' to '%s'\n", body, to)
	}

	err := client.SendMessage(msg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
