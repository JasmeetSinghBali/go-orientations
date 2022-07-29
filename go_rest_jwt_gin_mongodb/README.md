> REST api with go,jwt,gin,mongodb [with a basic structure]

                        # core STEPS
                        cd  go_rest_jwt_gin_mongodb


                        go mod init github.com/Jasmeet-1998/go-orientations/go_rest_jwt_gin_mongodb

                        # get gin-gonic
                        go get github.com/gin-gonic/gin


                        # api_structure
                        controllers [buisness logic]
                        repository [for DB interactions]
                        models [schema for DB storage]
                        routes [routing layer with all routes]
                        utils [misc helpers]
                        main.go [entry point]
                        go.mod [depend management]
