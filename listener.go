package bingo_events

// 监听器接口

type Listener func(event interface{},next func(event interface{}))
