# REST api in GO using PSQL and Docker

## Run

### In Docker
`docker-compose up`

### Directly

* Install dependencies
```
go get -u github.com/gorilla/mux
go get -u github.com/lib/pq
```
* Set DEVMACHINE environment variable
`export DEVMACHINE='true'`
* Finally run
`go run main.go`
* Or
`go install && $GOPATH/bin/psqlrestapi`

## TESTS

//TODO
