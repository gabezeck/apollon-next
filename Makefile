BINARY=reddy.so
 
build:
    go build -o ${BINARY} ./cmd/main.go
 
run-local:
    go build -o ${BINARY} ./cmd/main.go -- config-file .env
    ./${BINARY}
 
clean:
    go clean
    rm ${BINARY}