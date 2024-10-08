package cryptkit

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// KeyStoreFactoryMock implements KeyStoreFactory
type KeyStoreFactoryMock struct {
	t minimock.Tester

	funcCreatePublicKeyStore          func(s1 SigningKeyHolder) (p1 PublicKeyStore)
	inspectFuncCreatePublicKeyStore   func(s1 SigningKeyHolder)
	afterCreatePublicKeyStoreCounter  uint64
	beforeCreatePublicKeyStoreCounter uint64
	CreatePublicKeyStoreMock          mKeyStoreFactoryMockCreatePublicKeyStore
}

// NewKeyStoreFactoryMock returns a mock for KeyStoreFactory
func NewKeyStoreFactoryMock(t minimock.Tester) *KeyStoreFactoryMock {
	m := &KeyStoreFactoryMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreatePublicKeyStoreMock = mKeyStoreFactoryMockCreatePublicKeyStore{mock: m}
	m.CreatePublicKeyStoreMock.callArgs = []*KeyStoreFactoryMockCreatePublicKeyStoreParams{}

	return m
}

type mKeyStoreFactoryMockCreatePublicKeyStore struct {
	mock               *KeyStoreFactoryMock
	defaultExpectation *KeyStoreFactoryMockCreatePublicKeyStoreExpectation
	expectations       []*KeyStoreFactoryMockCreatePublicKeyStoreExpectation

	callArgs []*KeyStoreFactoryMockCreatePublicKeyStoreParams
	mutex    sync.RWMutex
}

// KeyStoreFactoryMockCreatePublicKeyStoreExpectation specifies expectation struct of the KeyStoreFactory.CreatePublicKeyStore
type KeyStoreFactoryMockCreatePublicKeyStoreExpectation struct {
	mock    *KeyStoreFactoryMock
	params  *KeyStoreFactoryMockCreatePublicKeyStoreParams
	results *KeyStoreFactoryMockCreatePublicKeyStoreResults
	Counter uint64
}

// KeyStoreFactoryMockCreatePublicKeyStoreParams contains parameters of the KeyStoreFactory.CreatePublicKeyStore
type KeyStoreFactoryMockCreatePublicKeyStoreParams struct {
	s1 SigningKeyHolder
}

// KeyStoreFactoryMockCreatePublicKeyStoreResults contains results of the KeyStoreFactory.CreatePublicKeyStore
type KeyStoreFactoryMockCreatePublicKeyStoreResults struct {
	p1 PublicKeyStore
}

// Expect sets up expected params for KeyStoreFactory.CreatePublicKeyStore
func (mmCreatePublicKeyStore *mKeyStoreFactoryMockCreatePublicKeyStore) Expect(s1 SigningKeyHolder) *mKeyStoreFactoryMockCreatePublicKeyStore {
	if mmCreatePublicKeyStore.mock.funcCreatePublicKeyStore != nil {
		mmCreatePublicKeyStore.mock.t.Fatalf("KeyStoreFactoryMock.CreatePublicKeyStore mock is already set by Set")
	}

	if mmCreatePublicKeyStore.defaultExpectation == nil {
		mmCreatePublicKeyStore.defaultExpectation = &KeyStoreFactoryMockCreatePublicKeyStoreExpectation{}
	}

	mmCreatePublicKeyStore.defaultExpectation.params = &KeyStoreFactoryMockCreatePublicKeyStoreParams{s1}
	for _, e := range mmCreatePublicKeyStore.expectations {
		if minimock.Equal(e.params, mmCreatePublicKeyStore.defaultExpectation.params) {
			mmCreatePublicKeyStore.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreatePublicKeyStore.defaultExpectation.params)
		}
	}

	return mmCreatePublicKeyStore
}

// Inspect accepts an inspector function that has same arguments as the KeyStoreFactory.CreatePublicKeyStore
func (mmCreatePublicKeyStore *mKeyStoreFactoryMockCreatePublicKeyStore) Inspect(f func(s1 SigningKeyHolder)) *mKeyStoreFactoryMockCreatePublicKeyStore {
	if mmCreatePublicKeyStore.mock.inspectFuncCreatePublicKeyStore != nil {
		mmCreatePublicKeyStore.mock.t.Fatalf("Inspect function is already set for KeyStoreFactoryMock.CreatePublicKeyStore")
	}

	mmCreatePublicKeyStore.mock.inspectFuncCreatePublicKeyStore = f

	return mmCreatePublicKeyStore
}

