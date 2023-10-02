package cache

import (
	"context"
	"sync"
	"time"

	"github.com/soverenio/vanilla/defcon"
	"github.com/soverenio/vanilla/zero"

	"github.com/soverenio/vanilla/throw"
)

type Item[V any] struct {
	Object     V
	Expiration time.Time
}

// Expired Returns true if the item has expired.
func (item Item[V]) Expired() bool {
	if item.Expiration.IsZero() {
		return false
	}
	return time.Now().After(item.Expiration)
}

const (
	// NoExpiration is userd with functions that take an expiration time.
	NoExpiration time.Duration = -2

	// DefaultExpiration is used with functions that take an expiration time. Equivalent to
	// passing in the same expiration duration as was given to New() or NewFrom() when the
	// cache was created (e.g. 5 minutes.)
	DefaultExpiration time.Duration = -1
)

type OnDeletionStatus int

const (
	_ OnDeletionStatus = iota
	OnDeletionStatusEvicted
	OnDeletionStatusDeleted
)

var (
	ErrNotExists     = throw.New("record doesn't exists")
	ErrAlreadyExists = throw.New("record already exists")
)

type Cache[K comparable, V any] struct {
	*cache[K, V]
	janitor *janitor[K, V]
}

type cache[K comparable, V any] struct {
	defaultExpiration time.Duration
	items             map[K]Item[V]
	mu                sync.RWMutex
	onDeletion        func(K, V, OnDeletionStatus)
}

func (c *cache[K, V]) getExpiration(exp time.Duration) time.Time {
	var rv time.Time
	if exp == DefaultExpiration {
		exp = c.defaultExpiration
	}
	if exp > 0 {
		rv = time.Now().Add(exp)
	}
	return rv
}

// Set an item to the cache, replacing any existing item. If the duration is 0
// (DefaultExpiration), the cache's default expiration time is used. If it is -1
// (NoExpiration), the item never expires.
func (c *cache[K, V]) Set(k K, x V, d time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.set(k, x, d)
}

func (c *cache[K, V]) Touch(k K, d time.Duration) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, ok := c.items[k]
	if !ok || val.Expired() {
		return false
	}

	val.Expiration = c.getExpiration(d)
	c.items[k] = val
	return true
}

func (c *cache[K, V]) set(k K, x V, d time.Duration) {
	c.items[k] = Item[V]{
		Object:     x,
		Expiration: c.getExpiration(d),
	}
}

// SetDefault sets an item to the cache, replacing any existing item, using the default
// expiration.
func (c *cache[K, V]) SetDefault(k K, x V) {
	c.Set(k, x, DefaultExpiration)
}

// Add an item to the cache only if an item doesn't already exist for the given
// key, or if the existing item has expired. Returns an error otherwise.
func (c *cache[K, V]) Add(k K, x V, d time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, found := c.get(k)
	if found {
		return ErrAlreadyExists
	}
	c.set(k, x, d)
	return nil
}

// Replace a new value for the cache key only if it already exists, and the existing
// item hasn't expired. Returns an error otherwise.
func (c *cache[K, V]) Replace(k K, x V, d time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, found := c.get(k)
	if !found {
		return ErrNotExists
	}
	c.set(k, x, d)
	return nil
}

// Get an item from the cache. Returns the item or nil, and a bool indicating
// whether the key was found.
func (c *cache[K, V]) Get(k K) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[k]
	if !found || item.Expired() {
		return zero.Zero[V](), false
	}
	return item.Object, true
}

// GetWithTouch gets an item from the cache and update its expiration time. Returns the item or nil, and a bool indicating
// whether the key was found.
func (c *cache[K, V]) GetWithTouch(k K, d time.Duration) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, found := c.items[k]
	if !found || item.Expired() {
		return zero.Zero[V](), false
	}

	item.Expiration = c.getExpiration(d)
	c.items[k] = item
	return item.Object, true
}

// GetWithExpiration returns an item and its expiration time from the cache.
// It returns the item or nil, the expiration time if one is set (if the item
// never expires a zero value for time.Time is returned), and a bool indicating
// whether the key was found.
func (c *cache[K, V]) GetWithExpiration(k K) (V, time.Time, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// "Inlining" of get and Expired
	item, found := c.items[k]
	if !found {
		return zero.Zero[V](), time.Time{}, false
	}

	if item.Expired() {
		return zero.Zero[V](), time.Time{}, false
	}
	return item.Object, item.Expiration, true
}

func (c *cache[K, V]) get(k K) (V, bool) {
	item, found := c.items[k]
	if !found {
		return zero.Zero[V](), false
	}

	if item.Expired() {
		return zero.Zero[V](), false
	}
	return item.Object, true
}

// Delete an item from the cache. Does nothing if the key is not in the cache.
func (c *cache[K, V]) Delete(k K) bool {
	c.mu.Lock()
	v, deleted, evicted := c.delete(k)
	c.mu.Unlock()

	if evicted {
		c.onDeletion(k, v, OnDeletionStatusDeleted)
	}

	return deleted
}

func (c *cache[K, V]) delete(key K) (V, bool, bool) {
	switch v, found := c.items[key]; {
	case !found:
		return zero.Zero[V](), false, false
	case c.onDeletion != nil:
		delete(c.items, key)
		return v.Object, true, true
	default:
		delete(c.items, key)
		return zero.Zero[V](), true, false
	}
}

