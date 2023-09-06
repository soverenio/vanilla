package bstream

import (
	"encoding/binary"
	"io"

	"github.com/soverenio/vanilla/throw"
)

var (
	int8ByteSize  = binary.Size(int8(1))
	int16ByteSize = binary.Size(int16(1))
	int32ByteSize = binary.Size(int32(1))
	int64ByteSize = binary.Size(int64(1))

	uint8ByteSize  = binary.Size(uint8(1))
	uint16ByteSize = binary.Size(uint16(1))
	uint32ByteSize = binary.Size(uint32(1))
	uint64ByteSize = binary.Size(uint64(1))
)

type UnderlineReaderStream interface {
	Read(p []byte) (n int, err error)
	Peek(n int) ([]byte, error)
	Discard(n int) (discarded int, err error)
}

type BinaryStreamReader struct {
	u  UnderlineReaderStream
	ur io.Reader
	bo EndiannessReader

	cnt int
	err error
}

func NewBinaryReaderStream(reader UnderlineReaderStream) *BinaryStreamReader {
	return &BinaryStreamReader{
		u:   reader,
		ur:  reader,
		bo:  BigEndian,
		cnt: 0,
		err: nil,
	}
}

func (s *BinaryStreamReader) Reset() *BinaryStreamReader {
	s.cnt = 0
	s.err = nil

	return s
}

func (s *BinaryStreamReader) SetEndianness(e EndiannessReader) *BinaryStreamReader {
	s.bo = e
	return s
}

func (s *BinaryStreamReader) Endianness() Endianness {
	return s.bo.(Endianness)
}

func (s *BinaryStreamReader) Error() error {
	return s.err
}

func (s *BinaryStreamReader) WithError(err error) {
	s.err = err
}

func (s *BinaryStreamReader) ByteCount() int {
	return s.cnt
}

func (s *BinaryStreamReader) Discard(byteCount int) {
	switch {
	case s.err != nil:
		return
	case byteCount == 0:
		return
	case byteCount < 0:
		panic(throw.IllegalValue())
	}

	_, err := s.u.Discard(byteCount)
	if err != nil {
		s.err = err
	} else {
		s.cnt += byteCount
	}
}

func (s *BinaryStreamReader) ReadInt8() int8 {
	if s.err != nil {
		return 0
	}

	b, err := s.u.Peek(int8ByteSize)
	switch {
	case err != nil:
		s.err = err
		return 0
	case len(b) != int8ByteSize:
		panic(throw.IllegalState())
	default:
		s.cnt += int8ByteSize
	}

	_, err = s.u.Discard(int8ByteSize)
	if err != nil {
		s.err = err
		return 0
	}

	return int8(b[0])
}

func (s *BinaryStreamReader) ReadInt16() int16 {
	if s.err != nil {
		return 0
	}

	b, err := s.u.Peek(int16ByteSize)
	switch {
	case err != nil:
		s.err = err
		return 0
	case len(b) != int16ByteSize:
		panic(throw.IllegalState())
	default:
		s.cnt += int16ByteSize
	}

	_, err = s.u.Discard(int16ByteSize)
	if err != nil {
		s.err = err
		return 0
	}

	return s.bo.ReadInt16Slice(b)
}

func (s *BinaryStreamReader) ReadInt32() int32 {
	if s.err != nil {
		return 0
	}

	b, err := s.u.Peek(int32ByteSize)
	switch {
	case err != nil:
		s.err = err
		return 0
	case len(b) != int32ByteSize:
		panic(throw.IllegalState())
	default:
		s.cnt += int32ByteSize
	}

	_, err = s.u.Discard(int32ByteSize)
	if err != nil {
		s.err = err
		return 0
	}

	return s.bo.ReadInt32Slice(b)
}

