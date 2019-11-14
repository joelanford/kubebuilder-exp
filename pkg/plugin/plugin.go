package plugin

import (
	"fmt"
	"sort"
	"sync"
)

type Plugin interface {
	Version() string
}

var (
	plugins map[string]Plugin
	m       sync.RWMutex
)

func init() {
	plugins = make(map[string]Plugin)
}

func Register(s Plugin) error {
	m.Lock()
	defer m.Unlock()
	v := s.Version()
	if _, ok := plugins[v]; ok {
		return fmt.Errorf("scaffolder already registered for version %q", v)
	}
	plugins[v] = s
	return nil
}

func Lookup(v string) (Plugin, bool) {
	m.RLock()
	defer m.RUnlock()
	s, ok := plugins[v]
	return s, ok
}

func List() []Plugin {
	m.RLock()
	defer m.RUnlock()

	var keys []string
	for k := range plugins {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var list []Plugin
	for _, k := range keys {
		list = append(list, plugins[k])
	}
	return list
}
