package bstream

type Endianness interface {
	EndiannessReader
	EndiannessWriter

	String() string
}

type EndiannessReader interface {
	ReadInt16Array(inp [2]byte) int16
	ReadInt32Array(inp [4]byte) int32
	ReadInt64Array(inp [8]byte) int64
	ReadUint16Array(inp [2]byte) uint16
	ReadUint32Array(inp [4]byte) uint32
	ReadUint64Array(inp [8]byte) uint64

	ReadInt16Slice(inp []byte) int16
	ReadInt32Slice(inp []byte) int32
	ReadInt64Slice(inp []byte) int64
	ReadUint16Slice(inp []byte) uint16
	ReadUint32Slice(inp []byte) uint32
	ReadUint64Slice(inp []byte) uint64
}

type EndiannessWriter interface {
	WriteInt16Array(inp int16) [2]byte
	WriteInt32Array(inp int32) [4]byte
	WriteInt64Array(inp int64) [8]byte
	WriteUint16Array(inp uint16) [2]byte
	WriteUint32Array(inp uint32) [4]byte
	WriteUint64Array(inp uint64) [8]byte
}

var (
	BigEndian    Endianness = bigEndian{}
	LittleEndian Endianness = littleEndian{}
)

type littleEndian struct{}

func (e littleEndian) String() string {
	return "bstream.LittleEndian"
}

func (e littleEndian) WriteInt16Array(inp int16) [2]byte {
	return [2]byte{
		byte(inp),
		byte(inp >> 8),
	}
}

func (e littleEndian) WriteInt32Array(inp int32) [4]byte {
	return [4]byte{
		byte(inp),
		byte(inp >> 8),
		byte(inp >> 16),
		byte(inp >> 24),
	}
}

func (e littleEndian) WriteInt64Array(inp int64) [8]byte {
	return [8]byte{
		byte(inp),
		byte(inp >> 8),
		byte(inp >> 16),
		byte(inp >> 24),
		byte(inp >> 32),
		byte(inp >> 40),
		byte(inp >> 48),
		byte(inp >> 56),
	}
}

func (e littleEndian) WriteUint16Array(inp uint16) [2]byte {
	return [2]byte{
		byte(inp),
		byte(inp >> 8),
	}
}

func (e littleEndian) WriteUint32Array(inp uint32) [4]byte {
	return [4]byte{
		byte(inp),
		byte(inp >> 8),
		byte(inp >> 16),
		byte(inp >> 24),
	}
}

func (e littleEndian) WriteUint64Array(inp uint64) [8]byte {
	return [8]byte{
		byte(inp),
		byte(inp >> 8),
		byte(inp >> 16),
		byte(inp >> 24),
		byte(inp >> 32),
		byte(inp >> 40),
		byte(inp >> 48),
		byte(inp >> 56),
	}
}

func (e littleEndian) ReadInt16Array(inp [2]byte) int16 {
	return int16(inp[0]) | int16(inp[1])<<8
}

func (e littleEndian) ReadInt32Array(inp [4]byte) int32 {
	return int32(inp[0]) | int32(inp[1])<<8 | int32(inp[2])<<16 | int32(inp[3])<<24
}

func (e littleEndian) ReadInt64Array(inp [8]byte) int64 {
	return int64(inp[0]) | int64(inp[1])<<8 | int64(inp[2])<<16 | int64(inp[3])<<24 | int64(inp[4])<<32 | int64(inp[5])<<40 | int64(inp[6])<<48 | int64(inp[7])<<56
}

func (e littleEndian) ReadUint16Array(inp [2]byte) uint16 {
	return uint16(inp[0]) | uint16(inp[1])<<8
}

func (e littleEndian) ReadUint32Array(inp [4]byte) uint32 {
	return uint32(inp[0]) | uint32(inp[1])<<8 | uint32(inp[2])<<16 | uint32(inp[3])<<24
}

func (e littleEndian) ReadUint64Array(inp [8]byte) uint64 {
	return uint64(inp[0]) | uint64(inp[1])<<8 | uint64(inp[2])<<16 | uint64(inp[3])<<24 | uint64(inp[4])<<32 | uint64(inp[5])<<40 | uint64(inp[6])<<48 | uint64(inp[7])<<56
}

func (e littleEndian) ReadInt16Slice(inp []byte) int16 {
	_ = inp[0]
	return int16(inp[0]) | int16(inp[1])<<8
}

func (e littleEndian) ReadInt32Slice(inp []byte) int32 {
	_ = inp[0]
	return int32(inp[0]) | int32(inp[1])<<8 | int32(inp[2])<<16 | int32(inp[3])<<24
}

func (e littleEndian) ReadInt64Slice(inp []byte) int64 {
	_ = inp[0]
	return int64(inp[0]) | int64(inp[1])<<8 | int64(inp[2])<<16 | int64(inp[3])<<24 | int64(inp[4])<<32 | int64(inp[5])<<40 | int64(inp[6])<<48 | int64(inp[7])<<56
}

func (e littleEndian) ReadUint16Slice(inp []byte) uint16 {
	_ = inp[0]
	return uint16(inp[0]) | uint16(inp[1])<<8
}

