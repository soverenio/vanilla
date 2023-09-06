package random

import (
	"math/rand"
	"sync"
)

type safeRand struct {
	lock       sync.Mutex
	parentRand *rand.Rand
}

func newSafeRand(seed int64) safeRand {
	return safeRand{parentRand: rand.New(rand.NewSource(seed))} //nolint:gosec //  G404: Use of weak random number generator (math/rand instead of crypto/rand)
}

func (sr *safeRand) Intn(n int) int {
	sr.lock.Lock()
	v := sr.parentRand.Intn(n)
	sr.lock.Unlock()
	return v
}

func (sr *safeRand) Perm(n int) []int {
	sr.lock.Lock()
	v := sr.parentRand.Perm(n)
	sr.lock.Unlock()
	return v

}
