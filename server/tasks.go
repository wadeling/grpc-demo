package main

import (
	"context"
	"fmt"
	"github.com/wadeling/grpc-demo/pb"
	"log"
	"math/rand"
	"sync"
	"time"
)

type TaskManager struct {
	lock  sync.Mutex
	tasks map[string]chan *pb.Task // client-ip -> task
}

func (t *TaskManager) AddClient(clientIp string) (chan *pb.Task, error) {
	t.lock.Lock()
	defer t.lock.Unlock()

	value, ok := t.tasks[clientIp]
	if ok {
		log.Printf("[warn] client allready exist.%s", clientIp)
		return value, nil
	}
	log.Printf("create new task chan for client:%s", clientIp)

	ct := make(chan *pb.Task, 1)
	t.tasks[clientIp] = ct
	return t.tasks[clientIp], nil
}

func (t *TaskManager) PutTask(clientIp string, pt *pb.Task) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	_, ok := t.tasks[clientIp]
	if !ok {
		log.Printf("[error] not found client %s", clientIp)
		return fmt.Errorf("not found client %s", clientIp)
	}
	t.tasks[clientIp] <- pt
	return nil
}

func (t *TaskManager) RandGenerateTask() error {
	log.Print("task manager generate task")

	t.lock.Lock()
	defer t.lock.Unlock()

	if len(t.tasks) == 0 {
		log.Print("not any client connected,abort generate task")
		return nil
	}

	// get all cluster ip
	arr := make([]string, 0)
	for k := range t.tasks {
		arr = append(arr, k)
	}
	index := rand.Intn(len(arr))
	clientIp := arr[index]
	log.Printf("ready to generate task for client %v", clientIp)

	// mock task
	task := pb.Task{
		TaskID: fmt.Sprintf("task-%s-%d", clientIp, time.Now().Unix()),
		Repo:   "wade23",
		Tag:    "latest",
	}
	log.Printf("prepare task ok,ready to write:%+v", &task)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	select {
	case t.tasks[clientIp] <- &task:
		log.Printf("write task ok.%s", task.TaskID)
		return nil
	case <-ctx.Done():
		log.Printf("[error] write task failed.timeout")
		return fmt.Errorf("write task failed.timeout")
	}
}

func (t *TaskManager) run() error {
	for {
		time.Sleep(10 * time.Second)

		// generate task
		_ = t.RandGenerateTask()

	}
}

func (t *TaskManager) Run() error {
	rand.Seed(time.Now().UnixNano())
	log.Printf("task manager start running")

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		_ = t.run()
	}()
	wg.Wait()
	log.Printf("task manager finish.")
	return nil
}

func NewTaskManager() *TaskManager {
	t := TaskManager{
		tasks: make(map[string]chan *pb.Task),
	}
	return &t
}
