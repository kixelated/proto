package proto

func (s *Sizer) Int(x int) {
	s.n += SizeInt(x)
}

func (s *Sizer) UInt(x uint) {
	s.n += SizeUInt(x)
}

func (s *Sizer) Int32(x int32) {
	s.n += SizeInt32(x)
}

func (s *Sizer) UInt32(x uint32) {
	s.n += SizeUInt32(x)
}

func (s *Sizer) SInt32(x int32) {
	s.n += SizeSInt32(x)
}

func (s *Sizer) Int64(x int64) {
	s.n += SizeInt64(x)
}

func (s *Sizer) UInt64(x uint64) {
	s.n += SizeUInt64(x)
}

func (s *Sizer) SInt64(x int64) {
	s.n += SizeSInt64(x)
}

func (s *Sizer) Bool(x bool) {
	s.n += SizeBool(x)
}

func (s *Sizer) Tag(id int) {
	s.n += SizeUInt64(uint64(id) << 3)
}

func (s *Sizer) Float32(x float32) {
	s.n += SizeFloat32(x)
}

func (s *Sizer) Float64(x float64) {
	s.n += SizeFloat64(x)
}

func (s *Sizer) FixedInt32(x int32) {
	s.n += SizeFixedInt32(x)
}

func (s *Sizer) FixedUInt32(x uint32) {
	s.n += SizeFixedUInt32(x)
}

func (s *Sizer) FixedInt64(x int64) {
	s.n += SizeFixedInt64(x)
}

func (s *Sizer) FixedUInt64(x uint64) {
	s.n += SizeFixedUInt64(x)
}

func (s *Sizer) String(x string) {
	s.n += SizeString(x)
}

func (s *Sizer) Message(x MarshallerSize) {
	x.MarshalSize(s)
}

func (s *Sizer) Bytes(x []byte) {
	s.n += SizeBytes(x)
}
