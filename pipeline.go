package bingo_events

type Pipeline struct {
	send    interface{} // 穿过管道的上下文
	through []Listener   // 中间件数组
	current int          // 当前执行到第几个中间件
}

func (p *Pipeline) Send(context interface{}) *Pipeline {
	p.send = context
	return p
}

func (p *Pipeline) Through(middlewares []Listener) *Pipeline {
	p.through = middlewares
	return p
}

func (p *Pipeline) Exec() {
	if len(p.through) > p.current {
		m := p.through[p.current]
		p.current += 1
		m(p.send, func(c interface{}) {
			p.Exec()
		})
	}

}

// 这里是路由的最后一站
func (p *Pipeline) Then(then func(context interface{})) {
	// 按照顺序执行
	// 将then作为最后一站的中间件
	var m Listener
	m = func(c interface{}, next func(c interface{})) {
		then(c)
		next(c)
	}
	p.through = append(p.through, m)
	p.Exec()
}
