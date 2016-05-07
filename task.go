package bosonnlp

import (
	"crypto/rand"
	"fmt"
)

type TaskPushResponse struct {
	Count  int    `json:"count" bson:"count"`
	TaskID string `json:"task_id" bson:"task_id"`
}

const (
	STATUS_NOT_FOUND = "NOT FOUND"
	STATUS_RECEIVED  = "RECEIVED"
	STATUS_RUNNING   = "RUNNING"
	STATUS_ERROR     = "ERROR"
	STATUS_DONE      = "DONE"
)

type TaskStatus struct {
	TaskID string `json:"_id" bson:"_id"`
	Count  int    `json:"count" bson:"count"`
	Status string `json:"status" bson:"status"`
}

type Task struct {
	ID   string `json:"_id" bson:"_id"`
	Text string `json:"text" bson:"text"`
}

func (t *Task) Check() {
	if t.ID == "" {
		t.ID = GenerateID()
	}
}

func GenerateID() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

type CommonTask struct {
	TaskID string
	Tasks  []Task
}

func (t *CommonTask) AddTask(task Task) {
	t.Tasks = append(t.Tasks, task)
}

func (t *CommonTask) AddText(text string) {
	t.Tasks = append(t.Tasks, Task{Text: text})
}

func (t *CommonTask) Check() {
	if t.TaskID == "" {
		t.TaskID = GenerateID()
	}
	for i := range t.Tasks {
		t.Tasks[i].Check()
	}
}

func (c *BosonNLPClient) NewCommentsTask(taskIDs ...string) *CommonTask {
	var taskID string
	if len(taskID) != 0 {
		taskID = taskIDs[0]
	} else {
		taskID = GenerateID()
	}
	return &CommonTask{TaskID: taskID}
}
