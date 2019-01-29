package proto

func (w *Writer) Tag(id int, t WireType) {
	w.data = EncodeUInt64(w.data, uint64(id)<<3|uint64(t))
}

func (w *Writer) IntField(tagRaw []byte, x int) {
	if x == 0 {
		return
	}

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeInt(w.data, x)
}

func (w *Writer) UIntField(tagRaw []byte, x uint) {
	if x == 0 {
		return
	}

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeUInt(w.data, x)
}

func (w *Writer) Int32Field(tagRaw []byte, x int32) {
	if x == 0 {
		return
	}

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeInt32(w.data, x)
}

func (w *Writer) SInt32Field(tagRaw []byte, x int32) {
	if x == 0 {
		return
	}

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeSInt32(w.data, x)
}

func (w *Writer) UInt32Field(tagRaw []byte, x uint32) {
	if x == 0 {
		return
	}

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeUInt32(w.data, x)
}

func (w *Writer) Int64Field(tagRaw []byte, x int64) {
	if x == 0 {
		return
	}

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeInt64(w.data, x)
}

func (w *Writer) SInt64Field(tagRaw []byte, x int64) {
	if x == 0 {
		return
	}

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeSInt64(w.data, x)
}

func (w *Writer) UInt64Field(tagRaw []byte, x uint64) {
	if x == 0 {
		return
	}

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeUInt64(w.data, x)
}

func (w *Writer) FixedInt32Field(tagRaw []byte, x int32) {
	if x == 0 {
		return
	}

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeFixedInt32(w.data, x)
}

func (w *Writer) FixedUInt32Field(tagRaw []byte, x uint32) {
	if x == 0 {
		return
	}

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeFixedUInt32(w.data, x)
}

func (w *Writer) FixedInt64Field(tagRaw []byte, x int64) {
	if x == 0 {
		return
	}

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeFixedInt64(w.data, x)
}

func (w *Writer) FixedUInt64Field(tagRaw []byte, x uint64) {
	if x == 0 {
		return
	}

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeFixedUInt64(w.data, x)
}

func (w *Writer) Float32Field(tagRaw []byte, x float32) {
	if x == 0 {
		return
	}

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeFloat32(w.data, x)
}

func (w *Writer) Float64Field(tagRaw []byte, x float64) {
	if x == 0 {
		return
	}

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeFloat64(w.data, x)
}

func (w *Writer) BoolField(tagRaw []byte, x bool) {
	if !x {
		return
	}

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeInt(w.data, 1)
}

func (w *Writer) StringField(tagRaw []byte, x string) {
	size := len(x)
	if size == 0 {
		return
	}

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeInt(w.data, size)
	w.data = EncodeString(w.data, x)
}

func (w *Writer) MessageField(tagRaw []byte, x MarshallerTo) {
	size := w.s.Shift()
	if size == 0 {
		return
	}

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeInt(w.data, size)
	x.MarshalTo(w)
}

func (w *Writer) BytesField(tagRaw []byte, x []byte) {
	size := len(x)
	if size == 0 {
		return
	}

	w.data = EncodeBytes(w.data, tagRaw)
	w.data = EncodeInt(w.data, size)
	w.data = EncodeBytes(w.data, x)
}