type keyAndValue[K comparable, V any] struct {
	key   K
	value V
}

// DeleteExpired deletes all expired items from the cache.
func (c *cache[K, V]) DeleteExpired() {
	var evictedItems []keyAndValue[K, V]

	c.mu.Lock()
	for k, v := range c.items {
		if !v.Expired() {
			continue
		}
		ov, _, evicted := c.delete(k)
		if evicted {
			evictedItems = append(evictedItems, keyAndValue[K, V]{k, ov})
		}
	}
	c.mu.Unlock()

	for _, v := range evictedItems {
		c.onDeletion(v.key, v.value, OnDeletionStatusEvicted)
	}
}

// OnDeletion sets an (optional) function that is called with the key and value when an
// item is evicted from the cache. (Including when it is deleted manually, but  not when
// it is overwritten). Set to nil to disable.
func (c *cache[K, V]) OnDeletion(f func(K, V, OnDeletionStatus)) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.onDeletion = f
}

// ItemList copies all unexpired items in the cache into a new map and returns it.
func (c *cache[K, V]) ItemList() map[K]Item[V] {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return defcon.IteratorFromMap(c.items, nil).Filter(func(k K, v Item[V]) bool {
		return !v.Expired()
	}).Finalize()
}

// ItemIterator copies all unexpired items in the cache into a new map and returns iterator.
func (c *cache[K, V]) ItemIterator() defcon.MapIterator[K, Item[V]] {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return defcon.IteratorFromMap(c.items, nil).Filter(func(k K, v Item[V]) bool {
		return !v.Expired()
	}).Intermediate(nil)
}

// ItemCount returns the number of items in the cache. This may include items that have
// expired, but have not yet been cleaned up.
func (c *cache[K, V]) ItemCount() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.items)
}

// Flush deletes all items from the cache.
func (c *cache[K, V]) Flush() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = map[K]Item[V]{}
}

func (c *Cache[K, V]) Stop() {
	c.janitor.Stop()
}

func (c *Cache[K, V]) Start(ctx context.Context) {
	c.janitor.Start(ctx)
}

func newCache[K comparable, V any](de time.Duration, m map[K]Item[V]) *cache[K, V] {
	if de == 0 {
		de = -1
	}
	c := &cache[K, V]{
		defaultExpiration: de,
		items:             m,
	}
	return c
}

func newCacheWithJanitor[K comparable, V any](
	ctx context.Context, de time.Duration, ci time.Duration, m map[K]Item[V], started bool,
) *Cache[K, V] {
	c := newCache[K, V](de, m)
	// This trick ensures that the janitor goroutine (which--granted it
	// was enabled--is running DeleteExpired on c forever) does not keep
	// the returned C object from being garbage collected. When it is
	// garbage collected, the finalizer stops the janitor goroutine, after
	// which c can be collected.

	janitorInstance := newJanitor[K, V](c, ci)
	if started {
		janitorInstance.Start(ctx)
	}

	return &Cache[K, V]{
		cache:   c,
		janitor: janitorInstance,
	}
}

// New return a new cache with a given default expiration duration and cleanup
// interval. If the expiration duration is less than one (or NoExpiration),
// the items in the cache never expire (by default), and must be deleted
// manually. If the cleanup interval is less than one, expired items are not
// deleted from the cache before calling c.DeleteExpired().
func New[K comparable, V any](
	defaultExpiration, cleanupInterval time.Duration,
) *Cache[K, V] {
	return newCacheWithJanitor[K, V](
		context.Background(),
		defaultExpiration,
		cleanupInterval,
		make(map[K]Item[V]),
		true,
	)
}

// NewInitialized works like a New, but it does not start the janitor.
// Use Start() to start the janitor.
func NewInitialized[K comparable, V any](
	ctx context.Context, defaultExpiration, cleanupInterval time.Duration,
) *Cache[K, V] {
	return newCacheWithJanitor[K, V](
		ctx,
		defaultExpiration,
		cleanupInterval,
		make(map[K]Item[V]),
		false,
	)
}

// NewFromItems returns a new cache with a given default expiration duration and cleanup
// interval. If the expiration duration is less than one (or NoExpiration),
// the items in the cache never expire (by default), and must be deleted
// manually. If the cleanup interval is less than one, expired items are not
// deleted from the cache before calling c.DeleteExpired().
//
// NewFromItems() also accepts an items map which will serve as the underlying map
// for the cache. This is useful for starting from a deserialized cache
// (serialized using e.g. gob.Encode() on c.ItemList()), or passing in e.g.
// make(map[string]Item, 500) to improve startup performance when the cache
// is expected to reach a certain minimum size.
//
// Only the cache's methods synchronize access to this map, so it is not
// recommended to keep any references to the map around after creating a cache.
// If need be, the map can be accessed at a later point using c.ItemList() (subject
// to the same caveat.)
//
// Note regarding serialization: When using e.g. gob, make sure to
// gob.Register() the individual types stored in the cache before encoding a
// map retrieved with c.ItemList(), and to register those same types before
// decoding a blob containing an items map.
func NewFromItems[K comparable, V any](
	defaultExpiration, cleanupInterval time.Duration, items map[K]Item[V],
) *Cache[K, V] {
	return newCacheWithJanitor[K, V](context.Background(), defaultExpiration, cleanupInterval, items, true)
}
