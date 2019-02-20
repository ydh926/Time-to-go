package interfacex

type Person interface {
	Say() string
}

type Programer interface {
	Person
	Debug()
}

type Bob struct {
}

func (*Bob) Say() string {
	return "I am Bob"
}
func (*Bob) Debug() {
	println("I can debug")
}

func DoDebug(p Programer) {
	switch m := p.(type) {
	case *Bob:
		println("This is Bob")
		m.Debug()
	default:
		print("I don't konw who")
	}
}

func DoDebug2(p Programer) {
	switch p.(type) {
	case *Bob:
		println("This is Bob")
		p.Debug()
	default:
		print("I don't konw who")
	}
}
