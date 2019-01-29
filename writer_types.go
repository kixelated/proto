package proto

func (w *Writer) Int(x int) {
	w.data = EncodeInt(w.data, x)
}

func (w *Writer) UInt(x uint) {
	w.data = EncodeUInt(w.data, x)
}

func (w *Writer) Int32(x int32) {
	w.data = EncodeInt32(w.data, x)
}

func (w *Writer) SInt32(x int32) {
	w.data = EncodeSInt32(w.data, x)
}

func (w *Writer) UInt32(x uint32) {
	w.data = EncodeUInt32(w.data, x)
}

func (w *Writer) Int64(x int64) {
	w.data = EncodeInt64(w.data, x)
}

func (w *Writer) SInt64(x int64) {
	w.data = EncodeSInt64(w.data, x)
}

func (w *Writer) UInt64(x uint64) {
	w.data = EncodeUInt64(w.data, x)
}

func (w *Writer) FixedInt32(x int32) {
	w.data = EncodeFixedInt32(w.data, x)
}

func (w *Writer) FixedUInt32(x uint32) {
	w.data = EncodeFixedUInt32(w.data, x)
}

func (w *Writer) FixedInt64(x int64) {
	w.data = EncodeFixedInt64(w.data, x)
}

func (w *Writer) FixedUInt64(x uint64) {
	w.data = EncodeFixedUInt64(w.data, x)
}

func (w *Writer) Float32(x float32) {
	w.data = EncodeFloat32(w.data, x)
}

func (w *Writer) Float64(x float64) {
	w.data = EncodeFloat64(w.data, x)
}

func (w *Writer) Bool(x bool) {
	w.data = EncodeBool(w.data, x)
}

func (w *Writer) String(x string) {
	w.data = EncodeString(w.data, x)
}

func (w *Writer) Byte(x byte) {
	w.data = append(w.data, x)
}

func (w *Writer) Bytes(x []byte) {
	w.data = EncodeBytes(w.data, x)
}

func (w *Writer) Message(x MarshallerTo) {
	x.MarshalTo(w)
}
