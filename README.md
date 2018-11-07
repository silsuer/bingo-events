# bingo-events
bingo框架的事件模块

## 简介
   
   `bingo-events` 是 [bingo](https://github.com/silsuer/bingo) 框架的基础子模块，很多子模块都依赖本模块
   
   `bingo-events` 实现了一个简单的观察者模式，`bingo`中的数据分发均基于事件。
   
## 优势
  
   基于事件的编程模式主要用于业务逻辑的解耦，当事件发生后，通知所有的观察者进行下一步的操作。

## 使用方式

   1. 使用 `NewApp()` 来创建一个应用对象
    
   2. 使用 `App.Listen()` 方法来添加一个监听器，或在实现 `IEvent` 接口的结构体上使用 `Attach()` 方法来添加一个观察者
   
   3. 使用 `App.Dispatch()` 来进行事件分发，将会触发所有匹配该事件的观察者（监听器）

## 示例
   
   ```go
         // 创建一个结构体，实现事件接口
        type listen struct {
        	bingo_events.Event
        	Name string
        }
  
        func main() {
        	// 事件对象
        	app := bingo_events.NewApp()
        	// 添加观察者
        	app.Listen("*main.listen", ListenStruct)  // 直接使用 Listen 方法，为监听的结构体添加一个回调 
        	app.Listen("*main.listen", L2)
        	l := new(listen)  // 创建刚刚那个实现 IEvent 接口的对象
        
        	// 添加观察者
        	l.Attach(ListenStruct)   // 使用 IEvent 的Attach方法为该对象添加观察者
        
        	l.Name = "silsuer"  // 为对象属性赋值
  	
        	// 复制完毕，开始分发事件，从上可知，共添加了两个观察者： ListenStruct 和 L2
  	        // 会按照监听的顺序执行，由于最开始已经添加了 ListenStruct 监听器，所以第二次再次添加的时候不会重复添加
            // 此处的分发，就是将参数顺序传入每一个监听器，进行后续操作
  	        app.Dispatch(l)
        }
  
        
        func ListenStruct(event interface{}, next func(event interface{})) {
        	// 由于监听器可监听的对象不一定非要实现 IEvent 接口，所以这里需要使用类型断言，将对象转换回原本的类型
  	        a := event.(*listen)
        	fmt.Println(a.Name) // output: silsuer
        	a.Name = "god"   // 更改结构体的属性
        	next(a)   // 调用next函数继续走下一个监听器，如果此处不调用，程序将会终止在此处，不会继续往下执行
        }
        
        func L2(event interface{}, next func(event interface{})) {      
  	       next(event)  // 当在函数中间调用next方法时，将会先调用其他的监听器，最后在继续往下走
  	       fmt.Println(event.(*listen).Name)
        }

   ```
 

