package injector

import (
	"github.com/soverenio/vanilla/throw"
)

func NewStaticContainer(parentRegistry DependencyRegistry, contentRegistry ScanDependencyRegistry) StaticContainer {
	sc := StaticContainer{parentRegistry: parentRegistry}
	if contentRegistry != nil {
		sc.localRegistry = map[string]interface{}{}
		contentRegistry.ScanDependencies(func(id string, v interface{}) bool {
			sc.localRegistry[id] = v
			return false
		})
	}
	return sc
}

type StaticContainer struct {
	parentRegistry DependencyRegistry
	localRegistry  map[string]interface{}
}

func (m StaticContainer) FindDependency(id string) (interface{}, bool) {
	if v, ok := m.localRegistry[id]; ok {
		return v, true
	}
	if m.parentRegistry != nil {
		return m.parentRegistry.FindDependency(id)
	}
	return nil, false
}

func (m StaticContainer) ScanDependencies(fn func(id string, v interface{}) bool) (found bool) {
	if fn == nil {
		panic(throw.IllegalValue())
	}

	for key, value := range m.localRegistry {
		if fn(key, value) {
			return true
		}
	}

	if sp, ok := m.parentRegistry.(ScanDependencyRegistry); ok {
		return sp.ScanDependencies(fn)
	}
	return false
}

func (m StaticContainer) AsRegistry() DependencyRegistry {
	if m.parentRegistry == nil && len(m.localRegistry) == 0 {
		return nil
	}
	return m
}
