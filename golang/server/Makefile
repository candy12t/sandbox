BIN := bin/server

.PHONY: all
all: test clean build

.PHONY: build
build:
	go build -o $(BIN) .

.PHONY: clean
clean:
	rm -rf $(BIN)

.PHONY: test
test:
	go test -v ./... -count=1
