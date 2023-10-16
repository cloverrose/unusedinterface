package a

type a1 interface { // want "interface a1 is defined but not used within the same package"
	A()
}

type b1 interface { // OK
	B()
}

func m(b b1) {

}

func n(c c1) {

}
