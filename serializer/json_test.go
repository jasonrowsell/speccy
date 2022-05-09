package serializer_test

import (
	"testing"

	"github.com/jasonrowsell/speccy/sample"
	"github.com/jasonrowsell/speccy/serializer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProtobufToJSON(t *testing.T) {
	t.Run("should marshal protobuf to JSON", func(t *testing.T) {
		t.Parallel()
		// given
		laptop := sample.NewLaptop()

		// when
		data, err := serializer.ProtobufToJSON(laptop)

		// then
		require.NoError(t, err)
		assert.NotEmpty(t, data)
	})
}

func TestJSONToProtobuf(t *testing.T) {
	t.Run("should unmarshal JSON to protobuf", func(t *testing.T) {
		t.Parallel()
		// given
		laptop := sample.NewLaptop()
		data, err := serializer.ProtobufToJSON(laptop)
		require.NoError(t, err)

		// when
		err = serializer.JSONToProtobuf(data, laptop)

		// then
		require.NoError(t, err)
	})
}
