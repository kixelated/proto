package proto

func (w *Writer) IntRepeated(tagRaw []byte, x []int) {
	for _, y := range x {
		w.data = EncodeBytes(w.data, tagRaw)
		w.data = EncodeInt(w.data, y)
	}
}

func (w *Writer) IntPacked(tagRaw []byte, x []int) {
	if len(x) == 0 {
		return
	}

	size := w.s.Shift()

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeInt(w.data, size)

	for _, y := range x {
		w.data = EncodeInt(w.data, y)
	}
}

func (w *Writer) UIntRepeated(tagRaw []byte, x []uint) {
	for _, y := range x {
		w.data = EncodeBytes(w.data, tagRaw)
		w.data = EncodeUInt(w.data, y)
	}
}

func (w *Writer) UIntPacked(tagRaw []byte, x []uint) {
	if len(x) == 0 {
		return
	}

	size := w.s.Shift()

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeInt(w.data, size)

	for _, y := range x {
		w.data = EncodeUInt(w.data, y)
	}
}

func (w *Writer) Int32Repeated(tagRaw []byte, x []int32) {
	for _, y := range x {
		w.data = EncodeBytes(w.data, tagRaw)
		w.data = EncodeInt32(w.data, y)
	}
}

func (w *Writer) Int32Packed(tagRaw []byte, x []int32) {
	if len(x) == 0 {
		return
	}

	size := w.s.Shift()

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeInt(w.data, size)

	for _, y := range x {
		w.data = EncodeInt32(w.data, y)
	}
}

func (w *Writer) UInt32Repeated(tagRaw []byte, x []uint32) {
	for _, y := range x {
		w.data = EncodeBytes(w.data, tagRaw)
		w.data = EncodeUInt32(w.data, y)
	}
}

func (w *Writer) UInt32Packed(tagRaw []byte, x []uint32) {
	if len(x) == 0 {
		return
	}

	size := w.s.Shift()

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeInt(w.data, size)

	for _, y := range x {
		w.data = EncodeUInt32(w.data, y)
	}
}

func (w *Writer) SInt32Repeated(tagRaw []byte, x []int32) {
	for _, y := range x {
		w.data = EncodeBytes(w.data, tagRaw)
		w.data = EncodeSInt32(w.data, y)
	}
}

func (w *Writer) SInt32Packed(tagRaw []byte, x []int32) {
	if len(x) == 0 {
		return
	}

	size := w.s.Shift()

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeInt(w.data, size)

	for _, y := range x {
		w.data = EncodeSInt32(w.data, y)
	}
}

func (w *Writer) Int64Repeated(tagRaw []byte, x []int64) {
	for _, y := range x {
		w.data = EncodeBytes(w.data, tagRaw)
		w.data = EncodeInt64(w.data, y)
	}
}

func (w *Writer) Int64Packed(tagRaw []byte, x []int64) {
	if len(x) == 0 {
		return
	}

	size := w.s.Shift()

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeInt(w.data, size)

	for _, y := range x {
		w.data = EncodeInt64(w.data, y)
	}
}

func (w *Writer) UInt64Repeated(tagRaw []byte, x []uint64) {
	for _, y := range x {
		w.data = EncodeBytes(w.data, tagRaw)
		w.data = EncodeUInt64(w.data, y)
	}
}

func (w *Writer) UInt64Packed(tagRaw []byte, x []uint64) {
	if len(x) == 0 {
		return
	}

	size := w.s.Shift()

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeInt(w.data, size)

	for _, y := range x {
		w.data = EncodeUInt64(w.data, y)
	}
}

func (w *Writer) SInt64Repeated(tagRaw []byte, x []int64) {
	for _, y := range x {
		w.data = EncodeBytes(w.data, tagRaw)
		w.data = EncodeSInt64(w.data, y)
	}
}

func (w *Writer) SInt64Packed(tagRaw []byte, x []int64) {
	if len(x) == 0 {
		return
	}

	size := w.s.Shift()

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeInt(w.data, size)

	for _, y := range x {
		w.data = EncodeSInt64(w.data, y)
	}
}

func (w *Writer) FixedInt32Repeated(tagRaw []byte, x []int32) {
	for _, y := range x {
		w.data = EncodeBytes(w.data, tagRaw)
		w.data = EncodeFixedInt32(w.data, y)
	}
}

