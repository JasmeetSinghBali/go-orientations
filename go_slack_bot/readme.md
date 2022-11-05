> Preq:

1. login to slack account & create an app [at](https://api.slack.com/)
2. from scratch
3. enable socket mode
4. use generated app token
5. enable all event subscription option and enable events & add permissions read/write messages to bot along with Oauth scopes.
6. go to Oauth & permissions install bot to workspace & and copy the bot token keep it safe would be needed to setup commands/intents for bot.
7. each time permissions or settings of bot changes it has to be reinstalled in slack workspace in the Oauth section
8. build the bot and mention on the channel & interact with it

> Notes

Under features

1. Oauth bot token scopes

appmentions :read
channel: history
channel: read
chat:write
im:history
im:read
im:write
mpim:history
mpim:read
mpim:write

2. Event subscriptions

app_mentions
message.im
message.mpim
message.groups
message.channels

> Dev walkthrough

                # cd to poc dir
                go mod init github.com/Jasmeet-1998/go-orientations/go_slack_bot
                go get "github.com/shomali11/slacker"

                #build
                go build

                #run
                go run main.go

                # go to slack workspace for which u created app
                #navigate to channel say general
                #enter
                @yourBotName ping
