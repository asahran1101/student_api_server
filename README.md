# Backend API Server for Student Details

This project implements a backend server in GoLang using Gin package, complete with unit testing and implementation. It has 4 end points as follows.

* DELETE   /students/:rollNo   --> To delete a student by his roll number.
* GET      /students           --> To get the list of all the students.
* GET      /students/:rollNo   --> To get the details of a particuar student, using his roll number.
* POST     /students           --> To register a new student in the database.
* PUT      /students/:rollNo   --> To update the details of a student, using his roll number.

## Setting up PostgreSQL Database

Make sure that PostgreSQL is installed on your machine before running the server. The original code was developed on PostgreSQL14.10. After downloading, run the following command in your terminal. Please note that postgres here is the username, which is given by default. In case you chose a different username during installation, enter that instead.

```terminal
psql -U postgres
```

The terminal then asks for your PostgreSQL password, and opens up the PostgreSQL CLI. Here, enter the following commands.

```postgresql cli
CREATE DATABASE students;
```

## Updating the .env File in the Project

Create a .env file in your local project directory and update it with the follwing code. Please note that the port number and username taken here are the ones assigned by default during installation of postgres. If you chose differently, enter those custom parameters accordingly.

```.env
HOST=localhost
PORT=5432
USER=postgres
DB_NAME=students
PASSWORD=<Your Password>
```

## Running the Server

Run the following command in your project directory terminal to start the server.

```terminal
./build-tools/run-server.sh
```

The server is now running on port 8080. You can send in API requests to this server and get a response. 

## Unit Testing for the server

This server can further be tested by running the following commands in your project directory terminal.

```terminal
cd ./student_core
go test
```

## Sending Requests

You can send API requests to this server through VS Extensions like REST client or Postman. Please note that the server is running on localhost:8080.

## Shutting Down the Server

The server can be shut down by going into the terminal window where the server is running and pressing Control + C on your keyboard.