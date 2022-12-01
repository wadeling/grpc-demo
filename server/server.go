package main

import (
	"flag"
	"fmt"
	"github.com/wadeling/grpc-demo/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 5051, "The server port")
)

type TaskServer struct {
	pb.UnimplementedTaskProxyServer
	taskManager *TaskManager
}

func (t *TaskServer) GetTask(req *pb.Req, stream pb.TaskProxy_GetTaskServer) error {
	log.Printf("recv client req:%v,%v", req.ClusterKey, req.NodeIp)
	taskChan, _ := t.taskManager.AddClient(req.NodeIp)

	clientIp := req.NodeIp

	for {
		//time.Sleep(5 * time.Second)

		log.Printf("waiting task for client %v", clientIp)
		task := <-taskChan
		log.Printf("get task for client %v", clientIp)

		err := stream.Send(task)
		if err != nil {
			log.Printf("send task to client %v failed.%v", clientIp, err)
			return err
		} else {
			log.Printf("send task to client %v ok", clientIp)
		}
	}
}

func NewTaskServer(manager *TaskManager) *TaskServer {
	t := &TaskServer{
		taskManager: manager,
	}
	return t
}

func main() {
	log.Print("start server")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// new task manager
	taskManager := NewTaskManager()

	go func() {
		_ = taskManager.Run()
	}()

	// create grpc server
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterTaskProxyServer(grpcServer, NewTaskServer(taskManager))
	err = grpcServer.Serve(lis)
	log.Printf("grpc server exit:%v", err)
}
