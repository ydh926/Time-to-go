package main

import (
	"fmt"
	"net"
	"time"
)

func Process(){
	stop := time.NewTimer(30*time.Second)
	chr := make(chan string)
	chn := make(chan string)
	go readcin(chr)
	go readNet(chn)
	for{
		select {
		case <-stop.C:
			break
		case s,ok:=<-chr:
			if ok{
				println(s)
			}
		case s,ok:=<-chn:
			if ok{
				println(s)
			}
		}
	}
}

func readcin(ch chan string){
	var input string
	fmt.Scanln(&input)
	ch <- input
	close(ch)
}

func readNet(ch chan string){
	defer close(ch)
	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8080,
	})
	if err != nil {
		fmt.Println("监听失败!", err)
		return
	}
	defer socket.Close()

	data := make([]byte, 4096)
	_, _, err = socket.ReadFromUDP(data)
	if err != nil {
		fmt.Println("读取数据失败!", err)
	}
	ch <- string(data[:])
}

func main(){
	Process()
}