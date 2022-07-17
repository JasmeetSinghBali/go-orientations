> ## aws lambda with Go

- Lambda serverless functions that executes business logic and each time this executes then only u need to pay

- 📝 great for MVP | UAT product stage development as we are unaware of the user traffic u only have to pay when some end user use your service via lambda's

- 📝 further from POC's -> UAT|MVP -> final launch aws lambda can be used bcz scaling is easy in aws lambda as it is fully managed service, +you dont pay for server, you only pay when someone is using the program.

- 📝aws lambda can be used to expose your api's also.

> ## ⛔ Drawbacks of aws lambda

- coldstart, lambda takes time to start when someone is trying to access your lambda functions.

> ### 🎯 Aim- build and deploy aws lambda function via aws-cli in golang

1.  init

                    #cd go-aws-lambda-snippet
                    go mod init github.com/Jasmeet-1998/go-orientations/go_basics_snippets/src/go-aws-lambda-snippet

2.  basic go snippet
