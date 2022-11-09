awsbuild: clean
	GOOS=linux go build -o bin/emailReader ./cmd/email_reader/main.go
build: clean
	go build -o bin/emailReader ./cmd/email_reader/main.go
run:
	go run emailReader.go
clean:
	rm -rf bin/*
test: build
	go test ./test
tdd:
	air
zip: awsbuild
	cd bin && zip emailReader.zip emailReader
