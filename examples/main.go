package main

import (
	"github.com/silsuer/bingo-events"
	"fmt"
)

//var m map[bingo_events.Event]bingo_events.Listener

type listen struct {
	//bingo_events.Event
	Name string
}

func main() {
	// 事件对象
	app := bingo_events.NewApp()
	// 添加观察者
	app.Listen("*main.listen", ListenStruct)
	app.Listen("*main.listen", L2)
	l := new(listen)

	// 添加观察者
	//l.Attach(ListenStruct)

	l.Name = "silsuer"
	app.Dispatch(l)
}

func ListenStruct(event interface{}, next func(event interface{})) {
	a := event.(*listen)
	fmt.Println(a.Name)
	a.Name = "god"
	next(a)
}

func L2(event interface{}, next func(event interface{})) {
	fmt.Println(event.(*listen).Name)
	next(event)
}
