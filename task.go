package bosonnlp

import (
	"crypto/rand"
	"fmt"
	"time"
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

	// task polling interval
	INTERVAL int = 1000
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

func (c *BosonNLPClient) NewCommonTask(taskIDs ...string) *CommonTask {
	var taskID string
	if len(taskID) != 0 {
		taskID = taskIDs[0]
	} else {
		taskID = GenerateID()
	}
	return &CommonTask{TaskID: taskID}
}

func (c *BosonNLPClient) taskPush(task *CommonTask, typeName string) (string, error) {
	task.Check()
	var pushResponse TaskPushResponse
	err := c.post(typeName+"/push/"+task.TaskID, nil, task.Tasks, &pushResponse)
	if err != nil {
		return "", err
	}

	if pushResponse.Count != len(task.Tasks) || pushResponse.TaskID != task.TaskID {
		return "", ErrResponseType
	}

	return task.TaskID, nil
}

func (c *BosonNLPClient) taskAnalysis(taskID string, typeName string, query ...map[string]string) error {
	var q map[string]string
	if query != nil {
		q = query[0]
	}

	var s TaskStatus
	c.get(typeName+"/analysis/"+taskID, q, &s)
	if s.Status != STATUS_RECEIVED {
		return ErrTaskNotReceive
	}
	return nil
}

func (c *BosonNLPClient) taskStatus(taskID string, typeName string, intervals ...int) error {
	interval := INTERVAL
	if intervals != nil {
		interval = intervals[0]
	}

	var s TaskStatus
	for {
		err := c.get(typeName+"/status/"+taskID, nil, &s)
		if err != nil {
			return err
		}

		if s.Status == STATUS_ERROR {
			return ErrTaskError
		} else if s.Status == STATUS_NOT_FOUND {
			return ErrTaskNotFound
		} else if s.Status == STATUS_DONE {
			break
		}

		time.Sleep(time.Duration(interval) * time.Millisecond)
	}

	return nil
}

func (c *BosonNLPClient) taskResult(taskID string, typeName string, response interface{}) error {
	err := c.get(typeName+"/result/"+taskID, nil, &response)
	return err
}
