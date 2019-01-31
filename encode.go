package proto

import (
	"math"
)

func EncodeInt(data []byte, x int) []byte {
	// TODO Optimize
	return EncodeUInt(data, uint(x))
}

func EncodeUInt(data []byte, x uint) []byte {
	// what have i done
	if x < (1 << 14) {
		if x < (1 << 7) {
			return append(data, byte(x))
		} else { // if x < (1 << 14) {
			return append(data, byte(x)|0x80, byte(x>>7))
		}
	} else if x < (1 << 42) {
		if x < (1 << 28) {
			if x < (1 << 21) {
				return append(data, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14))
			} else { // if x < (1 << 28) {
				return append(data, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14)|0x80, byte(x>>21))
			}
		} else if x < (1 << 35) {
			return append(data, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14)|0x80, byte(x>>21)|0x80, byte(x>>28))
		} else { // if x < (1 << 42) {
			return append(data, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14)|0x80, byte(x>>21)|0x80, byte(x>>28)|0x80, byte(x>>35))
		}
	} else if x < (1 << 56) {
		if x < (1 << 49) {
			return append(data, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14)|0x80, byte(x>>21)|0x80, byte(x>>28)|0x80, byte(x>>35)|0x80, byte(x>>42))
		} else { // if x < (1 << 56) {
			return append(data, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14)|0x80, byte(x>>21)|0x80, byte(x>>28)|0x80, byte(x>>35)|0x80, byte(x>>42)|0x80, byte(x>>49))
		}
	} else if x < (1 << 63) {
		return append(data, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14)|0x80, byte(x>>21)|0x80, byte(x>>28)|0x80, byte(x>>35)|0x80, byte(x>>42)|0x80, byte(x>>49)|0x80, byte(x>>56))
	} else { // x < 1 << 70
		return append(data, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14)|0x80, byte(x>>21)|0x80, byte(x>>28)|0x80, byte(x>>35)|0x80, byte(x>>42)|0x80, byte(x>>49)|0x80, byte(x>>56)|0x80, byte(x>>63))
	}
}

func EncodeSInt(data []byte, x int) []byte {
	// TODO Optimize
	return EncodeSInt64(data, int64(x))
}

func EncodeInt32(data []byte, x int32) []byte {
	// TODO Optimize
	return EncodeUInt32(data, uint32(x))
}

func EncodeUInt32(data []byte, x uint32) []byte {
	if x < (1 << 7) {
		return append(data, byte(x))
	} else if x < (1 << 14) {
		return append(data, byte(x)|0x80, byte(x>>7))
	} else if x < (1 << 21) {
		return append(data, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14))
	} else if x < (1 << 28) {
		return append(data, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14)|0x80, byte(x>>21))
	} else {
		return append(data, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14)|0x80, byte(x>>21)|0x80, byte(x>>28))
	}
}

func EncodeSInt32(data []byte, x int32) []byte {
	// TODO Optimize
	return EncodeUInt32(data, uint32((x<<1)^(x>>31)))
}

func EncodeInt64(data []byte, x int64) []byte {
	// TODO Optimize
	return EncodeUInt64(data, uint64(x))
}

func EncodeUInt64(data []byte, x uint64) []byte {
	// what have i done
	if x < (1 << 14) {
		if x < (1 << 7) {
			return append(data, byte(x))
		} else { // if x < (1 << 14) {
			return append(data, byte(x)|0x80, byte(x>>7))
		}
	} else if x < (1 << 42) {
		if x < (1 << 28) {
			if x < (1 << 21) {
				return append(data, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14))
			} else { // if x < (1 << 28) {
				return append(data, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14)|0x80, byte(x>>21))
			}
		} else if x < (1 << 35) {
			return append(data, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14)|0x80, byte(x>>21)|0x80, byte(x>>28))
		} else { // if x < (1 << 42) {
			return append(data, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14)|0x80, byte(x>>21)|0x80, byte(x>>28)|0x80, byte(x>>35))
		}
	} else if x < (1 << 56) {
		if x < (1 << 49) {
			return append(data, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14)|0x80, byte(x>>21)|0x80, byte(x>>28)|0x80, byte(x>>35)|0x80, byte(x>>42))
		} else { // if x < (1 << 56) {
			return append(data, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14)|0x80, byte(x>>21)|0x80, byte(x>>28)|0x80, byte(x>>35)|0x80, byte(x>>42)|0x80, byte(x>>49))
		}
	} else if x < (1 << 63) {
		return append(data, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14)|0x80, byte(x>>21)|0x80, byte(x>>28)|0x80, byte(x>>35)|0x80, byte(x>>42)|0x80, byte(x>>49)|0x80, byte(x>>56))
	} else { // x < 1 << 70
		return append(data, byte(x)|0x80, byte(x>>7)|0x80, byte(x>>14)|0x80, byte(x>>21)|0x80, byte(x>>28)|0x80, byte(x>>35)|0x80, byte(x>>42)|0x80, byte(x>>49)|0x80, byte(x>>56)|0x80, byte(x>>63))
	}
}

func EncodeSInt64(data []byte, x int64) []byte {
	// TODO Optimize
	return EncodeUInt64(data, uint64((x<<1)^(x>>63)))
}

func EncodeBool(data []byte, x bool) []byte {
	if x {
		return append(data, 1)
	} else {
		return append(data, 0)
	}
}

func EncodeFixedInt32(data []byte, x int32) []byte {
	// TODO Optimize
	return EncodeFixedUInt32(data, uint32(x))
}

func EncodeFixedUInt32(data []byte, x uint32) []byte {
	return append(data, byte(x), byte(x>>8), byte(x>>16), byte(x>>24))
}

func EncodeFixedInt64(data []byte, x int64) []byte {
	// TODO Optimize
	return EncodeFixedUInt64(data, uint64(x))
}

func EncodeFixedUInt64(data []byte, x uint64) []byte {
	return append(data, byte(x), byte(x>>8), byte(x>>16), byte(x>>24), byte(x>>32), byte(x>>40), byte(x>>48), byte(x>>56))
}

func EncodeFloat64(data []byte, x float64) []byte {
	return EncodeFixedUInt64(data, math.Float64bits(x))
}

func EncodeFloat32(data []byte, x float32) []byte {
	return EncodeFixedUInt32(data, math.Float32bits(x))
}

func grow(data []byte, size int) []byte {
	if size > cap(data) {
		temp := make([]byte, len(data), 2*size)
		copy(temp, data)
		data = temp
	}

	return data[:size]
}

func EncodeString(data []byte, x string) []byte {
	data = grow(data, len(data)+len(x))
	copy(data[len(data)-len(x):], x)
	return data
}

func EncodeBytes(data []byte, x []byte) []byte {
	data = grow(data, len(data)+len(x))
	copy(data[len(data)-len(x):], x)
	return data
}

func EncodeTag(data []byte, id int, t WireType) []byte {
	// TODO Optimize maybe
	return EncodeUInt64(data, (uint64(id)<<3)|uint64(t))
}
