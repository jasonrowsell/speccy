gen:
	protoc --proto_path=protobuf protobuf/*.proto --go_out=plugins=grpc:pb

clean:
	rm pb/*.go
