### Type
- Locker
```
type Locker interface {
    Lock()
    Unlock()
} //Locker接口代表一个可以加锁和解锁的对象。
```
- Cond
```
type Cond struct {
    L Locker     // 锁的具体实现，通常为 mutex 或者rwmutex
}
```
### Cond
sync.Cond 实现了一个条件变量，在 Locker 的基础上增加了一个消息通知的功能，其内部维护了一个等待队列，队列中存放的是所有等待在这个 sync.Cond 的 Go 程，即保存了一个通知列表。
sync.Cond 可以用来唤醒一个或所有因等待条件变量而阻塞的 Go 程，以此来实现多个 Go 程间的同步
- NewCond
    - 使用锁L创建一个*Cond
- Brodcast()
    - Broadcast唤醒所有等待cond的线程。调用者在调用本方法时，建议（但并非必须）保持c.L的锁定。
- Signal()
    - Signal唤醒等待c的一个线程（如果存在）。调用者在调用本方法时，建议（但并非必须）保持c.L的锁定。
- Wait()
    - Wait自行解锁c.L并阻塞当前线程，在之后线程恢复执行时，Wait方法会在返回前锁定c.L
    - Wait除非被Broadcast或者Signal唤醒，不会主动返回。
    - 在调用 Signal，Broadcast 之前，应确保目标 Go 程进入 Wait 阻塞状态。
### Pool
Pool用于存储那些被分配了但是没有被使用，而未来可能会使用的值，以减小垃圾回收的压力。
- Get() 从池中获取一个interface{}类型的值
- Put() 把一个interface{}类型的值放置于池中
### Mutex
- Mutex 为互斥锁，Lock() 加锁，Unlock() 解锁
- 在一个 goroutine 获得 Mutex 后，其他 goroutine 只能等到这个 goroutine 释放该 Mutex
- 使用 Lock() 加锁后，不能再继续对其加锁，直到利用 Unlock() 解锁后才能再加锁，否则会导致死锁
- 在 Lock() 之前使用 Unlock() 会导致 panic 异常
- 已经锁定的 Mutex 并不与特定的 goroutine 相关联，允许一个 goroutine 对其加锁，再利用其他 goroutine 对其解锁
- 适用于读写不确定，并且只有一个读或者写的场景
### RWMutex
- RWMutex 是单写多读锁，该锁可以加一个写锁或者多个读锁
- 读锁占用的情况下会阻止写，不会阻止读，多个 goroutine 可以同时获取读锁
- 写锁会阻止其他 goroutine（无论读和写）进来，整个锁由该 goroutine 独占
- 已经锁定的 RWMutex 并不与特定的 goroutine 相关联，允许一个 goroutine 对其加锁，再利用其他 goroutine 对其解锁
- 适用于读多写少的场景
#### Lock()和Unlock()
- Lock() 加写锁，Unlock() 解写锁
- 如果在加写锁之前已经有其他的读锁和写锁，则 Lock() 会阻塞直到该锁可用，为确保该锁可用，已经阻塞的 Lock() 调用会从获得的锁中排除新的读取器，即写锁权限高于读锁，有写锁时优先进行写锁定
- 在 Lock() 之前使用 Unlock() 会导致 panic 错误
#### RLock()和RUnlock()
- RLock() 加读锁，RUnlock() 解读锁
- RLock() 加读锁时，如果存在写锁，则无法加读锁；当只有读锁或者没有锁时，可以加读锁，读锁可以加多个
- RUnlock() 解读锁，RUnlock() 撤销单独RLock() 调用，对于其他同时存在的读锁则没有效果
- 在没有读锁的情况下调用 RUnlock() 会导致 panic 错误
- RUnlock() 的个数不得多余 RLock()，否则会导致 panic 错误

### WaitGroup
- WaitGroup对象不是一个引用类型，在通过函数传值的时候需要使用地址：
- WaitGroup 对象内部有一个计数器，最初从0开始.
- 可以使用WaitGroup进行多个任务的同步，WaitGroup可以保证在并发环境中完成指定数量的任务
- WaitGroup 提供三种方法来控制计数器的数量
    - Add(n) 把计数器设置为n
    - Done() 每次把计数器-1
    - wait() 阻塞代码的运行，直到计数器的值减为0。
    <!--n可以为负数，但是如果counter如果有可能变成负数，需添加panic-->