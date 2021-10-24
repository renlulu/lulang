package object

type Environment struct {
	store map[string]Obejct
}

func NewEnvironment() *Environment {
	s := make(map[string]Obejct)
	return &Environment{
		store: s,
	}
}

func (e *Environment) Get(name string) (Obejct, bool) {
	obj, ok := e.store[name]
	return obj, ok
}
func (e *Environment) Set(name string, val Obejct) Obejct {
	e.store[name] = val
	return val
}
