# letsgo
[![Build Status](https://travis-ci.org/letsgo-framework/letsgo.svg?branch=master)](https://travis-ci.org/letsgo-framework/letsgo)
[![Go Report Card](https://goreportcard.com/badge/github.com/letsgo-framework/letsgo)](https://goreportcard.com/report/github.com/letsgo-framework/letsgo)
[![Coverage Status](https://coveralls.io/repos/github/letsgo-framework/letsgo/badge.svg?branch=master)](https://coveralls.io/github/letsgo-framework/letsgo?branch=master)
[![Sourcegraph](https://sourcegraph.com/github.com/letsgo-framework/letsgo/-/badge.svg)](https://sourcegraph.com/github.com/letsgo-framework/letsgo?badge)
[![Join the chat at https://gitter.im/letsgo-framework/community](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/letsgo-framework/community)

## Go api starter
### Ingredients

- Go
- [gin] ( https://github.com/gin-gonic/gin )
- [mongodb] ( https://www.mongodb.com/ )
- [mongo-go-driver] ( https://github.com/mongodb/mongo-go-driver )
- [oauth2] ( https://github.com/golang/oauth2 )
- [check] ( https://godoc.org/gopkg.in/check.v1 )
- [godotenv] ( https://github.com/joho/godotenv )
- [go-oauth2/gin-server] ( github.com/go-oauth2/gin-server )
- [cors] ( github.com/gin-contrib/cors )
***
### Directory Structure

By default, your project's structure will look like this:

- `/controllers`: contains the core code of your application.
- `/database`: contains mongo-go-driver connector.
- `/helpers`: contains helpers functions of your application.
- `/middlewares`: contains middlewares of your application.
- `/routes`: directory contains RESTful api routes of your application.
- `/tests`: contains tests of your application.
- `/types`: contains the types/structures of your application.
***
### Environment Configuration

letsGo uses `godotenv` for setting environment variables. The root directory of your application will contain a `.env.example` file.
copy and rename it to `.env` to set your environment variables.

You need to create a `.env.testing` file from `.env.example` for running tests.
***
### Setting up

- clone letsGo
- change package name in `go.mod` to your package name
- change the internal package (controllers, tests, helpers etc.) paths as per your requirement
- setup `.env` and `.env.testing`
- run `go mod download` to install dependencies

#### OR `letsgo-cli` can be used to setup new project

### install letsgo-cli
```
go get github.com/letsgo-framework/letsgo-cli
```


### Create a new project

```bash
letsgo-cli init <import_namespace> <project_name>
```

- **letsgo-cli init github.com myapp**<br/>
  Generates a new project called **myapp** in your `GOPATH` inside `github.com` and installs the default plugins through the glide.
***
### Run : ```go run main.go```
***
### Build : ```go build```
***
### Test : ```go test tests/main_test.go```

### Coverall :
```
go test -v -coverpkg=./... -coverprofile=coverage.out ./...

goveralls -coverprofile=coverage.out -service=travis-ci -repotoken $COVERALLS_TOKEN
```
***
### Authentication

letsgo uses Go OAuth2 (https://godoc.org/golang.org/x/oauth2) for authentication.
***

### Deploy into Docker

```
sudo docker run --rm -v "$PWD":/go/src/github.com/letsgo-framework/letsgo -w /go/src/github.com/letsgo-framework/letsgo iron/go:dev go build -o letsgo
```
```
sudo docker build -t sab94/letsgo .
```
```
sudo docker run --rm -p 8080:8080 sab94/letsgo
```

# Thank You
[![https://www.jetbrains.com/?from=letsgo](https://user-images.githubusercontent.com/15252513/68403814-8f147100-01a3-11ea-9dbc-f51d36de3ef3.png)](https://www.jetbrains.com/?from=letsgo)
