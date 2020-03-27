package gepoll

// Handler Server 注册接口
type Handler interface {
	OnConnect(fd int)
	OnMessage(fd int, message interface{})
	OnClose(fd int)
}
