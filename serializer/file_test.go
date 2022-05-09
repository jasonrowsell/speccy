package serializer_test

import (
	"os"
	"testing"

	"github.com/jasonrowsell/speccy/sample"
	"github.com/jasonrowsell/speccy/serializer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriteProtobufToBinaryFile(t *testing.T) {
	t.Run("should write protobuf to file", func(t *testing.T) {
		t.Parallel()
		// given
		filename := "../tmp/laptop.bin"
		laptop := sample.NewLaptop()

		// when
		err := serializer.WriteProtobufToBinaryFile(filename, laptop)

		// then
		require.NoError(t, err)
		assert.FileExists(t, filename)

		// cleanup
		err = os.Remove(filename)
		require.NoError(t, err)
	})
}

func TestReadProtobufFromBinaryFile(t *testing.T) {
	t.Run("should read protobuf from file", func(t *testing.T) {
		t.Parallel()
		// given
		filename := "../tmp/laptop.bin"
		laptop := sample.NewLaptop()

		// when
		err := serializer.WriteProtobufToBinaryFile(filename, laptop)
		require.NoError(t, err)
		err = serializer.ReadProtobufFromBinaryFile(filename, laptop)

		// then
		require.NoError(t, err)

		// cleanup
		err = os.Remove(filename)
		require.NoError(t, err)
	})
}

func TestWriteProtobufToJSONFile(t *testing.T) {
	t.Run("should write protobuf to file", func(t *testing.T) {
		t.Parallel()
		// given
		filename := "../tmp/laptop.json"
		laptop := sample.NewLaptop()

		// when
		err := serializer.WriteProtobufToJSONFile(filename, laptop)

		// then
		require.NoError(t, err)
		assert.FileExists(t, filename)

		// cleanup
		err = os.Remove(filename)
		require.NoError(t, err)
	})
}