func (s *BinaryStreamReader) ReadInt64() int64 {
	if s.err != nil {
		return 0
	}

	b, err := s.u.Peek(int64ByteSize)
	switch {
	case err != nil:
		s.err = err
		return 0
	case len(b) != int64ByteSize:
		panic(throw.IllegalState())
	default:
		s.cnt += int64ByteSize
	}

	_, err = s.u.Discard(int64ByteSize)
	if err != nil {
		s.err = err
		return 0
	}

	return s.bo.ReadInt64Slice(b)
}

func (s *BinaryStreamReader) ReadUint8() uint8 {
	if s.err != nil {
		return 0
	}

	b, err := s.u.Peek(uint8ByteSize)
	switch {
	case err != nil:
		s.err = err
		return 0
	case len(b) != uint8ByteSize:
		panic(throw.IllegalState())
	default:
		s.cnt += uint8ByteSize
	}

	_, err = s.u.Discard(uint8ByteSize)
	if err != nil {
		s.err = err
		return 0
	}

	return b[0]
}

func (s *BinaryStreamReader) ReadUint16() uint16 {
	if s.err != nil {
		return 0
	}

	b, err := s.u.Peek(uint16ByteSize)
	switch {
	case err != nil:
		s.err = err
		return 0
	case len(b) != uint16ByteSize:
		panic(throw.IllegalState())
	default:
		s.cnt += uint16ByteSize
	}

	_, err = s.u.Discard(uint16ByteSize)
	if err != nil {
		s.err = err
		return 0
	}

	return s.bo.ReadUint16Slice(b)
}

func (s *BinaryStreamReader) ReadUint32() uint32 {
	if s.err != nil {
		return 0
	}

	b, err := s.u.Peek(uint32ByteSize)
	switch {
	case err != nil:
		s.err = err
		return 0
	case len(b) != uint32ByteSize:
		panic(throw.IllegalState())
	default:
		s.cnt += uint32ByteSize
	}

	_, err = s.u.Discard(uint32ByteSize)
	if err != nil {
		s.err = err
		return 0
	}

	return s.bo.ReadUint32Slice(b)
}

func (s *BinaryStreamReader) ReadUint64() uint64 {
	if s.err != nil {
		return 0
	}

	b, err := s.u.Peek(uint64ByteSize)
	switch {
	case err != nil:
		s.err = err
		return 0
	case len(b) != uint64ByteSize:
		panic(throw.IllegalState())
	default:
		s.cnt += uint64ByteSize
	}

	_, err = s.u.Discard(uint64ByteSize)
	if err != nil {
		s.err = err
		return 0
	}

	return s.bo.ReadUint64Slice(b)
}

func (s *BinaryStreamReader) ReadString(byteCount int) string {
	switch {
	case s.err != nil:
		return ""
	case byteCount == 0:
		return ""
	case byteCount < 0:
		panic(throw.IllegalValue())
	}

	b := make([]byte, byteCount)

	switch n, err := io.ReadFull(s.ur, b); {
	case err != nil:
		s.err = err
		return ""
	case n != len(b):
		panic(throw.IllegalState())
	default:
		s.cnt += n

		return string(b)
	}
}

func (s *BinaryStreamReader) ReadBytes(byteCount int) []byte {
	switch {
	case s.err != nil:
		return nil
	case byteCount == 0:
		return nil
	case byteCount < 0:
		panic(throw.IllegalValue())
	}

	b := make([]byte, byteCount)

	switch n, err := io.ReadFull(s.ur, b); {
	case err != nil:
		s.err = err
		return nil
	case n != len(b):
		panic(throw.IllegalState())
	default:
		s.cnt += n

		return b
	}
}

func (s *BinaryStreamReader) Read(out []byte) {
	if s.err != nil {
		return
	} else if len(out) == 0 {
		return
	}

	switch n, err := io.ReadFull(s.ur, out); {
	case err != nil:
		s.err = err
	case n != len(out):
		panic(throw.IllegalState())
	default:
		s.cnt += n
	}
}
