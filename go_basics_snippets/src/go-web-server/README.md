> a simple golang web-server

- net/http

                # simple server with some routes
                Server------  /  ----> index.html
                |
                ------------ /hello ---> hello func
                |
                ------------ /form ---> form func ---> form.html

- run the server

                # navigate to go-web-server
                go run main.go

                routes
                /
                /home
                /form.html
