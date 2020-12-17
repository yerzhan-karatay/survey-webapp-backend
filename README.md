# _Survey backend service on GoLang_

This is part of Survey web application. To run the whole service, you need [**survey-webapp-frontend**](https://github.com/yerzhan-karatay/survey-webapp-frontend) and [**survey-webapp-backend**](https://github.com/yerzhan-karatay/survey-webapp-backend) to be run on your local machine together.

### 1. Install [GO - 1.15.6](https://golang.org/doc/install)

### 2. Set up local environment
Clone [**survey-webapp-backend**](https://github.com/yerzhan-karatay/survey-webapp-backend) project to your local.
```
git clone git@github.com:yerzhan-karatay/survey-webapp-backend.git
cd survey-webapp-backend
```

### 3. Run the service locally
```
make run
```
or
```
go build -o build/survey-webapp-backend cmd/survey-webapp-backend/main.go
```

### API documentation is accessible via Swagger after `make run` command on this link - [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

**NOTE!** 
##### Most APIs are only accessible for authorized users.
You should Authorize in Swagger by adding a user token 
The **Bearer token** should be added like:
```
Bearer YOUR_TOKEN
```
![Bearer token example](/docs/authorization-example.png)

### DB ER diagram
![ER diagram](/docs/survey-db-ER.png)

### APIs draft
![APIs draft](/docs/apis-structure.png)

### Other commands:
#### - Build

```
make
```
or
```
go build -o build/survey-webapp-backend cmd/survey-webapp-backend/main.go
```

#### - Test
```
make test
```
or
```
go test ./cmd/...
```
