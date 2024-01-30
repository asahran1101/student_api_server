# Backend API Server for Student Details

This project implements a backend server in GoLang using Gin package, complete with unit testing and implementation. It has 4 end points as follows.

* DELETE   /students/:rollNo   --> To delete a student by his roll number.
* GET      /students           --> To get the list of all the students.
* GET      /students/:rollNo   --> To get the details of a particuar student, using his roll number.
* POST     /students           --> To register a new student in the database.
* PUT      /students/:rollNo   --> To update the details of a student, using his roll number.

## Getting the Server Started

To get started with the server, please ensure that Go is installed in your local machine. The original code was developed in Go1.21.6

After installing Go. Run the following commands in your terminal.

```terminal
$ go get -u github.com/gin-gonic/gin
$ go get -u github.com/stretchr/testify/assert
$ go get -u github.com/golang/mock
```

## Setting up PostgreSQL



## Updating the .env File in the Project

Create a .env file in your local project directory and update it with the follwing code.

```.env
HOST=localhost
PORT=5432
USER=<Your Username>
DB_NAME=students
PASSWORD=<Your Password>
```

## Running the Server

Run the following command in your project directory terminal to start the server.

```terminal
go run .
```


The server is now running on port 8080. You can send in API requests to this server and get a response. 

## Unit Testing for the server

This server can further be tested by running the following commands in your working directory.

```terminal
cd ./student_core
go test
```

## Sending Requests

You can send API requests to this server through VS Extensions like REST client or Postman. Please note that the server is running on localhost:8080.

## Shutting Down the Server

The server can be shut down by going into the terminal window where the server is running and pressing Control + C on your keyboard.