package proto_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kixelated/proto"
)

func TestEncode(t *testing.T) {
	assert := assert.New(t)

	out, err := proto.Encode(nil, int(300))
	if assert.NoError(err) {
		assert.Equal(proto.EncodeSInt(nil, 300), out)
	}

	out, err = proto.Encode(nil, uint(300))
	if assert.NoError(err) {
		assert.Equal(proto.EncodeUInt(nil, 300), out)
	}
}
