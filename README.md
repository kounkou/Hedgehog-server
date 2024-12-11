# Hedgehog-server

Hedgehog-server is the server side for the Hedgehog application. Hedgehog-server is made of 2 parts : 

1- The server application in Golang
2- The mongo db that has all the questions

When Hedgehog connects to Hedgehog-server currently on : "http://localhost:8080/questions", the Golang
server will make a request to the mongo-db side and retrieve all the questions.
The questions will then be used for further processing on the Hedgehog application.

### How to launch

Make sure your docker engine is running

In a Terminal, Inside `Hedgehog-server`

```bash
docker-compose down
docker-compose build --no-cache
docker-compose up
```

Then you can launch `Hedgehog` in another Terminal
