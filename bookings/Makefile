awsbuild: clean
	GOOS=linux go build -o bin/getBookings ./cmd/get_bookings/main.go
build: clean
	go build -o bin/storeBooking ./cmd/get_bookings/main.go
clean:
	rm -rf bin/*
run:
	go run cmd/get_bookings/main.go
tdd:
	air
test: build
	go test ./test
zip: awsbuild
	cd bin && zip getBookings.zip getBookings
