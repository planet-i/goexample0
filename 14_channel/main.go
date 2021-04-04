package main

func main() {
	done := make(chan struct{}) //结束事件  创建通道，因为仅是通知，数据并没有实际意义
	c := make(chan string)      //数据传输通道

	go func() {
		s := <-c //s从通道接收消息
		println(s)
		close(done) //关闭通道，发出信号
	}()

	c <- "hi" // 发送消息
	<-done    //阻塞，直到有数据或者通道关闭
}