func (w *Writer) FixedInt32Packed(tagRaw []byte, x []int32) {
	if len(x) == 0 {
		return
	}

	size := len(x) * 4

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeInt(w.data, size)

	for _, y := range x {
		w.data = EncodeFixedInt32(w.data, y)
	}
}

func (w *Writer) FixedUInt32Repeated(tagRaw []byte, x []uint32) {
	for _, y := range x {
		w.data = EncodeBytes(w.data, tagRaw)
		w.data = EncodeFixedUInt32(w.data, y)
	}
}

func (w *Writer) FixedUInt32Packed(tagRaw []byte, x []uint32) {
	if len(x) == 0 {
		return
	}

	size := len(x) * 4

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeInt(w.data, size)

	for _, y := range x {
		w.data = EncodeFixedUInt32(w.data, y)
	}
}

func (w *Writer) FixedInt64Repeated(tagRaw []byte, x []int64) {
	for _, y := range x {
		w.data = EncodeBytes(w.data, tagRaw)
		w.data = EncodeFixedInt64(w.data, y)
	}
}

func (w *Writer) FixedInt64Packed(tagRaw []byte, x []int64) {
	if len(x) == 0 {
		return
	}

	size := len(x) * 8

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeInt(w.data, size)

	for _, y := range x {
		w.data = EncodeFixedInt64(w.data, y)
	}
}

func (w *Writer) FixedUInt64Repeated(tagRaw []byte, x []uint64) {
	for _, y := range x {
		w.data = EncodeBytes(w.data, tagRaw)
		w.data = EncodeFixedUInt64(w.data, y)
	}
}

func (w *Writer) FixedUInt64Packed(tagRaw []byte, x []uint64) {
	if len(x) == 0 {
		return
	}

	size := len(x) * 8

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeInt(w.data, size)

	for _, y := range x {
		w.data = EncodeFixedUInt64(w.data, y)
	}
}

func (w *Writer) Float32Repeated(tagRaw []byte, x []float32) {
	for _, y := range x {
		w.data = EncodeBytes(w.data, tagRaw)
		w.data = EncodeFloat32(w.data, y)
	}
}

func (w *Writer) Float32Packed(tagRaw []byte, x []float32) {
	if len(x) == 0 {
		return
	}

	size := len(x) * 4

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeInt(w.data, size)

	for _, y := range x {
		w.data = EncodeFloat32(w.data, y)
	}
}

func (w *Writer) Float64Repeated(tagRaw []byte, x []float64) {
	for _, y := range x {
		w.data = EncodeBytes(w.data, tagRaw)
		w.data = EncodeFloat64(w.data, y)
	}
}

func (w *Writer) Float64Packed(tagRaw []byte, x []float64) {
	if len(x) == 0 {
		return
	}

	size := len(x) * 8

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeInt(w.data, size)

	for _, y := range x {
		w.data = EncodeFloat64(w.data, y)
	}
}

func (w *Writer) BoolRepeated(tagRaw []byte, x []bool) {
	for _, y := range x {
		w.data = EncodeBytes(w.data, tagRaw)
		w.data = EncodeBool(w.data, y)
	}
}

func (w *Writer) BoolPacked(tagRaw []byte, x []bool) {
	if len(x) == 0 {
		return
	}

	size := len(x) * 1

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeInt(w.data, size)

	for _, y := range x {
		w.data = EncodeBool(w.data, y)
	}
}

func (w *Writer) BytesRepeated(tagRaw []byte, x [][]byte) {
	for _, y := range x {
		size := len(y)

		w.data = EncodeBytes(w.data, tagRaw)
		w.data = EncodeInt(w.data, size)
		w.data = EncodeBytes(w.data, y)
	}
}

func (w *Writer) StringRepeated(tagRaw []byte, x []string) {
	for _, y := range x {
		size := len(y)

		w.data = EncodeBytes(w.data, tagRaw)
		w.data = EncodeInt(w.data, size)
		w.data = EncodeString(w.data, y)
	}
}

func (w *Writer) MessageRepeated(tagRaw []byte, x []MarshallerTo) {
	for _, y := range x {
		size := w.s.Shift()

		w.data = EncodeBytes(w.data, tagRaw)
		w.data = EncodeInt(w.data, size)
		y.MarshalTo(w)
	}
}