// Return sets up results that will be returned by KeyStoreFactory.CreatePublicKeyStore
func (mmCreatePublicKeyStore *mKeyStoreFactoryMockCreatePublicKeyStore) Return(p1 PublicKeyStore) *KeyStoreFactoryMock {
	if mmCreatePublicKeyStore.mock.funcCreatePublicKeyStore != nil {
		mmCreatePublicKeyStore.mock.t.Fatalf("KeyStoreFactoryMock.CreatePublicKeyStore mock is already set by Set")
	}

	if mmCreatePublicKeyStore.defaultExpectation == nil {
		mmCreatePublicKeyStore.defaultExpectation = &KeyStoreFactoryMockCreatePublicKeyStoreExpectation{mock: mmCreatePublicKeyStore.mock}
	}
	mmCreatePublicKeyStore.defaultExpectation.results = &KeyStoreFactoryMockCreatePublicKeyStoreResults{p1}
	return mmCreatePublicKeyStore.mock
}

// Set uses given function f to mock the KeyStoreFactory.CreatePublicKeyStore method
func (mmCreatePublicKeyStore *mKeyStoreFactoryMockCreatePublicKeyStore) Set(f func(s1 SigningKeyHolder) (p1 PublicKeyStore)) *KeyStoreFactoryMock {
	if mmCreatePublicKeyStore.defaultExpectation != nil {
		mmCreatePublicKeyStore.mock.t.Fatalf("Default expectation is already set for the KeyStoreFactory.CreatePublicKeyStore method")
	}

	if len(mmCreatePublicKeyStore.expectations) > 0 {
		mmCreatePublicKeyStore.mock.t.Fatalf("Some expectations are already set for the KeyStoreFactory.CreatePublicKeyStore method")
	}

	mmCreatePublicKeyStore.mock.funcCreatePublicKeyStore = f
	return mmCreatePublicKeyStore.mock
}

// When sets expectation for the KeyStoreFactory.CreatePublicKeyStore which will trigger the result defined by the following
// Then helper
func (mmCreatePublicKeyStore *mKeyStoreFactoryMockCreatePublicKeyStore) When(s1 SigningKeyHolder) *KeyStoreFactoryMockCreatePublicKeyStoreExpectation {
	if mmCreatePublicKeyStore.mock.funcCreatePublicKeyStore != nil {
		mmCreatePublicKeyStore.mock.t.Fatalf("KeyStoreFactoryMock.CreatePublicKeyStore mock is already set by Set")
	}

	expectation := &KeyStoreFactoryMockCreatePublicKeyStoreExpectation{
		mock:   mmCreatePublicKeyStore.mock,
		params: &KeyStoreFactoryMockCreatePublicKeyStoreParams{s1},
	}
	mmCreatePublicKeyStore.expectations = append(mmCreatePublicKeyStore.expectations, expectation)
	return expectation
}

// Then sets up KeyStoreFactory.CreatePublicKeyStore return parameters for the expectation previously defined by the When method
func (e *KeyStoreFactoryMockCreatePublicKeyStoreExpectation) Then(p1 PublicKeyStore) *KeyStoreFactoryMock {
	e.results = &KeyStoreFactoryMockCreatePublicKeyStoreResults{p1}
	return e.mock
}

