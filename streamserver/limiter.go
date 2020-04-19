package streamserver

//设置connLimitNum 代表支持同一时间最大连接数
//最大连接数使用connchan来限制
type ConnLimiter struct {
	connChan  chan uint
	connLimitNum   uint
}

func NewConnLimiter(num uint)ConnLimiter{
	return ConnLimiter{
		connChan:     make(chan uint,num),
		connLimitNum: 20,
	}
}
//增加一个链接
func (c *ConnLimiter)AddConn() bool{
	if int(c.connLimitNum) <= len(c.connChan){
		return  false
	}
	 c.connChan <- 1
	 return true
}

//丢弃一个链接
func (c *ConnLimiter)ReduceConn(){
	<-c.connChan
}
