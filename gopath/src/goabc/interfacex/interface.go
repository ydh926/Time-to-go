package interfacex

type Person interface {
	Say() string
}

type Programer interface {
	Person
	Debug()
}

type Singer struct {
	Fans []string
}

func (*Singer) Sing() string {
	return "燃烧我的卡路里"
}

type Dancer struct{
	Fans []string
}

type Bob struct {
	Singer
	Dancer
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
