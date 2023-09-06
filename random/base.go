package random

import (
	"bufio"
	cRand "crypto/rand"
	"time"

	"github.com/soverenio/vanilla/bstream"
	"github.com/soverenio/vanilla/throw"
	"golang.org/x/exp/constraints"
)

var (
	bStreamRandom *bstream.BinaryStreamReader
	mRandInstance safeRand
)

func SignedMathRandBetween[T constraints.Signed](l, r T) T {
	if r < l {
		panic(throw.IllegalValue())
	}
	return T(mRandInstance.Intn(int(r-l+1))) + l
}

func SignedCryptoRandBetween[T constraints.Signed](l, r T) T {
	if r < l {
		panic(throw.IllegalValue())
	}
	return T(bStreamRandom.ReadInt64()%(int64(r)-int64(l)+1) + int64(l))
}

func init() {
	mRandInstance = newSafeRand(time.Now().UnixNano())
	bStreamRandom = bstream.NewBinaryReaderStream(bufio.NewReader(cRand.Reader))
}