func (e littleEndian) ReadUint32Slice(inp []byte) uint32 {
	_ = inp[0]
	return uint32(inp[0]) | uint32(inp[1])<<8 | uint32(inp[2])<<16 | uint32(inp[3])<<24
}

func (e littleEndian) ReadUint64Slice(inp []byte) uint64 {
	_ = inp[0]
	return uint64(inp[0]) | uint64(inp[1])<<8 | uint64(inp[2])<<16 | uint64(inp[3])<<24 | uint64(inp[4])<<32 | uint64(inp[5])<<40 | uint64(inp[6])<<48 | uint64(inp[7])<<56
}

type bigEndian struct{}

func (e bigEndian) WriteInt16Array(inp int16) [2]byte {
	return [2]byte{
		byte(inp >> 8),
		byte(inp),
	}
}

func (e bigEndian) WriteInt32Array(inp int32) [4]byte {
	return [4]byte{
		byte(inp >> 24),
		byte(inp >> 16),
		byte(inp >> 8),
		byte(inp),
	}
}

func (e bigEndian) WriteInt64Array(inp int64) [8]byte {
	return [8]byte{
		byte(inp >> 56),
		byte(inp >> 48),
		byte(inp >> 40),
		byte(inp >> 32),
		byte(inp >> 24),
		byte(inp >> 16),
		byte(inp >> 8),
		byte(inp),
	}
}

func (e bigEndian) WriteUint16Array(inp uint16) [2]byte {
	return [2]byte{
		byte(inp >> 8),
		byte(inp),
	}
}

func (e bigEndian) WriteUint32Array(inp uint32) [4]byte {
	return [4]byte{
		byte(inp >> 24),
		byte(inp >> 16),
		byte(inp >> 8),
		byte(inp),
	}
}

func (e bigEndian) WriteUint64Array(inp uint64) [8]byte {
	return [8]byte{
		byte(inp >> 56),
		byte(inp >> 48),
		byte(inp >> 40),
		byte(inp >> 32),
		byte(inp >> 24),
		byte(inp >> 16),
		byte(inp >> 8),
		byte(inp),
	}
}

func (e bigEndian) ReadInt16Array(inp [2]byte) int16 {
	return int16(inp[1]) | int16(inp[0])<<8
}

func (e bigEndian) ReadInt32Array(inp [4]byte) int32 {
	return int32(inp[3]) | int32(inp[2])<<8 | int32(inp[1])<<16 | int32(inp[0])<<24
}

func (e bigEndian) ReadInt64Array(inp [8]byte) int64 {
	return int64(inp[7]) | int64(inp[6])<<8 | int64(inp[5])<<16 | int64(inp[4])<<24 | int64(inp[3])<<32 | int64(inp[2])<<40 | int64(inp[1])<<48 | int64(inp[0])<<56
}

func (e bigEndian) ReadUint16Array(inp [2]byte) uint16 {
	return uint16(inp[1]) | uint16(inp[0])<<8
}

func (e bigEndian) ReadUint32Array(inp [4]byte) uint32 {
	return uint32(inp[3]) | uint32(inp[2])<<8 | uint32(inp[1])<<16 | uint32(inp[0])<<24
}

func (e bigEndian) ReadUint64Array(inp [8]byte) uint64 {
	return uint64(inp[7]) | uint64(inp[6])<<8 | uint64(inp[5])<<16 | uint64(inp[4])<<24 | uint64(inp[3])<<32 | uint64(inp[2])<<40 | uint64(inp[1])<<48 | uint64(inp[0])<<56
}

func (e bigEndian) ReadInt16Slice(inp []byte) int16 {
	_ = inp[0]
	return int16(inp[1]) | int16(inp[0])<<8
}

func (e bigEndian) ReadInt32Slice(inp []byte) int32 {
	_ = inp[0]
	return int32(inp[3]) | int32(inp[2])<<8 | int32(inp[1])<<16 | int32(inp[0])<<24
}

func (e bigEndian) ReadInt64Slice(inp []byte) int64 {
	_ = inp[0]
	return int64(inp[7]) | int64(inp[6])<<8 | int64(inp[5])<<16 | int64(inp[4])<<24 | int64(inp[3])<<32 | int64(inp[2])<<40 | int64(inp[1])<<48 | int64(inp[0])<<56
}

func (e bigEndian) ReadUint16Slice(inp []byte) uint16 {
	_ = inp[0]
	return uint16(inp[1]) | uint16(inp[0])<<8
}

func (e bigEndian) ReadUint32Slice(inp []byte) uint32 {
	_ = inp[0]
	return uint32(inp[3]) | uint32(inp[2])<<8 | uint32(inp[1])<<16 | uint32(inp[0])<<24
}

func (e bigEndian) ReadUint64Slice(inp []byte) uint64 {
	_ = inp[0]
	return uint64(inp[7]) | uint64(inp[6])<<8 | uint64(inp[5])<<16 | uint64(inp[4])<<24 | uint64(inp[3])<<32 | uint64(inp[2])<<40 | uint64(inp[1])<<48 | uint64(inp[0])<<56
}

func (e bigEndian) String() string {
	return "bstream.BigEndian"
}
