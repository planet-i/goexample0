package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var lockers = new(sync.Mutex)
var cond = sync.NewCond(lockers)

func test(x int) {
	cond.L.Lock() //获取锁
	fmt.Println("aaa: ", x)
	cond.Wait() //等待通知  暂时阻塞
	fmt.Println("bbb: ", x)
	time.Sleep(time.Second * 2)
	cond.L.Unlock() //释放锁
}

var bytePool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 1024)
		return &b
	},
}

func main() {
	UseCond()
	UsePool()
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	once.Do(onceBody)     //当且仅当第一次被调用时才执行函数f
	UseMutex()            //互斥锁
	UseRWMutex()          //读写锁
	var wg sync.WaitGroup // 声明一个等待组，一组等待任务只需要一个等待组
	FetchURL(&wg)
}

func UseCond() {
	for i := 0; i < 5; i++ {
		go test(i)
	}
	fmt.Println("start all")
	time.Sleep(time.Second * 1)
	fmt.Println("signal")
	cond.Signal()

	time.Sleep(time.Second * 1)
	fmt.Println("signal")
	cond.Signal() //下发一个通知给已经获取锁的goroutine

	fmt.Println("broadcast")
	time.Sleep(time.Second * 3)
	cond.Broadcast() //发广播给所有等待的goroutine
	time.Sleep(time.Second * 5)
	fmt.Println("finish all")
}
func UsePool() {
	a := time.Now()
	// 使用对象池
	for i := 0; i < 1000000000; i++ {
		obj := bytePool.Get().(*[]byte)
		_ = obj
		bytePool.Put(obj)
	}
	b := time.Now()
	fmt.Println("with pool %v", b.Sub(a))
}
func UseMutex() { //Lock()和UnLock()成对使用
	var mutex sync.Mutex
	fmt.Println("Lock the lock")
	mutex.Lock()
	fmt.Println("The lock is locked")
	channels := make([]chan int, 4)
	for i := 0; i < 4; i++ {
		channels[i] = make(chan int)
		go func(i int, c chan int) {
			fmt.Println("Not lock: ", i)
			mutex.Lock()
			fmt.Println("Locked: ", i)
			time.Sleep(time.Second)
			fmt.Println("Unlock the lock: ", i)
			mutex.Unlock()
			c <- i
		}(i, channels[i])
	}
	time.Sleep(time.Second)
	fmt.Println("Unlock the lock")
	mutex.Unlock()
	time.Sleep(time.Second)

	for _, c := range channels {
		println(<-c)
	}
}
func UseRWMutex() { //Lock()和UnLock()成对使用
	var mutex *sync.RWMutex
	mutex = new(sync.RWMutex)
	fmt.Println("Lock the lock")
	mutex.Lock()
	fmt.Println("The lock is locked")
	channels := make([]chan int, 4)
	for i := 0; i < 4; i++ {
		channels[i] = make(chan int)
		go func(i int, c chan int) {
			fmt.Println("Not read lock: ", i)
			mutex.RLock()
			fmt.Println("Read Locked: ", i)
			fmt.Println("Unlock the read lock: ", i)
			time.Sleep(time.Second)
			mutex.RUnlock()
			c <- i
		}(i, channels[i])
	}
	time.Sleep(time.Second)
	fmt.Println("Unlock the lock")
	mutex.Unlock()
	time.Sleep(time.Second)

	for _, c := range channels {
		println(<-c)
	}
}
func FetchURL(wg *sync.WaitGroup) {
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	} //可访问的网站地址的字符串切片。
	for _, url := range urls {
		wg.Add(1) //将等待组的计数器值加1
		go func(url string) {
			defer wg.Done()         //函数完成时将等待组值减1
			_, err := http.Get(url) //Get() 函数会一直阻塞直到网站响应或者超时。
			fmt.Println(url, err)
		}(url)
	}
	wg.Wait()           //，等待所有的网站都响应或者超时后，任务完成，Wait 就会停止阻塞。
	fmt.Println("over") //如果不使用WaitGroup，将会直接输出over,线程任务未完成
}
