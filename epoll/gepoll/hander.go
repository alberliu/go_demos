package gepoll

// Handler Server 注册接口
type Handler interface {
	OnConnect(fd int)
	OnMessage(fd int, message interface{})
	OnError(fd int, err error)
	OnClose(fd int)
}
