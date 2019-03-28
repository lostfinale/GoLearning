package work

import (
	"sync"
	"log"
	"time"
	"github.com/goinaction/code/chapter7/patterns/work"
)

// Worker 必须满足接口类型，才能使用工作池
type Worker interface {
	Task()
}

//Pool 提供一个 goroutine 池，这个池可以完成任何已提交的 Worker 任务
type Pool struct {
	work chan Worker
	wg sync.WaitGroup
}

// New 创建一个新工作池
func New(maxGoroutines int) * Pool {
	p := Pool {
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)

	for i:=0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

// Run 提交工作到工作池
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// Shutdown 等待所有 goroutine 停止工作
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}


//-----------测试代码--------------------------------


// names 提供了一组用来显示的名字
var names = []string {
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}
// namePrinter 使用特定方式打印名字
type namePrinter struct {
	name string
}

// Task 实现 Worker 接口
func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func Main() {
	p := work.New(2)
	var wg sync.WaitGroup
	wg.Add(100 * len(names))
	for i:=0; i < 100; i++ {
		for _, name := range names {
			//创建一个 namePrinter 并提供指定的名字
			np := namePrinter{name: name,}
			go func (){
				//将任务提交执行。当 Run 返回时我们就知道任务已经处理完成
				p.Run(&np)
				wg.Done()
			}()


		}
	}

	wg.Wait()
	//让工作池停止工作，等待所有现有的工作完成
	p.Shutdown()
}