package injector

import (
	"github.com/soverenio/vanilla/throw"
)

func NewMultiMapRegistry(maps []map[string]interface{}) MultiMapRegistry {
	return MultiMapRegistry{maps}
}

var _ DependencyRegistry = MultiMapRegistry{}
var _ ScanDependencyRegistry = MultiMapRegistry{}

type MultiMapRegistry struct {
	maps []map[string]interface{}
}

func (v MultiMapRegistry) ScanDependencies(fn func(id string, v interface{}) bool) bool {
	if fn == nil {
		panic(throw.IllegalValue())
	}
	for _, om := range v.maps {
		for id, v := range om {
			if fn(id, v) {
				return true
			}
		}
	}
	return false
}

func (v MultiMapRegistry) FindDependency(id string) (interface{}, bool) {
	for _, om := range v.maps {
		if v, ok := om[id]; ok {
			return v, true
		}
	}
	return nil, false
}
