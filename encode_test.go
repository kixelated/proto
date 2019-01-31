package proto_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kixelated/proto"
)

func TestUInt32(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		Input    uint32
		Expected []byte
	}{
		{
			Input:    0,
			Expected: []byte{0},
		},
		{
			Input:    1,
			Expected: []byte{1},
		},
		{
			Input:    2,
			Expected: []byte{2},
		},
		{
			Input:    127,
			Expected: []byte{0x7f},
		},
		{
			Input:    128,
			Expected: []byte{0x80, 0x1},
		},
		{
			Input:    129,
			Expected: []byte{0x81, 0x1},
		},
		{
			Input:    270,
			Expected: []byte{0x8e, 0x2},
		},
		{
			Input:    86942,
			Expected: []byte{0x9e, 0xa7, 0x05},
		},
	}

	for i, test := range tests {
		output := proto.EncodeUInt32(nil, test.Input)
		assert.Equal(test.Expected, output, "%d: wrong output", i)

		size := proto.SizeUInt32(test.Input)
		assert.Equal(len(test.Expected), size, "%d: wrong size", i)
	}
}

func TestInt32(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		Input    int32
		Expected []byte
	}{
		{
			Input:    0,
			Expected: []byte{0},
		},
		{
			Input:    1,
			Expected: []byte{1},
		},
		{
			Input:    -1,
			Expected: []byte{0xff, 0xff, 0xff, 0xff, 0xf},
		},
		{
			Input:    2,
			Expected: []byte{2},
		},
		{
			Input:    -2,
			Expected: []byte{0xfe, 0xff, 0xff, 0xff, 0xf},
		},
		{
			Input:    127,
			Expected: []byte{0x7f},
		},
		{
			Input:    -127,
			Expected: []byte{0x81, 0xff, 0xff, 0xff, 0x0f},
		},
		{
			Input:    128,
			Expected: []byte{0x80, 0x1},
		},
		{
			Input:    -128,
			Expected: []byte{0x80, 0xff, 0xff, 0xff, 0x0f},
		},
		{
			Input:    129,
			Expected: []byte{0x81, 0x1},
		},
		{
			Input:    -129,
			Expected: []byte{0xff, 0xfe, 0xff, 0xff, 0x0f},
		},
	}

	for i, test := range tests {
		output := proto.EncodeInt32(nil, test.Input)
		assert.Equal(test.Expected, output, "%d: wrong output", i)

		size := proto.SizeInt32(test.Input)
		assert.Equal(len(test.Expected), size, "%d: wrong size", i)
	}
}

func TestSInt32(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		Input    int32
		Expected []byte
	}{
		{
			Input:    0,
			Expected: []byte{0},
		},
		{
			Input:    -1,
			Expected: []byte{1},
		},
		{
			Input:    1,
			Expected: []byte{2},
		},
		{
			Input:    -2,
			Expected: []byte{3},
		},
		{
			Input:    2,
			Expected: []byte{4},
		},
		{
			Input:    63,
			Expected: []byte{0x7e},
		},
		{
			Input:    -64,
			Expected: []byte{0x7f},
		},
		{
			Input:    64,
			Expected: []byte{0x80, 0x1},
		},
		{
			Input:    -65,
			Expected: []byte{0x81, 0x1},
		},
	}

	for i, test := range tests {
		output := proto.EncodeSInt32(nil, test.Input)
		assert.Equal(test.Expected, output, "%d: wrong output", i)

		size := proto.SizeSInt32(test.Input)
		assert.Equal(len(test.Expected), size, "%d: wrong size", i)
	}
}

func TestUInt64(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		Input    uint64
		Expected []byte
	}{
		{
			Input:    0,
			Expected: []byte{0},
		},
		{
			Input:    1,
			Expected: []byte{1},
		},
		{
			Input:    2,
			Expected: []byte{2},
		},
		{
			Input:    127,
			Expected: []byte{0x7f},
		},
		{
			Input:    128,
			Expected: []byte{0x80, 0x1},
		},
		{
			Input:    129,
			Expected: []byte{0x81, 0x1},
		},
		{
			Input:    270,
			Expected: []byte{0x8e, 0x2},
		},
		{
			Input:    86942,
			Expected: []byte{0x9e, 0xa7, 0x05},
		},
	}

	for i, test := range tests {
		output := proto.EncodeUInt64(nil, test.Input)
		assert.Equal(test.Expected, output, "%d: wrong output", i)

		size := proto.SizeUInt64(test.Input)
		assert.Equal(len(test.Expected), size, "%d: wrong size", i)
	}
}

