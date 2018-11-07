package bingo_events

import (
	"sync"
	"reflect"
)

// 服务容器
type App struct {
	//inject.Graph
	sync.RWMutex
	events map[string][]Listener // 事件，允许正则匹配
}

// 容器的目的是为了创建事件之后，为事件注入观察者，所以为了方便，直接全部解耦
// 创建

// 新建一个App，可以在其中注册事件
func NewApp() *App {
	app := &App{}
	app.events = make(map[string][]Listener)
	return app
}

// 给容器绑定事件,传入 类型-> 观察者的绑定
// 可以使用反射，传入类型的字符串
// 否则bind时也可以放置包括通配符的
func (app *App) Bind(t string, listeners []Listener) {
	for k := range listeners {
		app.Listen(t, listeners[k])
	}
}

// 监听[事件][监听器]，单独绑定一个监听器
func (app *App) Listen(str string, listener Listener) {
	app.Lock()
	app.events[str] = append(app.events[str], listener)
	app.Unlock()
}

// 分发事件，传入各种事件，如果是
func (app *App) Dispatch(events ...interface{}) {
	// 容器分发数据
	var event string
	for k := range events {
		if _, ok := events[k].(string); ok { // 如果传入的是字符串类型的
			event = events[k].(string)
		} else {
			// 不是字符串类型的，那么得到其类型
			event = reflect.TypeOf(events[k]).String()
		}

		// 如果实现了 事件接口 IEvent，则调用事件的观察者模式，得到所有的
		var observers []Listener
		if _, ok := events[k].(IEvent); ok {
			obs := events[k].(IEvent).Observers()
			observers = append(observers, obs...) // 将事件中自行添加的观察者，放在所有观察者之后
		}

		if obs, exist := app.events[event]; exist {
			observers = append(observers, obs...)
		}

		if len(observers) > 0 {
			// 得到了所有的观察者，这里通过pipeline来执行，通过next来控制什么时候调用这个观察者
			new(Pipeline).Send(events[k]).Through(observers).Then(func(context interface{}) {

			})
		}

	}

}
