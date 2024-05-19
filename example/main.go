package main

import (
	"fmt"
	"time"

	workerpool "github.com/swift-pace/worker-pool"
)

type EmailSender struct {
	email string
}

func (s *EmailSender) RunTask() {
	// Simulate execute time
	// time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	time.Sleep(2 * time.Second)

	fmt.Printf("Email %s sent\n", s.email)
}

func main() {
	tasks := make([]workerpool.TaskRunner, 0, 20)
	for index := 0; index < 20; index++ {
		tasks = append(tasks, &EmailSender{
			email: fmt.Sprintf("email%d@example.com", index+1),
		})
	}

	wp := workerpool.NewWorkerPool(tasks, 5)
	wp.Start()
}
