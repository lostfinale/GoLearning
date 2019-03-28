package pool

import (
	"sync"
	"io"
	"errors"
	"log"
	"sync/atomic"
	"time"
	"math/rand"
)
// Pool 管理一组可以安全地在多个 goroutine 间共享的资源
// 被管理的资源必须实现 io.Closer 接口
type Pool struct {
	m sync.Mutex
	resources chan io.Closer
	factory func() (io.Closer, error)
	closed bool
}
// ErrPoolClosed 表示请求（Acquire）了一个已经关闭的池
var ErrPoolClosed = errors.New("Pool has been closed.")


// New 创建一个用来管理资源的池。
// 这个池需要一个可以分配新资源的函数，并规定池的大小
func New(fun func() (io.Closer, error), size uint) (*Pool, error) {

	if size <= 0 {
		return nil, errors.New("Size value too small.")
	}

	return &Pool{
		factory: fun,
		resources:make(chan io.Closer, size),
	}, nil
}

// Acquire 从池中获取一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <- p.resources:
		// 检查是否有空闲的资源
		log.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil

	default:
		// 因为没有空闲资源可用，所以提供一个新资源
		log.Println("Acquire:", "New Resource")
		return p.factory()
	}

}

// Release 将一个使用后的资源放回池里
func (p *Pool) Release(r io.Closer) {
	// 保证本操作和 Close 操作的安全
	p.m.Lock()

	defer p.m.Unlock()

	// 如果池已经被关闭，销毁这个资源
	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resources <- r:
		// 试图将这个资源放入队列
		log.Println("Release:", "In Queue")
	default:
		// 如果队列已满，则关闭这个资源
		log.Println("Release:", "Closing")
		r.Close()
	}

}

// Close 会让资源池停止工作，并关闭所有现有的资源
func (p *Pool) Close() {

	// 保证本操作与 Release 操作的安全
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	// 在清空通道里的资源之前，将通道关闭 如果不这样做，会发生死锁
	close(p.resources)

	// 关闭资源
	for r := range p.resources {
		r.Close()
	}
}

//---------------------测试代码------------------------------

const (
	maxGoroutines = 25// 要使用的 goroutine 的
	pooledResources = 2// 池中的资源的数量
)

//dbConnection 模拟要共享的资源
type dbConnection struct {
	ID int32
}
// Close 实现了 io.Closer 接口，以便 dbConnection可以被池管理。
// Close 用来完成任意资源的释放管理
func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

// idCounter 用来给每个连接分配一个独一无二的 id
var idCounter int32

// createConnection 是一个工厂函数
// 当需要一个新连接时，资源池会调用这个函数
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)
	return &dbConnection{id}, nil
}




// performQueries 用来测试连接的资源池
func performQueries(query int, p *Pool) {
	// 从池里请求一个连接
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	// 将该连接释放回池里
	defer p.Release(conn)

	// 用等待来模拟查询响应
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}

func Main() {

	var wg sync.WaitGroup
	wg.Add(maxGoroutines)
	// 创建用来管理连接的池
	p, err := New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}
	// 使用池里的连接来完成查询
	for query := 0; query < maxGoroutines; query++ {

		// 每个 goroutine 需要自己复制一份要查询值的副本，
		// 不然所有的查询会共享同一个查询变量
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()
	log.Println("Shutdown Program.")
	p.Close()
}



