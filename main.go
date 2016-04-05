package main

import (
	"restful/controller"
)

func main() {
	//init config
	//init log
	controller.Init()
	// 最后添加个系统信号处理，阻塞接收kill命令什么的，为了程序不退出
}
