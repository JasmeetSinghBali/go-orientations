> ## aws lambda with Go

- Lambda serverless functions that executes business logic and each time this executes then only u need to pay

- 📝 great for MVP | UAT product stage development as we are unaware of the user traffic u only have to pay when some end user use your service via lambda's

- 📝 further from POC's -> UAT|MVP -> final launch aws lambda can be used bcz scaling is easy in aws lambda as it is fully managed service, +you dont pay for server, you only pay when someone is using the program.

- 📝aws lambda can be used to expose your api's also.

> ## ⛔ Drawbacks of aws lambda

- coldstart, lambda takes time to start when someone is trying to access your lambda functions.

> ### 🎯 Aim- build and deploy aws lambda function via aws-cli in golang
>
> ref: https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html

1.  init

                    #cd go-aws-lambda-snippet
                    go mod init github.com/Jasmeet-1998/go-orientations/go_basics_snippets/src/go-aws-lambda-snippet

2.  basic go snippet

3.  deploying as aws lambda

a) Create an IAM user [from terminal]

                    # make sure aws-cli setup properly in your system


                    aws iam create-role --role-name lambda-ex --assume-role-policy-document '{"Version": "2012-10-17","Statement": [{"Effect": "Allow","Principal":{"Service":"lambda.amazonaws.com"},"Action":"sts:AssumeRole"}]}'

b) create a file called 'trust-policy.json' in the project root i.e go-aws-lambda-snippet

c) signal and configure aws-cli & point to this trust-policy.json [from terminal]

                    aws iam create-role --role-name lambda-ex --assume-role-policy-document file://trust-policy.json

d) finally attach policy and iam roles together [from terminal]

                    aws iam attach-role-policy --role-name lambda-ex --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole

e) build the go file

                    go build main.go
                    # error
                    go mod tidy
                    go build main.go

f) create a zip for the main file created after build

                    #homebrew for mac, choco for windows powershell, apt install in linux to install zip
                    # takes the main file and zips into function.zip

                    zip function.zip main

g)📝 create aws lambda function from terminal from the zip file created

                    aws lambda create-function --function-name go-aws-lambda-snippet \--zip-file fileb://function.zip --handler main --runtime go1.x \--role arn:aws:iam::318784325735:role/lambda-ex

                    # --function-name yourFunctionName cross check in go.mod
                    # handler is main as main is inside function.zip
                    # zip file
                    # runtime golang as go 1.x version
                    # 318784325735 - amazon account ID

h) 📝 invoking the lambda function

                    aws lambda invoke --function-name go-aws-lambda-snippet --cli-binary-format raw-in-base64-out --payload '{"what is your name?":"John","How old are you": 33}' output.txt

                    # the response or err will go into output.txt
                    # raw-in-base64-out to pass json as payload to the lambda function
