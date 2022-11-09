awsbuild: clean
	GOOS=linux go build -o bin/storeBooking ./cmd/store_booking/main.go
build: clean
	go build -o bin/storeBooking ./cmd/store_booking/main.go
clean:
	rm -rf bin/*
run:
	go run cmd/store_booking/main.go
tdd:
	air
test: build
	go test ./test
zip: awsbuild
	cd bin && zip  storeBooking.zip storeBooking
