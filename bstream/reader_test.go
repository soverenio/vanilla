package bstream

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type ContinuousByteReader struct{}

func (c ContinuousByteReader) Read(p []byte) (n int, err error) {
	bts := 0
	for i := 0; i < 100 && i < len(p); i++ {
		p[i] = 5
		bts++
	}
	return bts, nil
}

var (
	gr1 int8
	gr2 int16
	gr3 int32
	gr4 int64

	gr5 uint8
	gr6 uint16
	gr7 uint32
	gr8 uint64

	gs  string
	gbs []byte
)

func BenchmarkBinaryStreamReader1(b *testing.B) {
	runtime.MemProfileRate = 1

	b.ResetTimer()
	stream := NewBinaryReaderStream(bufio.NewReader(ContinuousByteReader{}))

	var (
		r1 int8
		r2 int16
		r3 int32
		r4 int64

		r5 uint8
		r6 uint16
		r7 uint32
		r8 uint64

		s  string
		bs []byte
	)

	for i := 0; i < b.N; i++ {
		r1 = stream.ReadInt8()
		r2 = stream.ReadInt16()
		r3 = stream.ReadInt32()
		r4 = stream.ReadInt64()

		r5 = stream.ReadUint8()
		r6 = stream.ReadUint16()
		r7 = stream.ReadUint32()
		r8 = stream.ReadUint64()

		s = stream.ReadString(10)
		bs = stream.ReadBytes(10)
		stream.Discard(10)
	}

	gr1 = r1
	gr2 = r2
	gr3 = r3
	gr4 = r4

	gr5 = r5
	gr6 = r6
	gr7 = r7
	gr8 = r8

	gs = s
	gbs = bs

	b.ReportAllocs()
}

func BenchmarkNewBinaryStream(b *testing.B) {
	runtime.MemProfileRate = 1

	var a []byte
	a = binary.BigEndian.AppendUint16(a, 100)
	a = binary.BigEndian.AppendUint32(a, 128)
	a = append(a, bytes.Repeat([]byte("hello"), 100)...)
	a = binary.BigEndian.AppendUint64(a, 0)
	a = binary.BigEndian.AppendUint64(a, 0)
	a = binary.BigEndian.AppendUint64(a, 0)
	a = binary.BigEndian.AppendUint64(a, 0)
	// a = binary.bigEndian.AppendUint64(a, 0)
	// a = binary.bigEndian.AppendUint64(a, 0)
	// a = binary.bigEndian.AppendUint64(a, 0)
	// a = binary.bigEndian.AppendUint64(a, 0)
	// a = binary.bigEndian.AppendUint64(a, 0)
	// a = binary.bigEndian.AppendUint64(a, 0)

	bb := bytes.NewBuffer(a)
	s := bufio.NewReader(bb)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream := NewBinaryReaderStream(s)
		val1 := stream.ReadUint16()
		assert.EqualValues(b, 100, val1)
		assert.EqualValues(b, 2, stream.cnt)
		val2 := stream.ReadUint32()
		assert.EqualValues(b, 128, val2)
		assert.EqualValues(b, 6, stream.cnt)
		stream.Discard(500)
		assert.EqualValues(b, 506, stream.cnt)
		require.NoError(b, stream.Error())
		b.StopTimer()

		s = bufio.NewReader(bytes.NewBuffer(a))
	}
	b.ReportAllocs()
}
