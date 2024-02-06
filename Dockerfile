FROM golang:latest AS student_api_server

WORKDIR /Users/asahran/Desktop/Internship Learning/task_2

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY . .
RUN go build -o main ./services/main/main.go
EXPOSE 8080
CMD ["./main"]


FROM alpine:latest AS db_server

RUN apk --no-cache add postgresql-client
WORKDIR /Users/asahran/Desktop/Internship Learning/task_2
COPY . .