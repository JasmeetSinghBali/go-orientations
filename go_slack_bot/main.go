package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type apiConfigData struct {
	OpenWeatherApiKey string `json:"OpenWeatherApiKey"`
}

type weatherData struct {
	Name string `json: "name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

/**
@desc- loads openweather api key from .weatherApiConfig
*/
func loadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return apiConfigData{}, err
	}
	var c apiConfigData

	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfigData{}, err
	}

	return c, nil
}

/**
@desc - loop over the command events available option & displays them in terminal
*/
func displayCommandEvents(eventChannel <-chan *slacker.CommandEvent) {

	fmt.Println("event channel info for slack-bot:")
	for event := range eventChannel {
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}

}

/**
@desc - takes city name and returns weather via openapiweather call
*/
func tellCurrentWeather(city string) (weatherData, error) {
	apiConfig, err := loadApiConfig("./.weatherApiConfig")
	if err != nil {
		return weatherData{}, err
	}

	result, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherApiKey + "&q=" + city)
	if err != nil {
		return weatherData{}, err
	}
	defer result.Body.Close()

	var d weatherData
	if err := json.NewDecoder(result.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}

	return d, nil
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

	/*display event logs for slack bot*/
	go displayCommandEvents(myBot.CommandEvents())

	/*Greetings*/
	myBot.Command("kida", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("wadiya")
		},
	})

	/*tell weather a/c to city*/
	myBot.Command("weather {city}", &slacker.CommandDefinition{
		Description: "fed in a city name to know its current weather via openWeatherAPI",
		Examples:    []string{"weather delhi"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			city := request.Param("city")
			var data weatherData
			data, err = tellCurrentWeather(city)
			if err != nil {
				log.Fatal(err)
				response.Reply(`Sorry failed to fetch weather for provided city`)
			}
			log.Println("helloooooooooo")
			log.Println(data)
			currentTemp := data.Main.Kelvin - 273.15
			response.Reply("(In Celsius) Temp: " + strconv.FormatFloat(currentTemp, 'f', 10, 64))
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
