package nano

type globalObject struct {
	innerData map[string]interface{}
}

func (g *globalObject) Define(name string, value interface{}) {
	g.innerData[name] = value
}

func (g *globalObject) Get(name string) interface{} {
	return g.innerData[name]
}

func newGlobalObject() *globalObject {
	obj := &globalObject{innerData: make(map[string]interface{})}

	return obj
}
