package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	pb "starbucks/grpcfo/hellostream"
	"strconv"
)

func main()  {
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	client := pb.NewHelloStreamClient(conn)


	//stream, err:=client.GetStream(context.Background(),&pb.ReqData{
	//	Data: "hello1",
	//})
	//if err != nil {
	//	panic(err.Error())
	//}
	//for{
	//	data, err:= stream.Recv()
	//	if err != nil {
	//		if err == io.EOF{
	//			log.Println("服务端发送完毕")
	//			break
	//		}
	//		log.Println("其它错误：",err.Error())
	//		break
	//	}
	//	log.Println("收到的服务端数据:", data.Data)
	//}

	//sendBuffer := []string{"hello", "abc.com", "areyou ok!"}
	//stream, err := client.SetStream(context.Background())
	//if err != nil {
	//	panic(err.Error())
	//}
	//if stream == nil {
	//	panic("stream is nil")
	//}
	//for _,v := range sendBuffer {
	//	if err := stream.Send(&pb.ReqData{
	//		Data: v,
	//	}); err != nil {
	//		log.Println("send failed", err.Error())
	//	}
	//	time.Sleep(time.Second)
	//}
	//repData, err := stream.CloseAndRecv()
	//if err != nil {
	//	log.Println("recv from server rep failed:", err.Error())
	//	return
	//}
	//log.Println("收到服务端响应 recv:", repData.Data)
	//客户端发起流，服务端返回流
	stream, err := client.AllStream(context.Background())
	if err != nil || stream == nil {
		panic("init failed")
	}
	//发送流
	go func() {
		cnt := 0
		for {
			if err := stream.Send(&pb.ReqData{
				Data: "client: " + strconv.Itoa(cnt),
			}); err != nil {
				log.Println("send failed:", err.Error())
				return
			}
			if cnt >= 10 {
				if err = stream.CloseSend(); err != nil {
					log.Println("send close err", err.Error())
					return
				}
				break
			}
			cnt ++
		}
	}()

	//接收流

	go func() {
		for {
			data, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					log.Println("server close conn")
				}else{
					log.Println("recr from server failed:", err.Error())
				}
				return
			}
			log.Println("recv from server:", data.Data)
		}
	}()
	select {
	}
}
