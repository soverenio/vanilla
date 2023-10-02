package cache

import (
	"runtime"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestStruct struct {
	Num      int
	Children []*TestStruct
}

func TestCache(t *testing.T) {
	tc := New[string, float64](0, 0)
	defer tc.Stop()

	a, found := tc.Get("a")
	assert.False(t, found)
	assert.EqualValues(t, 0, a)

	b, found := tc.Get("b")
	assert.False(t, found)
	assert.EqualValues(t, 0, b)

	c, found := tc.Get("c")
	assert.False(t, found)
	assert.EqualValues(t, 0, c)

	tc.Set("a", 1, DefaultExpiration)
	tc.Set("b", 2, DefaultExpiration)
	tc.Set("c", 3.5, DefaultExpiration)

	x, found := tc.Get("a")
	assert.True(t, found, "element 'a' not found")
	assert.EqualValues(t, 1, x)

	x, found = tc.Get("b")
	assert.True(t, found, "element 'b' not found")
	assert.EqualValues(t, 2, x)

	x, found = tc.Get("c")
	assert.True(t, found, "element 'c' fufnot found")
	assert.EqualValues(t, 3.5, x)
}

func TestCacheTimes(t *testing.T) {
	tc := New[string, int](50*time.Millisecond, 1*time.Millisecond)
	defer tc.Stop()

	tc.Set("a", 1, DefaultExpiration)
	tc.Set("b", 2, NoExpiration)
	tc.Set("c", 3, 20*time.Millisecond)
	tc.Set("d", 4, 200*time.Millisecond)

	<-time.After(25 * time.Millisecond)
	_, found := tc.Get("c")
	assert.False(t, found, "Found c when it should have been automatically deleted")

	<-time.After(30 * time.Millisecond)
	_, found = tc.Get("a")
	assert.False(t, found, "Found a when it should have been automatically deleted")

	_, found = tc.Get("b")
	assert.True(t, found, "Did not find b even though it was set to never expire")

	_, found = tc.Get("d")
	assert.True(t, found, "Did not find d even though it was set to expire later than the default")

	<-time.After(200 * time.Millisecond)
	_, found = tc.Get("d")
	assert.False(t, found, "Found d when it should have been automatically deleted (later than the default)")
}

func TestNewFromItems(t *testing.T) {
	m := map[string]Item[int]{
		"a": {Object: 1, Expiration: time.Time{}},
		"b": {Object: 2, Expiration: time.Time{}},
	}
	tc := NewFromItems(DefaultExpiration, 0, m)
	defer tc.Stop()

	a, found := tc.Get("a")
	assert.True(t, found)
	assert.Equal(t, 1, a)

	b, found := tc.Get("b")
	assert.True(t, found)
	assert.Equal(t, 2, b)
}

func TestStorePointerToStruct(t *testing.T) {
	tc := New[string, *TestStruct](DefaultExpiration, 0)
	defer tc.Stop()

	tc.Set("foo", &TestStruct{Num: 1}, DefaultExpiration)
	x, found := tc.Get("foo")
	assert.True(t, found, "*TestStruct was not found for foo")
	x.Num++

	y, found := tc.Get("foo")
	assert.True(t, found, "*TestStruct was not found for foo (second time)")
	assert.Equal(t, 2, y.Num)
}

func TestOnEvicted(t *testing.T) {
	tc := New[string, int](DefaultExpiration, 0)
	tc.Set("foo", 3, DefaultExpiration)
	require.Nil(t, tc.onDeletion, "tc.onDeletion is not nil")

	works := false
	tc.OnDeletion(func(k string, v int, s OnDeletionStatus) {
		if k == "foo" && v == 3 {
			works = true
		}
		tc.Set("bar", 4, DefaultExpiration)
	})

	tc.Delete("foo")
	x, _ := tc.Get("bar")
	assert.True(t, works)
	assert.Equal(t, 4, x)
}

func BenchmarkCacheGetExpiring(b *testing.B) {
	benchmarkCacheGet(b, 5*time.Minute)
}

func BenchmarkCacheGetNotExpiring(b *testing.B) {
	benchmarkCacheGet(b, NoExpiration)
}

func benchmarkCacheGet(b *testing.B, exp time.Duration) {
	b.StopTimer()
	tc := New[string, string](exp, 0)
	defer tc.Stop()

	tc.Set("foo", "bar", DefaultExpiration)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tc.Get("foo")
	}
}

