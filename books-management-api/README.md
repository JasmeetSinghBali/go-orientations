> Books managment api

- go.mod file just like package.json file in Nodejs.

            cd books-management-api
            go mod init github.com/githubUsername/go-orientations

- installing needed third party packages

            # gorm is ORM[object-relational mapper] to interact with the database
            go get "github.com/jinzhu/gorm"
            go get "github.com/jinzhu/gorm/dialects/mysql"

            # gorilla/mux web-server
            go get "github.com/gorilla/mux"

- project structure

            - cmd
             - main
                *main.go
            - pkg
             - config
             - controllers
             - models
             - routes
             - utils

> NOTE- in golang we need to provide absolute paths relative path with reff to current dir or like so is not possible in go.

> NOTE- to make sense and process data in controllers we need to always unmarshal i.e deserialize i.e parse the json data coming in the request payload. ref: utils.go in utils under pkg folder

            # navigate to cmd->main/main.go
            go build

            #resolve all errors if their are any

            #finally
            go run main.go
