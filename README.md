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

### SMTP settings for sending an email with the survey responses
This service is using SMTP to send an email with survey response data to the user.
This feature is **disabled** by default. To **enable** this feature you should update [**email**](https://github.com/yerzhan-karatay/survey-webapp-backend/blob/main/config/config.yaml#L14) and [**password**](https://github.com/yerzhan-karatay/survey-webapp-backend/blob/main/config/config.yaml#L15) in [**config.yaml**](https://github.com/yerzhan-karatay/survey-webapp-backend/blob/main/config/config.yaml) file by changing with your account credentials. I use GMAIL SMTP settings. If you want to GMAIL SMTP server, then you should create *App passwords* in your Google Account -> Security page. [Here](https://support.google.com/accounts/answer/185833?hl=en) is the tutorial to create App Passwords. 
1. Go to your Google Account and choose Security on the left panel.
2. On the Signing in to Google tab, select App Passwords.
3. At the bottom, choose Select app and choose the app you using and then Select device and choose the device you’re using and then Generate.

Follow the instructions to enter the App Password. The App Password is the 16-character code in the yellow bar on your device.


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
