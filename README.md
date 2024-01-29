**Getting the Server Started**

To get started with the server, please ensure that Go is installed in your local machine. The original code was developed in Go1.21.6

After installing Go. Run the following commands in your terminal.

```terminal
go get -u github.com/gin-gonic/gin
go get -u github.com/stretchr/testify/assert
```


**Running the Server**

Run the following command in your project directory terminal to start the server.

```terminal
go run .
```


The server is now running on port 8080. You can send in API requests to this server and get a response. 
This server can further be tested by going into routes or student interface directory, and running the following command.

```terminal
go test
```

You can send API requests to this server through VS Extensions like REST client or Postman.