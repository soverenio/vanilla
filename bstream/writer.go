package bstream

import (
	"io"
	"reflect"
	"unsafe"

	"github.com/soverenio/vanilla/throw"
)

type UnderlineWriterStream interface {
	Write(p []byte) (n int, err error)
}

type Flusher interface {
	Flush() error
}

type BinaryStreamWriter struct {
	u  UnderlineWriterStream
	bo EndiannessWriter

	cnt int
	err error
}

func NewBinaryWriterStream(writer UnderlineWriterStream) *BinaryStreamWriter {
	return &BinaryStreamWriter{
		u:   writer,
		bo:  BigEndian,
		cnt: 0,
		err: nil,
	}
}

func (s *BinaryStreamWriter) Flush() error {
	switch flusher, ok := s.u.(Flusher); {
	case s.err != nil:
		return s.err
	case ok:
		return flusher.Flush()
	default:
		return nil
	}
}

func (s *BinaryStreamWriter) Reset() *BinaryStreamWriter {
	s.cnt = 0
	s.err = nil

	return s
}

func (s *BinaryStreamWriter) SetEndianness(e EndiannessWriter) *BinaryStreamWriter {
	s.bo = e
	return s
}

func (s *BinaryStreamWriter) Endianness() Endianness {
	return s.bo.(Endianness)
}

func (s *BinaryStreamWriter) Error() error {
	return s.err
}

func (s *BinaryStreamWriter) Padding(byteCount int) {
	if s.err != nil {
		return
	}

	switch n, err := io.CopyN(s.u, Zero, int64(byteCount)); {
	case err != nil:
		s.err = err
	case int(n) != byteCount:
		panic(throw.IllegalState())
	default:
		s.cnt += int(n)
	}
}

func (s *BinaryStreamWriter) WriteInt8(inp int8) {
	if s.err != nil {
		return
	}

	b := [1]byte{byte(inp)}

	switch n, err := s.u.Write(b[:]); {
	case err != nil:
		s.err = err
	case n < len(b):
		s.err = io.EOF
	default:
		s.cnt += n
	}
}

func (s *BinaryStreamWriter) WriteInt16(inp int16) {
	if s.err != nil {
		return
	}

	b := s.bo.WriteInt16Array(inp)

	switch n, err := s.u.Write(b[:]); {
	case err != nil:
		s.err = err
	case n < len(b):
		s.err = io.EOF
	default:
		s.cnt += n
	}
}

func (s *BinaryStreamWriter) WriteInt32(inp int32) {
	if s.err != nil {
		return
	}

	b := s.bo.WriteInt32Array(inp)

	switch n, err := s.u.Write(b[:]); {
	case err != nil:
		s.err = err
	case n < len(b):
		s.err = io.EOF
	default:
		s.cnt += n
	}
}

func (s *BinaryStreamWriter) WriteInt64(inp int64) {
	if s.err != nil {
		return
	}

	b := s.bo.WriteInt64Array(inp)

	switch n, err := s.u.Write(b[:]); {
	case err != nil:
		s.err = err
	case n < len(b):
		s.err = io.EOF
	default:
		s.cnt += n
	}
}

func (s *BinaryStreamWriter) WriteUint8(inp uint8) {
	if s.err != nil {
		return
	}

	b := [1]byte{inp}

	switch n, err := s.u.Write(b[:]); {
	case err != nil:
		s.err = err
	case n < len(b):
		s.err = io.EOF
	default:
		s.cnt += n
	}
}

func (s *BinaryStreamWriter) WriteUint16(inp uint16) {
	if s.err != nil {
		return
	}

	b := s.bo.WriteUint16Array(inp)

	switch n, err := s.u.Write(b[:]); {
	case err != nil:
		s.err = err
	case n < len(b):
		s.err = io.EOF
	default:
		s.cnt += n
	}
}

func (s *BinaryStreamWriter) WriteUint32(inp uint32) {
	if s.err != nil {
		return
	}

	b := s.bo.WriteUint32Array(inp)

	switch n, err := s.u.Write(b[:]); {
	case err != nil:
		s.err = err
	case n < len(b):
		s.err = io.EOF
	default:
		s.cnt += n
	}
}

func (s *BinaryStreamWriter) WriteUint64(inp uint64) {
	if s.err != nil {
		return
	}

	b := s.bo.WriteUint64Array(inp)

	switch n, err := s.u.Write(b[:]); {
	case err != nil:
		s.err = err
	case n < len(b):
		s.err = io.EOF
	default:
		s.cnt += n
	}
}

func (s *BinaryStreamWriter) WriteString(inp string) {
	if s.err != nil {
		return
	}

	b := *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&inp))))

	switch n, err := s.u.Write(b); {
	case err != nil:
		s.err = err
	case n < len(b):
		s.err = io.EOF
	default:
		s.cnt += n
	}
}

func (s *BinaryStreamWriter) WriteBytes(inp []byte) {
	if s.err != nil {
		return
	}

	switch n, err := s.u.Write(inp); {
	case err != nil:
		s.err = err
	case n < len(inp):
		s.err = io.EOF
	default:
		s.cnt += n
	}
}
