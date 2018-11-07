package bingo_events

// 事件接口
type IEvent interface {
	Attach(listeners ...Listener) // 事件添加观察者
	Observers() []Listener        // 获取所有的观察者
	//Detach(listener Listener)
	DetachIndexOf(index int)
}

type Event struct {
	listeners []Listener
}

func (e *Event) Attach(listeners ...Listener) {
	e.listeners = append(e.listeners, listeners...)
}

//func (e *Event) Detach(listener Listener) {
//	for k := range e.listeners {
//		if listener == e.listeners[k] {
//
//		}
//	}
//}

// 移除某个监听器
func (e *Event) DetachIndexOf(index int) {
	if len(e.listeners) > index && index > 0 {
		e.listeners = append(e.listeners[:index], e.listeners[index+1:]...)
	}
}

func (e *Event) Observers() []Listener {
	return e.listeners
}
