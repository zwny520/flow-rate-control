package main

import (
	"../framework/business"
	"../framework/comm"
	"../proto"
	"fmt"
	"github.com/golang/protobuf/proto"
	"time"
)

type BusinessJob struct {
	Packet business.Packet
}

func (s * BusinessJob) Do() {
	fmt.Println(time.Now().Format("15:04:05")+" Recv client request")
	s.Packet.Conn.Write([]byte( " server send respone!"))
	//处理业务
}

func NewBusinessJob(packet business.Packet) business.Job {
	business_job := BusinessJob{}
	business_job.Packet=packet
	return &business_job
}

func main() {

	min_data:=pb.MinData{}
	min_data.Offset		=	1
	min_data.Moment		=	935
	min_data.Price		=	9.89
	min_data.AvePrice	=	9.98
	min_data.Volume		=	300000
	min_data.Open		=	9.55
	min_data.High		=	10.23
	min_data.Low		=	9.47
	min_data.Close		=	9.99
	min_data.IsWrite	=	1

	//序列化
	b, _ := proto.Marshal(&min_data)

	//反序列化
	var min_data1 pb.MinData
	err := proto.Unmarshal(b, &min_data1)
	if err != nil {
		return
	}

	fmt.Printf("Offset:%d,Moment:%d,Price:%f,AvePrice:%f,Volume:%d,Open:%f,High:%f,Low:%f,Close:%f",
		min_data1.GetOffset(),min_data1.GetMoment(),min_data1.GetPrice(),min_data1.GetAvePrice(),
		min_data1.GetVolume(),min_data1.GetOpen(),min_data1.GetHigh(),min_data1.GetLow(),min_data1.GetClose())

	s:=comm.NewServer("[v0.1]",NewBusinessJob)

	//启动
	s.Serve()
}
