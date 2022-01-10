package business

import "net"

type Packet struct {
	ConnectFlag int			//连接唯一标识
	Conn	* net.TCPConn	//连接句柄指针
	DataSize	int			//数据大小
	DataBuf []byte			//数据
}