func TestInt64(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		Input    int64
		Expected []byte
	}{
		{
			Input:    0,
			Expected: []byte{0},
		},
		{
			Input:    1,
			Expected: []byte{1},
		},
		{
			Input:    -1,
			Expected: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1},
		},
		{
			Input:    2,
			Expected: []byte{2},
		},
		{
			Input:    -2,
			Expected: []byte{0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1},
		},
		{
			Input:    127,
			Expected: []byte{0x7f},
		},
		{
			Input:    -127,
			Expected: []byte{0x81, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1},
		},
		{
			Input:    128,
			Expected: []byte{0x80, 0x1},
		},
		{
			Input:    -128,
			Expected: []byte{0x80, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1},
		},
		{
			Input:    129,
			Expected: []byte{0x81, 0x1},
		},
		{
			Input:    -129,
			Expected: []byte{0xff, 0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1},
		},
	}

	for i, test := range tests {
		output := proto.EncodeInt64(nil, test.Input)
		assert.Equal(test.Expected, output, "%d: wrong output", i)

		size := proto.SizeInt64(test.Input)
		assert.Equal(len(test.Expected), size, "%d: wrong size", i)
	}
}

func TestSInt64(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		Input    int64
		Expected []byte
	}{
		{
			Input:    0,
			Expected: []byte{0},
		},
		{
			Input:    -1,
			Expected: []byte{1},
		},
		{
			Input:    1,
			Expected: []byte{2},
		},
		{
			Input:    -2,
			Expected: []byte{3},
		},
		{
			Input:    2,
			Expected: []byte{4},
		},
		{
			Input:    63,
			Expected: []byte{0x7e},
		},
		{
			Input:    -64,
			Expected: []byte{0x7f},
		},
		{
			Input:    64,
			Expected: []byte{0x80, 0x1},
		},
		{
			Input:    -65,
			Expected: []byte{0x81, 0x1},
		},
	}

	for i, test := range tests {
		output := proto.EncodeSInt64(nil, test.Input)
		assert.Equal(test.Expected, output, "%d: wrong output", i)

		size := proto.SizeSInt64(test.Input)
		assert.Equal(len(test.Expected), size, "%d: wrong size", i)
	}
}

func TestUInt(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		Input    uint
		Expected []byte
	}{
		{
			Input:    0,
			Expected: []byte{0},
		},
		{
			Input:    1,
			Expected: []byte{1},
		},
		{
			Input:    2,
			Expected: []byte{2},
		},
		{
			Input:    127,
			Expected: []byte{0x7f},
		},
		{
			Input:    128,
			Expected: []byte{0x80, 0x1},
		},
		{
			Input:    129,
			Expected: []byte{0x81, 0x1},
		},
		{
			Input:    270,
			Expected: []byte{0x8e, 0x2},
		},
		{
			Input:    86942,
			Expected: []byte{0x9e, 0xa7, 0x05},
		},
	}

	for i, test := range tests {
		output := proto.EncodeUInt(nil, test.Input)
		assert.Equal(test.Expected, output, "%d: wrong output", i)

		size := proto.SizeUInt(test.Input)
		assert.Equal(len(test.Expected), size, "%d: wrong size", i)
	}
}

func TestInt(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		Input    int
		Expected []byte
	}{
		{
			Input:    0,
			Expected: []byte{0},
		},
		{
			Input:    1,
			Expected: []byte{1},
		},
		{
			Input:    -1,
			Expected: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1}, // 32-bit in 2017 LUL
		},
		{
			Input:    2,
			Expected: []byte{2},
		},
		{
			Input:    -2,
			Expected: []byte{0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1},
		},
		{
			Input:    127,
			Expected: []byte{0x7f},
		},
		{
			Input:    -127,
			Expected: []byte{0x81, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1},
		},
		{
			Input:    128,
			Expected: []byte{0x80, 0x1},
		},
		{
			Input:    -128,
			Expected: []byte{0x80, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1},
		},
		{
			Input:    129,
			Expected: []byte{0x81, 0x1},
		},
		{
			Input:    -129,
			Expected: []byte{0xff, 0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1},
		},
	}

	for i, test := range tests {
		output := proto.EncodeInt(nil, test.Input)
		assert.Equal(test.Expected, output, "%d: wrong output", i)

		size := proto.SizeInt(test.Input)
		assert.Equal(len(test.Expected), size, "%d: wrong size", i)
	}
}

