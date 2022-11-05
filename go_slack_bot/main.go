package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

/**
@desc - loop over the command events available option & displays them in terminal
*/
func displayCommandEvents(eventChannel <-chan *slacker.CommandEvent) {

	fmt.Println("Available command events for the SlackBot:")
	for event := range eventChannel {
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}

}

/**
@desc- entry point of the application
*/
func main() {

	/*load up required config env's*/
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		log.Fatal(err)
	}
	BotToken := os.Getenv("SLACK_BOT_TOKEN")
	AppToken := os.Getenv("SLACK_APP_TOKEN")
	myBot := slacker.NewClient(BotToken, AppToken)

	/*display available bot commands*/
	go displayCommandEvents(myBot.CommandEvents())
	myBot.Command("kida", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("wadiya")
		},
	})

	/* create a context with cancel event returned that is called for context termination to release up resources on current main function return*/
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	/*make the slackBot listen on the newly created context*/
	err = myBot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
