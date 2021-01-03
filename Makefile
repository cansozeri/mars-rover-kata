.PHONY:  lint

# build
build:
	go build ./...

# test
test:
	go test -v ./...

# run
run:
	go run cmd/mars-rover-kata/main.go
