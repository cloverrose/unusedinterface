package a

type c1 interface { // OK
	C()
}

type D1 interface { // want "interface D1 is defined but not used within the same package"
	D()
}
