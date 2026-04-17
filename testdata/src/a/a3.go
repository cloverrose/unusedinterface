package a

type closureParamName interface {
	Q()
}

type app struct {
	c closureParamName
}

func (a *app) doSomething() {
	a.c.Q()
}

func doSomething() {
	fn := func(closureParamName int) { // OK, parameter name is same as interface name but it is not a type name.
	}
	fn(42)
}
