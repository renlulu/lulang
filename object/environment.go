package object

type Environment struct {
	store map[string]Obejct
	outer *Environment
}

func NewEnvironment() *Environment {
	s := make(map[string]Obejct)
	return &Environment{
		store: s,
	}
}

func (e *Environment) Get(name string) (Obejct, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}
func (e *Environment) Set(name string, val Obejct) Obejct {
	e.store[name] = val
	return val
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}
