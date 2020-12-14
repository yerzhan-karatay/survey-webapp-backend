# Survey backend service on GoLang
### Install [GO - 1.15.6](https://golang.org/doc/install)

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

### API doc accessible via Swagger
[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
