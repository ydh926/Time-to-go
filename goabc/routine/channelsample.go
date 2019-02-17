package routine

import "time"

func testGo(){
		req := make(chan int)
		go doPlus(1,1,req)
		v,ok := <-req
		if ok{
			println("channel is opening")
		}else{
			println("channel is closed")
		}
		println(v)
	}
	func doPlus(a,b int, req chan int){
		time.Sleep(10000)
		req <- a+b
		close(req)
}