func BenchmarkRWMutexMapGet(b *testing.B) {
	b.StopTimer()
	m := map[string]string{
		"foo": "bar",
	}
	mu := sync.RWMutex{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mu.RLock()
		_, _ = m["foo"] //nolint:gosimple // S1005: unnecessary assignment to the blank identifier
		mu.RUnlock()
	}
}

func BenchmarkRWMutexInterfaceMapGetStruct(b *testing.B) {
	b.StopTimer()
	s := struct{ name string }{name: "foo"}
	m := map[interface{}]string{
		s: "bar",
	}
	mu := sync.RWMutex{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mu.RLock()
		_, _ = m[s] //nolint:gosimple // S1005: unnecessary assignment to the blank identifier
		mu.RUnlock()
	}
}

func BenchmarkRWMutexInterfaceMapGetString(b *testing.B) {
	b.StopTimer()
	m := map[interface{}]string{
		"foo": "bar",
	}
	mu := sync.RWMutex{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mu.RLock()
		_, _ = m["foo"] //nolint:gosimple // S1005: unnecessary assignment to the blank identifier
		mu.RUnlock()
	}
}

func BenchmarkCacheGetConcurrentExpiring(b *testing.B) {
	benchmarkCacheGetConcurrent(b, 5*time.Minute)
}

func BenchmarkCacheGetConcurrentNotExpiring(b *testing.B) {
	benchmarkCacheGetConcurrent(b, NoExpiration)
}

