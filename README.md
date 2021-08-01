# simple-golang-api
Golang API practice

Using [manato's article: Clean Architecture with GO](https://medium.com/manato/clean-architecture-with-go-bce409427d31).

# TODO

* v3 to practice concurrecy and channels
* implement Cloud Native Patterns
* utests and mocks
    * https://pkg.go.dev/github.com/gorilla/mux#readme-testing-handlers
* Add file logger
* kubernetes setup
* env file best practices
* implement more patterns from Cloud Native Go book
* use oauth
* add https security
* add db logger
* add db keystore 
* document API
* input validation in all layers
* DB seeder
* gRPC


# Usefull commands

```
docker build --tag kvstore key-store-api/ -f .
```

```
docker run --detach --publish 8080:8080 kvstore
```

```
go mod init
go mod tidy
```


