package map_slice

func TestSlice() {
	array := make([]int, 0)  //初始化
	array = append(array, 0) //增加元素
	printSlice(array)
	array2 := []int{1, 2, 3}
	array = append(array, array2...) //合并切片
	printSlice(array)
	array3 := array[1:3] //截取切片
	printSlice(array3)
	array = append(array[:1], array[2:]...) //删除元素
	printSlice(array)
	array = make([]int,5)
	printSlice(array)
}
func printSlice(array []int) {
	for i := range array {
		print(array[i], ",")
	}
	println()
}
