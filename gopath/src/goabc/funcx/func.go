package funcx

func TestFunc(args []string)(int, string){
	return 1,"1"
}

func DoFunc(f func([]string)(int, string)) string {
	_,v := f([]string{})
	return v
}