package framework

import (
	"github.com/hibiken/asynq"
	"log"
	"xgo/framework/task"
)

var Queue *asynq.Client

//消息队列redis实例
const redisAddr = "127.0.0.1:6379"

func init() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	Queue = client
}

func RunQueueServer() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// See the godoc for other configuration options
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.HandleFunc(task.TypeEmailDelivery, task.HandleEmailDeliveryTask)
	mux.Handle(task.TypeImageResize, task.NewImageProcessor())

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}