clean:
	rm -rf bin/*

build:
	go build -o bin/service main.go

image:
	docker build -t ghcr.io/bryopsida/go-libp2p-starter:local .

test:
	go test -v ./...
	
lint:
	go install golang.org/x/lint/golint@latest
	golint ./...
	go vet ./...
