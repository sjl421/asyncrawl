package main

import "./base"


func test()  {
	ch := make(chan string)
	rooms := base.Pindao(10)
	for _, i := range rooms{
		go base.Connect(i)
	}
	for i:=1;i<len(rooms) ;i++  {
		<-ch
	}
}
// todo fix bug。


func main()  {
	test()
}
