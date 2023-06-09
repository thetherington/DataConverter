BINARY_NAME=DataConverter

build:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} main.go

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}
