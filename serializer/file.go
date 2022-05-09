package serializer

import (
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/proto"
)

func WriteProtobufToBinaryFile(filename string, message proto.Message) error {
	// TODO: add tests for mock serializer
	data, err := proto.Marshal(message)

	if err != nil {
		log.Fatal("Failed to marshal proto message: ", err)
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)

	if err != nil {
		// TODO: improve error handling
		log.Fatal("Failed to write proto to file: ", err)
		return err
	}

	return nil
}

func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal("Failed to read proto from file: ", err)
		return err
	}

	err = proto.Unmarshal(data, message)

	if err != nil {
		log.Fatal("Failed to unmarshal proto message: ", err)
		return err
	}

	return nil
}

func WriteProtobufToJSONFile(filename string, message proto.Message) error {
	data, err := ProtobufToJSON(message)

	if err != nil {
		log.Fatal("Failed to marshal proto message to JSON: ", err)
		return err
	}

	err = ioutil.WriteFile(filename, data, 0644)

	if err != nil {
		log.Fatal("Failed to write JSON to file: ", err)
		return err
	}

	return nil
}
