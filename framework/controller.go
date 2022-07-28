package framework

import (
	"log"
	"net/http"
	"xgo/framework/task"
)

type ControllerHandler func(c *XContext) error

func IndexController(c *XContext) error {
	//异步任务test
	task, err := task.NewEmailDeliveryTask(42, "some:template:id")
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}
	info, err := Queue.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	//Json输出
	c.Json(http.StatusOK, map[string]interface{}{
		"name":    "hello",
		"address": "nihao",
	})

	return nil
}
