build:
    mkdir build
    go build -o build/basicwebapp main.go

clean:
    rm -rf build

check:
    go vet ./...
    errcheck ./...

fmt:
    go fmt ./...

test:
    go test -v ./...
