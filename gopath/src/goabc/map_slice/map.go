package map_slice

func TestMap(){
	Roster := make(map[string]string,0) //初始化
	//Roster2 := map[string]int{"bob":1, "bob2":2}
	Roster["bob"]="01"         //插入数据
	printBob(Roster)
	Roster["bob"]="02"         //修改数据
	printBob(Roster)
	delete(Roster, "bob") //删除数据
	printBob(Roster)
}
func printBob(Roster map[string]string){
	value,ok := Roster["bob"] //查询数据
	if ok{
		println("bob is "+value)
	}else{
		println("There is no bob")
	}
}