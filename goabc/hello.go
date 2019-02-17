package main

import (
	"goabc/routine"
	"time"
)

func main(){
/*	array := make([]int,10)
	for i:= range array{
		array[i] = rand.Intn(100)
	}
	routine.Waitgroup.Add(1)
	routine.QuickSort02(array)
	routine.Waitgroup.Wait()
	printSlice(array)*/
	//testGo()
	routine.Process()
}


func printSlice(array []int){
	for _,v:=range array{
		println(v)
	}
}
