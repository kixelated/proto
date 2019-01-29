package proto

type Writer struct {
	data []byte
	s    *Sizer
}

func (w *Writer) Grow(size int) {
	if cap(w.data) > size {
		return
	}

	temp := make([]byte, len(w.data), size)
	copy(temp, w.data)
	w.data = temp
}

func (w *Writer) WithSizer(s *Sizer) {
	w.Grow(s.Size())
	w.s = s
}

func (w *Writer) Size() int {
	return len(w.data)
}

func (w *Writer) Data() []byte {
	return w.data
}
