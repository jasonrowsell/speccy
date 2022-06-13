gen:
	protoc --proto_path=protobuf protobuf/*.proto --go_out=plugins=grpc:pb

clean:
	rm pb/*.go

client:
	go run cmd/client/main.go -server 0.0.0.0:8080

server:
	go run cmd/server/main.go -port 8080

test:
	go test -race -v ./...

test-coverage:
	go test -race -v ./... -cover -coverprofile=coverage.out -covermode=atomic

# Formats the coverage report to HTML.
format-coverage:
  go tool cover -html=coverage.out -o coverage.html
