# mango_message

做了一个简单的实例
运行方式：
启动两个cmd
分别运行
	go run main.go message.go boltdb.go 0 tcp://localhost:45454
	go run main.go message.go boltdb.go 1 tcp://localhost:45454
  
  go run main.go pubsub.go boltdb.go server tcp://localhost:45454
  go run main.go pubsub.go boltdb.go client tcp://localhost:45454
  
#mango 包作用说明  
pair 消息一对一传送(实现)
subpub 将消息分发给订阅消息的用户(实现)
bus 简单的多对多通信
reqrep 允许构建无状态服务集群来处理用户请求
pipeline  汇总来自多个来源的消息，并在许多目的点之间负载平衡
survey 允许一次查询多个应用状态

ipc 在单个机器上的进程间传输
tcp通过tcp的网络传输
