package comm

import (
	"../business"
	"fmt"
	"net"
)

type  NewTaskJob func(packet business.Packet) business.Job

type Server struct {
	//服务器名称
	Name string
	//服务器ip版本
	IPVersion string
	//服务器监听ip
	IP string
	//端口
	Port int
	//工作池
	Pool	* business.WorkerPool
	//消息接收器类型()
	NewJob NewTaskJob
}

func (this *Server) Start() {

	fmt.Printf("[Start] string Listenner at IP: %s,Prot %d,is start \n",this.IP,this.Port)

	//异步操作，不阻塞
	go func() {
		//获取一个TCP的Addr
		addr,err := net.ResolveTCPAddr(this.IPVersion,fmt.Sprintf("%s:%d",this.IP,this.Port))
		if err!=nil{
			fmt.Println("resolve tcp addr error:",err)
		}
		//监听路由的地址
		listen,err := net.ListenTCP(this.IPVersion,addr)
		if err!=nil{
			fmt.Println("Listen",this.IPVersion,"err",err)
			return
		}

		fmt.Println("start zinx server succ",this.Name,"succ,Listening..")

		//阻塞的等待客户端的链接，处理业务
		for {
			conn,err := listen.AcceptTCP()
			if err!=nil {
				fmt.Println("Accept err",err)
				continue
			}

			//客户端已经建立链接,做一些业务，因为比较简单，所以就回写吧
			go func() {
				for  {
					buf := make([]byte,512)
					cnt,err := conn.Read(buf)
					if err!= nil{
						fmt.Println("recv buf err",err)
						conn.Close()
						break
					}
					packet := business.Packet{ConnectFlag: 1,Conn:conn,DataBuf:buf,DataSize:cnt}
					job:=this.NewJob(packet)
					this.Pool.JobQueue <- job
				}
			}()
		}
	}()

}

func (this *Server) Stop() {
	//TODO 回收
}

func (this *Server) Serve() {
	//启动服务
	this.Start()

	//TODO 启动服务后做一些额外的业务功能

	//阻塞状态,因为Start()是异步的，如果不加主程式早就运行完了
	select {

	}
}

/**
  创建服务
*/
func NewServer(name string,newjob NewTaskJob) IServer  {
	num := 100 * 100 * 20
	// 注册工作池，传入任务
	// 参数1 worker并发个数
	p := business.NewWorkerPool(num)
	p.Run()

	s:= &Server{
		Name:name,
		IPVersion:"tcp4",
		IP:"0.0.0.0",
		Port:8999,
		Pool:p,
		NewJob:newjob,
	}
	return  s
}

