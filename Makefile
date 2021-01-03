.PHONY:

# build
build:
	go build -o build/main cmd/mars-rover-kata/main.go

# test
test:
	go test -v ./...

# run
run:
	go run cmd/mars-rover-kata/main.go -fPath="cmd/mars-rover-kata/instructions.txt"
