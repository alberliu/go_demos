package gepoll

// Handler Server 注册接口
type Handler interface {
	OnConnect(c *Conn)
	OnMessage(c *Conn, message interface{})
	OnError(c *Conn, err error)
	OnClose(c *Conn)
}