func benchmarkCacheGetConcurrent(b *testing.B, exp time.Duration) {
	b.StopTimer()
	tc := New[string, string](exp, 0)
	defer tc.Stop()

	tc.Set("foo", "bar", DefaultExpiration)
	wg := new(sync.WaitGroup)
	workers := runtime.NumCPU()
	each := b.N / workers
	wg.Add(workers)
	b.StartTimer()
	for i := 0; i < workers; i++ {
		go func() {
			for j := 0; j < each; j++ {
				tc.Get("foo")
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkRWMutexMapGetConcurrent(b *testing.B) {
	b.StopTimer()
	m := map[string]string{
		"foo": "bar",
	}
	mu := sync.RWMutex{}
	wg := new(sync.WaitGroup)
	workers := runtime.NumCPU()
	each := b.N / workers
	wg.Add(workers)
	b.StartTimer()
	for i := 0; i < workers; i++ {
		go func() {
			for j := 0; j < each; j++ {
				mu.RLock()
				_, _ = m["foo"] //nolint:gosimple // S1005: unnecessary assignment to the blank identifier
				mu.RUnlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkCacheGetManyConcurrentExpiring(b *testing.B) {
	benchmarkCacheGetManyConcurrent(b, 5*time.Minute)
}

func BenchmarkCacheGetManyConcurrentNotExpiring(b *testing.B) {
	benchmarkCacheGetManyConcurrent(b, NoExpiration)
}

func benchmarkCacheGetManyConcurrent(b *testing.B, exp time.Duration) {
	// This is the same as BenchmarkCacheGetConcurrent, but its result
	// can be compared against BenchmarkShardedCacheGetManyConcurrent
	// in sharded_test.go.
	b.StopTimer()
	n := 10000
	tc := New[string, string](exp, 0)
	defer tc.Stop()

	keys := make([]string, n)
	for i := 0; i < n; i++ {
		k := "foo" + strconv.Itoa(i)
		keys[i] = k
		tc.Set(k, "bar", DefaultExpiration)
	}
	each := b.N / n
	wg := new(sync.WaitGroup)
	wg.Add(n)
	for _, v := range keys {
		go func(k string) {
			for j := 0; j < each; j++ {
				tc.Get(k)
			}
			wg.Done()
		}(v)
	}
	b.StartTimer()
	wg.Wait()
}

func BenchmarkCacheSetExpiring(b *testing.B) {
	benchmarkCacheSet(b, 5*time.Minute)
}

func BenchmarkCacheSetNotExpiring(b *testing.B) {
	benchmarkCacheSet(b, NoExpiration)
}

func benchmarkCacheSet(b *testing.B, exp time.Duration) {
	b.StopTimer()
	tc := New[string, string](exp, 0)
	defer tc.Stop()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tc.Set("foo", "bar", DefaultExpiration)
	}
}

func BenchmarkRWMutexMapSet(b *testing.B) {
	b.StopTimer()
	m := map[string]string{}
	mu := sync.RWMutex{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		m["foo"] = "bar"
		mu.Unlock()
	}
}

func BenchmarkCacheSetDelete(b *testing.B) {
	b.StopTimer()
	tc := New[string, string](DefaultExpiration, 0)
	defer tc.Stop()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tc.Set("foo", "bar", DefaultExpiration)
		tc.Delete("foo")
	}
}

func BenchmarkRWMutexMapSetDelete(b *testing.B) {
	b.StopTimer()
	m := map[string]string{}
	mu := sync.RWMutex{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		m["foo"] = "bar"
		mu.Unlock()
		mu.Lock()
		delete(m, "foo")
		mu.Unlock()
	}
}

func BenchmarkCacheSetDeleteSingleLock(b *testing.B) {
	b.StopTimer()
	tc := New[string, string](DefaultExpiration, 0)
	defer tc.Stop()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tc.mu.Lock()
		tc.set("foo", "bar", DefaultExpiration)
		tc.delete("foo")
		tc.mu.Unlock()
	}
}

func BenchmarkRWMutexMapSetDeleteSingleLock(b *testing.B) {
	b.StopTimer()
	m := map[string]string{}
	mu := sync.RWMutex{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		m["foo"] = "bar"
		delete(m, "foo")
		mu.Unlock()
	}
}

func BenchmarkDeleteExpiredLoop(b *testing.B) {
	b.StopTimer()
	tc := New[string, string](5*time.Minute, 0)
	defer tc.Stop()

	tc.mu.Lock()
	for i := 0; i < 100000; i++ {
		tc.set(strconv.Itoa(i), "bar", DefaultExpiration)
	}
	tc.mu.Unlock()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tc.DeleteExpired()
	}
}

func TestGetWithExpiration(t *testing.T) {
	tc := New[string, float64](time.Second, 0)
	defer tc.Stop()

	a, expiration, found := tc.GetWithExpiration("a")
	assert.False(t, found)
	assert.EqualValues(t, 0, a)
	assert.True(t, expiration.IsZero())

	b, expiration, found := tc.GetWithExpiration("d")
	assert.False(t, found)
	assert.EqualValues(t, 0, b)
	assert.True(t, expiration.IsZero())

	c, expiration, found := tc.GetWithExpiration("e")
	assert.False(t, found)
	assert.EqualValues(t, 0, c)
	assert.True(t, expiration.IsZero())

	tc.Set("a", 1, DefaultExpiration)
	tc.Set("d", 1, NoExpiration)
	tc.Set("e", 1, 50*time.Millisecond)

	x, expiration, found := tc.GetWithExpiration("a")
	assert.True(t, found)
	assert.EqualValues(t, 1, x)
	assert.False(t, expiration.IsZero())

	x, expiration, found = tc.GetWithExpiration("d")
	assert.True(t, found)
	assert.EqualValues(t, 1, x)
	assert.True(t, expiration.IsZero())

	x, expiration, found = tc.GetWithExpiration("e")
	assert.True(t, found)
	assert.EqualValues(t, 1, x)
	assert.False(t, expiration.IsZero())
	assert.Equal(t, tc.items["e"].Expiration, expiration)
	assert.True(t, time.Now().Before(expiration))
}