func TestFixedInt32(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		Input    int32
		Expected []byte
	}{
		{
			Input:    0,
			Expected: []byte{0, 0, 0, 0},
		},
		{
			Input:    1,
			Expected: []byte{1, 0, 0, 0},
		},
		{
			Input:    -1,
			Expected: []byte{0xff, 0xff, 0xff, 0xff},
		},
		{
			Input:    2,
			Expected: []byte{2, 0, 0, 0},
		},
		{
			Input:    -2,
			Expected: []byte{0xfe, 0xff, 0xff, 0xff},
		},
		{
			Input:    256,
			Expected: []byte{0, 1, 0, 0},
		},
		{
			Input:    -256,
			Expected: []byte{0x0, 0xff, 0xff, 0xff},
		},
	}

	for i, test := range tests {
		output := proto.EncodeFixedInt32(nil, test.Input)
		assert.Equal(test.Expected, output, "%d: wrong output", i)

		size := proto.SizeFixedInt32(test.Input)
		assert.Equal(len(test.Expected), size, "%d: wrong size", i)
	}
}

func TestFixedInt64(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		Input    int64
		Expected []byte
	}{
		{
			Input:    0,
			Expected: []byte{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Input:    1,
			Expected: []byte{1, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Input:    -1,
			Expected: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		},
		{
			Input:    2,
			Expected: []byte{2, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Input:    -2,
			Expected: []byte{0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		},
		{
			Input:    256,
			Expected: []byte{0, 1, 0, 0, 0, 0, 0, 0},
		},
		{
			Input:    -256,
			Expected: []byte{0x0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		},
	}

	for i, test := range tests {
		output := proto.EncodeFixedInt64(nil, test.Input)
		assert.Equal(test.Expected, output, "%d: wrong output", i)

		size := proto.SizeFixedInt64(test.Input)
		assert.Equal(len(test.Expected), size, "%d: wrong size", i)
	}
}

func BenchmarkIntBest(b *testing.B) {
	data := make([]byte, 0, 10)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		proto.EncodeInt(data[0:], 1)
	}
}

func BenchmarkIntWorst(b *testing.B) {
	data := make([]byte, 0, 10)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		proto.EncodeInt(data[0:], ^int(0))
	}
}

func BenchmarkUIntBest(b *testing.B) {
	data := make([]byte, 0, 10)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		proto.EncodeUInt(data[0:], 1)
	}
}

func BenchmarkUIntWorst(b *testing.B) {
	data := make([]byte, 0, 10)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		proto.EncodeUInt(data[0:], ^uint(0))
	}
}

func BenchmarkInt32Best(b *testing.B) {
	data := make([]byte, 0, 5)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		proto.EncodeInt32(data[0:], 1)
	}
}

func BenchmarkInt32Worst(b *testing.B) {
	data := make([]byte, 0, 5)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		proto.EncodeInt32(data[0:], -1<<31)
	}
}

func BenchmarkUInt32Best(b *testing.B) {
	data := make([]byte, 0, 5)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		proto.EncodeUInt32(data[0:], 1)
	}
}

func BenchmarkUInt32Worst(b *testing.B) {
	data := make([]byte, 0, 5)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		proto.EncodeUInt32(data[0:], 1<<32-1)
	}
}

func BenchmarkInt64Best(b *testing.B) {
	data := make([]byte, 0, 10)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		proto.EncodeInt64(data[0:], 1)
	}
}

func BenchmarkInt64Worst(b *testing.B) {
	data := make([]byte, 0, 10)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		proto.EncodeInt64(data[0:], -1<<63)
	}
}

func BenchmarkUInt64Best(b *testing.B) {
	data := make([]byte, 0, 10)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		proto.EncodeUInt64(data[0:], 1)
	}
}

func BenchmarkUInt64Worst(b *testing.B) {
	data := make([]byte, 0, 10)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		proto.EncodeUInt64(data[0:], 1<<64-1)
	}
}
