# run
run:
	go run cmd/survey-webapp-backend/main.go

# build
build:
	mkdir build
	go build -o build/survey-webapp-backend cmd/survey-webapp-backend/main.go

# test
test:
	go test ./cmd/...

# validate swagger spec
swagger.validate:
	swagger validate pkg/swagger/swagger.yml
