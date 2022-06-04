BINARY_NAME=yahooFinance
MAIN=yahooFinance.go

all: build test

build:
	go build -o ${BINARY_NAME} ${MAIN}

test:
	go test -v

run:
	go build -o ${BINARY_NAME} ${MAIN}
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}
