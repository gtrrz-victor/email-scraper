build:
	GOOS=linux go build -o bin/emailReader ./cmd/email_reader/main.go
run:
	go run emailReader.go
test:
	go test ./test
zip: build
	zip bin/emailReader.zip bin/emailReader
