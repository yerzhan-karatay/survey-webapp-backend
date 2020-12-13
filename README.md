# Survey backend service on GoLang
### 1. Install [GO - 1.15.6](https://golang.org/doc/install)
### 2. Install [go-swagger](https://github.com/go-swagger/go-swagger/blob/master/docs/install.md)


### Run the service locally
```
make run
```
or
```
go build -o build/survey-webapp-backend cmd/survey-webapp-backend/main.go
```


### Build
```
make
```
or
```
go build -o build/survey-webapp-backend cmd/survey-webapp-backend/main.go
```


### Test
```
make test
```
or
```
go test ./cmd/...
```