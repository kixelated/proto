package proto

func (s *Sizer) IntRepeated(tag int, x []int) {
	for _, y := range x {
		s.TagSize(tag)
		s.Int(y)
	}
}

func (s *Sizer) IntPacked(tag int, x []int) {
	if len(x) == 0 {
		return
	}

	s.TagSize(tag)

	start := s.Size()
	for _, y := range x {
		s.Int(y)
	}

	size := s.Size() - start

	s.Int(size)
	s.Push(size)
}

func (s *Sizer) UIntRepeated(tag int, x []uint) {
	for _, y := range x {
		s.TagSize(tag)
		s.UInt(y)
	}
}

func (s *Sizer) UIntPacked(tag int, x []uint) {
	if len(x) == 0 {
		return
	}

	s.TagSize(tag)

	start := s.Size()
	for _, y := range x {
		s.UInt(y)
	}

	size := s.Size() - start

	s.Int(size)
	s.Push(size)
}

func (s *Sizer) Int32Repeated(tag int, x []int32) {
	for _, y := range x {
		s.TagSize(tag)
		s.Int32(y)
	}
}

func (s *Sizer) Int32Packed(tag int, x []int32) {
	if len(x) == 0 {
		return
	}

	s.TagSize(tag)

	start := s.Size()
	for _, y := range x {
		s.Int32(y)
	}

	size := s.Size() - start

	s.Int(size)
	s.Push(size)
}

func (s *Sizer) SInt32Repeated(tag int, x []int32) {
	for _, y := range x {
		s.TagSize(tag)
		s.SInt32(y)
	}
}

func (s *Sizer) SInt32Packed(tag int, x []int32) {
	if len(x) == 0 {
		return
	}

	s.TagSize(tag)

	start := s.Size()
	for _, y := range x {
		s.SInt32(y)
	}

	size := s.Size() - start

	s.Int(size)
	s.Push(size)
}

func (s *Sizer) UInt32Repeated(tag int, x []uint32) {
	for _, y := range x {
		s.TagSize(tag)
		s.UInt32(y)
	}
}

func (s *Sizer) UInt32Packed(tag int, x []uint32) {
	if len(x) == 0 {
		return
	}

	s.TagSize(tag)

	start := s.Size()
	for _, y := range x {
		s.UInt32(y)
	}

	size := s.Size() - start

	s.Int(size)
	s.Push(size)
}

func (s *Sizer) Int64Repeated(tag int, x []int64) {
	for _, y := range x {
		s.TagSize(tag)
		s.Int64(y)
	}
}

func (s *Sizer) Int64Packed(tag int, x []int64) {
	if len(x) == 0 {
		return
	}

	s.TagSize(tag)

	start := s.Size()
	for _, y := range x {
		s.Int64(y)
	}

	size := s.Size() - start

	s.Int(size)
	s.Push(size)
}

func (s *Sizer) SInt64Repeated(tag int, x []int64) {
	for _, y := range x {
		s.TagSize(tag)
		s.SInt64(y)
	}
}

func (s *Sizer) SInt64Packed(tag int, x []int64) {
	if len(x) == 0 {
		return
	}

	s.TagSize(tag)

	start := s.Size()
	for _, y := range x {
		s.SInt64(y)
	}

	size := s.Size() - start

	s.Int(size)
	s.Push(size)
}

func (s *Sizer) UInt64Repeated(tag int, x []uint64) {
	for _, y := range x {
		s.TagSize(tag)
		s.UInt64(y)
	}
}

func (s *Sizer) UInt64Packed(tag int, x []uint64) {
	if len(x) == 0 {
		return
	}

	s.TagSize(tag)

	start := s.Size()
	for _, y := range x {
		s.UInt64(y)
	}

	size := s.Size() - start

	s.Int(size)
	s.Push(size)
}

func (s *Sizer) BoolRepeated(tag int, x []bool) {
	for _, y := range x {
		s.TagSize(tag)
		s.Bool(y)
	}
}

func (s *Sizer) BoolPacked(tag int, x []bool) {
	if len(x) == 0 {
		return
	}

	s.TagSize(tag)

	size := len(x)

	s.Int(size)
	s.n += size
}

func (s *Sizer) Float32Repeated(tag int, x []float32) {
	for _, y := range x {
		s.TagSize(tag)
		s.Float32(y)
	}
}

func (s *Sizer) Float32Packed(tag int, x []float32) {
	if len(x) == 0 {
		return
	}

	s.TagSize(tag)

	size := len(x) * 4

	s.Int(size)
	s.n += size
}

func (s *Sizer) Float64Repeated(tag int, x []float64) {
	for _, y := range x {
		s.TagSize(tag)
		s.Float64(y)
	}
}

func (s *Sizer) Float64Packed(tag int, x []float64) {
	if len(x) == 0 {
		return
	}

	s.TagSize(tag)

	size := len(x) * 8

	s.Int(size)
	s.n += size
}

func (s *Sizer) FixedInt32Repeated(tag int, x []int32) {
	for _, y := range x {
		s.TagSize(tag)
		s.FixedInt32(y)
	}
}

func (s *Sizer) FixedInt32Packed(tag int, x []int32) {
	if len(x) == 0 {
		return
	}

	s.TagSize(tag)

	size := len(x) * 4
	s.Int(size)
	s.n += size
}

func (s *Sizer) FixedUInt32Repeated(tag int, x []uint32) {
	for _, y := range x {
		s.TagSize(tag)
		s.FixedUInt32(y)
	}
}

func (s *Sizer) FixedUInt32Packed(tag int, x []uint32) {
	if len(x) == 0 {
		return
	}

	s.TagSize(tag)

	size := len(x) * 4
	s.Int(size)
	s.n += size
}

func (s *Sizer) FixedInt64Repeated(tag int, x []int64) {
	for _, y := range x {
		s.TagSize(tag)
		s.FixedInt64(y)
	}
}

func (s *Sizer) FixedInt64Packed(tag int, x []int64) {
	if len(x) == 0 {
		return
	}

	s.TagSize(tag)

	size := len(x) * 8
	s.Int(size)
	s.n += size
}

func (s *Sizer) FixedUInt64Repeated(tag int, x []uint64) {
	for _, y := range x {
		s.TagSize(tag)
		s.FixedUInt64(y)
	}
}

func (s *Sizer) FixedUInt64Packed(tag int, x []uint64) {
	if len(x) == 0 {
		return
	}

	s.TagSize(tag)

	size := len(x) * 8
	s.Int(size)
	s.n += size
}

func (s *Sizer) StringRepeated(tag int, x []string) {
	if len(x) == 0 {
		return
	}

	for _, y := range x {
		s.TagSize(tag)

		size := SizeString(y)
		s.Int(size)
		s.n += size
	}
}

func (s *Sizer) MessageRepeated(tag int, x []MarshallerSize) {
	if len(x) == 0 {
		return
	}

	for _, y := range x {
		s.TagSize(tag)

		start := s.Size()
		y.MarshalSize(s)
		size := s.Size() - start

		s.Push(size)
		s.Int(size)
	}
}

func (s *Sizer) BytesRepeated(tag int, x [][]byte) {
	if len(x) == 0 {
		return
	}

	for _, y := range x {
		s.TagSize(tag)

		size := SizeBytes(y)
		s.Int(size)
		s.n += size
	}
}
