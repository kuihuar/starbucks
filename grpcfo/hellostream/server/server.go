package main

import (
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	pb "starbucks/grpcfo/hellostream"
	"strconv"
	"sync"
	"time"
)

type server struct {

}


func (me *server)GetStream(req *pb.ReqData, rep pb.HelloStream_GetStreamServer) error{
	log.Println("收到从客户端请求来的数据：", req.Data)
	if req.Data == "hello" {
		cnt := 0
		for  {
			err:= rep.Send(&pb.RepData{
				Data: time.Now().Format(time.RFC3339),
			})
			if err != nil {
				log.Println("send to client failed:", err.Error())
				break
			}
			time.Sleep(time.Second)
			if cnt >= 10 {
				break
			}
			cnt++
		}
	}else {
		log.Println("err request")
	}
	return  nil
}
func (me *server)SetStream(req pb.HelloStream_SetStreamServer) error{
	for{
		data, err := req.Recv()
		if err == nil {
			log.Println("收到从客户端的数据:", data.Data)
			continue
		}else{
			if err == io.EOF {
				log.Println("客户端发送完毕")
			}else{
				log.Println("其它错误")
			}
			break
		}
	}
	req.SendMsg(&pb.RepData{
		Data: "byebye",
	})
	return  nil
}
func (me *server)AllStream(req pb.HelloStream_AllStreamServer) error{
	w := sync.WaitGroup{}
	w.Add(2)
	go func() {
		for {
			recvMsg, err := req.Recv()
			if err != nil {
				if err == io.EOF {
					log.Println("客户端发送完毕", err.Error())
				}else{
					log.Println("其它错误：",err.Error())
				}
				break
			}
			log.Println("收到客户端的数据：", recvMsg.Data)
		}
		w.Done()
	}()
	go func() {
		cnt := 0
		for {
			if err := req.Send(&pb.RepData{
				Data: "从服务端返回 " + strconv.Itoa(cnt),
			});err != nil {
				log.Println("发送失败", err.Error())
				break
			}
			if cnt >=3 {
				break
			}
			cnt++
			time.Sleep(time.Second)
		}
		w.Done()
	}()
	w.Wait()
	return  nil
}
func main()  {
	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		panic(err.Error())
	}
	s := grpc.NewServer()
	pb.RegisterHelloStreamServer(s, &server{})
	s.Serve(lis)
}