// CreatePublicKeyStore implements KeyStoreFactory
func (mmCreatePublicKeyStore *KeyStoreFactoryMock) CreatePublicKeyStore(s1 SigningKeyHolder) (p1 PublicKeyStore) {
	mm_atomic.AddUint64(&mmCreatePublicKeyStore.beforeCreatePublicKeyStoreCounter, 1)
	defer mm_atomic.AddUint64(&mmCreatePublicKeyStore.afterCreatePublicKeyStoreCounter, 1)

	if mmCreatePublicKeyStore.inspectFuncCreatePublicKeyStore != nil {
		mmCreatePublicKeyStore.inspectFuncCreatePublicKeyStore(s1)
	}

	mm_params := &KeyStoreFactoryMockCreatePublicKeyStoreParams{s1}

	// Record call args
	mmCreatePublicKeyStore.CreatePublicKeyStoreMock.mutex.Lock()
	mmCreatePublicKeyStore.CreatePublicKeyStoreMock.callArgs = append(mmCreatePublicKeyStore.CreatePublicKeyStoreMock.callArgs, mm_params)
	mmCreatePublicKeyStore.CreatePublicKeyStoreMock.mutex.Unlock()

	for _, e := range mmCreatePublicKeyStore.CreatePublicKeyStoreMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.p1
		}
	}

	if mmCreatePublicKeyStore.CreatePublicKeyStoreMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreatePublicKeyStore.CreatePublicKeyStoreMock.defaultExpectation.Counter, 1)
		mm_want := mmCreatePublicKeyStore.CreatePublicKeyStoreMock.defaultExpectation.params
		mm_got := KeyStoreFactoryMockCreatePublicKeyStoreParams{s1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreatePublicKeyStore.t.Errorf("KeyStoreFactoryMock.CreatePublicKeyStore got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreatePublicKeyStore.CreatePublicKeyStoreMock.defaultExpectation.results
		if mm_results == nil {
			mmCreatePublicKeyStore.t.Fatal("No results are set for the KeyStoreFactoryMock.CreatePublicKeyStore")
		}
		return (*mm_results).p1
	}
	if mmCreatePublicKeyStore.funcCreatePublicKeyStore != nil {
		return mmCreatePublicKeyStore.funcCreatePublicKeyStore(s1)
	}
	mmCreatePublicKeyStore.t.Fatalf("Unexpected call to KeyStoreFactoryMock.CreatePublicKeyStore. %v", s1)
	return
}

// CreatePublicKeyStoreAfterCounter returns a count of finished KeyStoreFactoryMock.CreatePublicKeyStore invocations
func (mmCreatePublicKeyStore *KeyStoreFactoryMock) CreatePublicKeyStoreAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreatePublicKeyStore.afterCreatePublicKeyStoreCounter)
}

// CreatePublicKeyStoreBeforeCounter returns a count of KeyStoreFactoryMock.CreatePublicKeyStore invocations
func (mmCreatePublicKeyStore *KeyStoreFactoryMock) CreatePublicKeyStoreBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreatePublicKeyStore.beforeCreatePublicKeyStoreCounter)
}

// Calls returns a list of arguments used in each call to KeyStoreFactoryMock.CreatePublicKeyStore.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreatePublicKeyStore *mKeyStoreFactoryMockCreatePublicKeyStore) Calls() []*KeyStoreFactoryMockCreatePublicKeyStoreParams {
	mmCreatePublicKeyStore.mutex.RLock()

	argCopy := make([]*KeyStoreFactoryMockCreatePublicKeyStoreParams, len(mmCreatePublicKeyStore.callArgs))
	copy(argCopy, mmCreatePublicKeyStore.callArgs)

	mmCreatePublicKeyStore.mutex.RUnlock()

	return argCopy
}

// MinimockCreatePublicKeyStoreDone returns true if the count of the CreatePublicKeyStore invocations corresponds
// the number of defined expectations
func (m *KeyStoreFactoryMock) MinimockCreatePublicKeyStoreDone() bool {
	for _, e := range m.CreatePublicKeyStoreMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreatePublicKeyStoreMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreatePublicKeyStoreCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreatePublicKeyStore != nil && mm_atomic.LoadUint64(&m.afterCreatePublicKeyStoreCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreatePublicKeyStoreInspect logs each unmet expectation
func (m *KeyStoreFactoryMock) MinimockCreatePublicKeyStoreInspect() {
	for _, e := range m.CreatePublicKeyStoreMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to KeyStoreFactoryMock.CreatePublicKeyStore with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreatePublicKeyStoreMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreatePublicKeyStoreCounter) < 1 {
		if m.CreatePublicKeyStoreMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to KeyStoreFactoryMock.CreatePublicKeyStore")
		} else {
			m.t.Errorf("Expected call to KeyStoreFactoryMock.CreatePublicKeyStore with params: %#v", *m.CreatePublicKeyStoreMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreatePublicKeyStore != nil && mm_atomic.LoadUint64(&m.afterCreatePublicKeyStoreCounter) < 1 {
		m.t.Error("Expected call to KeyStoreFactoryMock.CreatePublicKeyStore")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *KeyStoreFactoryMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCreatePublicKeyStoreInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *KeyStoreFactoryMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *KeyStoreFactoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreatePublicKeyStoreDone()
}
