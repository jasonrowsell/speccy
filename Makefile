gen:
	protoc --proto_path=protobuf protobuf/*.proto --go_out=plugins=grpc:pb

clean:
	rm pb/*.go

test-unit:
	go test -v ./...

test-coverage:
	go test -v ./... -cover -coverprofile=coverage.out -covermode=atomic || exit 1

# Formats the coverage report to HTML.
format-coverage:
  go tool cover -html=coverage.out -o coverage.html
