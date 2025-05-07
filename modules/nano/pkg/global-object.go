package nano

import "sync"

type globalObject struct {
	innerData map[string]interface{}
	mu        sync.RWMutex
}

func (g *globalObject) Define(name string, value interface{}) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.innerData[name] = value
}

func (g *globalObject) UnDefine(name string) {
	g.mu.Lock()
	defer g.mu.Unlock()
	delete(g.innerData, name)
}

func (g *globalObject) Get(name string) interface{} {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.innerData[name]
}

func newGlobalObject() *globalObject {
	obj := &globalObject{innerData: make(map[string]interface{})}

	return obj
}
