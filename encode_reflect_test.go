package proto_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kixelated/proto"
)

func TestEncode(t *testing.T) {
	assert := assert.New(t)

	out, err := proto.Encode(int(300))
	if assert.NoError(err) {
		assert.Equal(proto.EncodeSInt(nil, 300), out)
	}

	out, err = proto.Encode(uint(300))
	if assert.NoError(err) {
		assert.Equal(proto.EncodeUInt(nil, 300), out)
	}
}

func TestEncodeStruct(t *testing.T) {
	assert := assert.New(t)

	type Test1 struct {
		A uint32 `proto:"1"`
	}

	in1 := Test1{A: 150}
	out, err := proto.Encode(in1)
	if assert.NoError(err) {
		assert.Equal([]byte{0x08, 0x96, 0x01}, out)
	}

	type Test2 struct {
		B string `proto:"2"`
	}

	in2 := Test2{B: "testing"}
	out, err = proto.Encode(in2)
	if assert.NoError(err) {
		assert.Equal([]byte{0x12, 0x07, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67}, out)
	}

	type Test3 struct {
		C Test1 `proto:"3"`
	}

	in3 := Test3{C: Test1{A: 150}}
	out, err = proto.Encode(in3)
	if assert.NoError(err) {
		assert.Equal([]byte{0x1a, 0x03, 0x08, 0x96, 0x01}, out)
	}
}
