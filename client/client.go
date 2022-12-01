package main

import (
	"context"
	"flag"
	"github.com/wadeling/grpc-demo/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
)

var (
	serverAddr = flag.String("addr", "localhost:5051", "The server address in the format of host:port")
	ip         = flag.String("ip", "10.0.0.1", "The client ip address")
)

type TaskClient struct {
	client pb.TaskProxyClient
}

func (t *TaskClient) ReqTask() {
	//ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	//defer cancel()

	req := t.MockReq()
	stream, err := t.client.GetTask(context.Background(), req)
	if err != nil {
		log.Fatalf("client.GetTask failed: %v", err)
		return
	}

	for {
		task, err := stream.Recv()
		if err == io.EOF {
			log.Print("client recv task eof")
			break
		}
		if err != nil {
			log.Fatalf("client.Reqtask failed: %v", err)
		}
		log.Printf("task: id: %v,repo %v,tag %v ", task.GetTaskID(), task.GetRepo(), task.GetTag())
	}
	log.Printf("req task end")
}

func (t *TaskClient) MockReq() *pb.Req {
	r := pb.Req{
		ClusterKey: "aaa",
		NodeIp:     *ip,
	}
	return &r
}

func NewTaskClient(client pb.TaskProxyClient) *TaskClient {
	t := TaskClient{
		client: client,
	}
	return &t
}

func main() {
	log.Print("start client")
	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer func() { _ = conn.Close() }()
	client := pb.NewTaskProxyClient(conn)

	tc := NewTaskClient(client)
	tc.ReqTask()

	log.Printf("client exit")
}
