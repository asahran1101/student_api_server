# Backend API Server for Student Details

This project implements a backend server in GoLang using Gin package, complete with unit testing and implementation. It has 4 end points as follows.

* DELETE   /students/:rollNo   --> To delete a student by his roll number.
* GET      /students           --> To get the list of all the students.
* GET      /students/:rollNo   --> To get the details of a particuar student, using his roll number.
* POST     /students           --> To register a new student in the database.
* PUT      /students/:rollNo   --> To update the details of a student, using his roll number.

## Updating the .env File in the Project

Create a .env file in your local project directory and update it with the following code. Please note that the port number and username taken here are the ones assigned by default during installation of postgres. If you chose differently, enter those custom parameters accordingly.

```.env
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_DB=students
POSTGRES_PASSWORD=<Your Password>
```

## Updating the postgres-secret.yaml File in the Project

Create a postgres-secret.yaml file in ./charts/student-api-server/templates directory of this project directory. Copy the following lines in the file. 

```postgres-secret.yaml
apiVersion: v1
kind: Secret
metadata:
  name: {{ .Values.db.name }}-secret
  namespace: {{ .Values.app }}
data:
  POSTGRES_PASSWORD: <Base-64 encoded database password>
```

Replace the Base-64 encoded database password in the above file. You can obtain this by running the following command with your password on the terminal.

```terminal
echo -n '<Your Password>' | base64
```

Now copy and paste the output of this command in the postgres-secret.yaml file.

## Running the Server

Run the following commands in your project directory terminal to start the server.

```terminal
minikube start

kubectl delete namespace student-api-server
kubectl delete pv postgres-pv
helm install student-server ./charts/student-api-server 
kubectl get pods --namespace=student-api-server
```
You will get a list of all the running pods on your machine after this. Copy the pods name which starts with student-api-deployment. The run the following command in your terminal.

```terminal
kubectl port-forward <pod-name> --namespace=student-api-server 9922:8080
```
The server is now running on port 9292 on your local machine. You can send in API requests to this server and get a response. 

## Unit Testing for the server

This server can further be tested by running the following commands in your project directory terminal.

```terminal
cd ./student_core
go test
```

## Sending Requests

You can send API requests to this server through VS Extensions like REST client or Postman. Please note that the server is running on localhost:9922.

## Shutting Down the Server

The server can be shut down by going into the terminal window where the server is running and pressing Control + C on your keyboard.