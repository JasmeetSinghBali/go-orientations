> REST api with go,jwt,gin,mongodb [with a basic structure]

                        # core STEPS
                        cd  go_rest_jwt_gin_mongodb


                        go mod init github.com/Jasmeet-1998/go-orientations/go_rest_jwt_gin_mongodb

                        # get gin-gonic
                        go get github.com/gin-gonic/gin


                        # api_structure
                        controllers [buisness logic]
                        models [schema for DB storage]
                        routes [routing layer with all routes]
                        utils [misc helpers]
                        middlewares [middlewares]
                        main.go [entry point]
                        go.mod [depend management]

> IMP📝 context ref:https://pkg.go.dev/context

> ### To get started

                clone the repo
                cd go_rest_jwt_***
                go mod tidy [to resolve dependency]
                go run main.go
