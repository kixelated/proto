package proto

func SizeInt(x int) (n int) {
	return SizeUInt(uint(x))
}

func SizeUInt(x uint) (n int) {
	for x >= 1<<7 {
		n += 1
		x >>= 7
	}

	return n + 1
}

func SizeInt32(x int32) (n int) {
	return SizeUInt32(uint32(x))
}

func SizeUInt32(x uint32) (n int) {
	for x >= 1<<7 {
		n += 1
		x >>= 7
	}

	return n + 1
}

func SizeSInt32(x int32) (n int) {
	// TODO Optimize
	return SizeUInt32(uint32(x<<1) ^ uint32(x>>31))
}

func SizeInt64(x int64) (n int) {
	return SizeUInt64(uint64(x))
}

func SizeUInt64(x uint64) (n int) {
	for x >= 1<<7 {
		n += 1
		x >>= 7
	}

	return n + 1
}

func SizeSInt64(x int64) (n int) {
	// TODO Optimize
	return SizeUInt64(uint64(x<<1) ^ uint64(x>>63))
}

func SizeBool(x bool) (n int) {
	return 1
}

func SizeFixedInt32(x int32) (n int) {
	return 4
}

func SizeFixedInt64(x int64) (n int) {
	return 8
}

func SizeFixedUInt32(x uint32) (n int) {
	return 4
}

func SizeFixedUInt64(x uint64) (n int) {
	return 8
}

func SizeFloat32(x float32) (n int) {
	return 4
}

func SizeFloat64(x float64) (n int) {
	return 8
}

func SizeString(x string) (n int) {
	return len(x)
}

func SizeBytes(x []byte) (n int) {
	return len(x)
}

func SizeTag(id int) (n int) {
	return SizeUInt64(uint64(id<<3))
}
