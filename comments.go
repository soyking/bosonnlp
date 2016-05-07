package bosonnlp

import "time"

type CommentsResponse []struct {
	ID     int      `json:"_id" bson:"_id"`
	Option string      `json:"option" bson:"option"`
	Num    int         `json:"num" bson:"num"`
	List   [][]string `json:"list" bson:"list"`
}

// query:
// alpha (0, 1] 默认 0.8 调节聚类最大cluster大小
// beta (0, 1) 默认 0.45 调节聚类平均cluster大小
func (c *BosonNLPClient) Comments(task *CommonTask, query ...map[string]string) (CommentsResponse, error) {
	taskID, err := c.CommentsPush(task)
	if err != nil {
		return nil, err
	}

	err = c.CommentsAnalysis(taskID, query...)
	if err != nil {
		return nil, err
	}

	err = c.CommentsStatus(taskID)
	if err != nil {
		return nil, err
	}

	return c.CommentsResult(taskID)
}

func (c *BosonNLPClient) CommentsPush(task *CommonTask) (string, error) {
	task.Check()
	var pushResponse TaskPushResponse
	err := c.post("comments/push/" + task.TaskID, nil, task.Tasks, &pushResponse)
	if err != nil {
		return "", err
	}

	if pushResponse.Count != len(task.Tasks) || pushResponse.TaskID != task.TaskID {
		return "", ErrResponseType
	}

	return task.TaskID, nil
}

func (c *BosonNLPClient) CommentsAnalysis(taskID string, query ...map[string]string) error {
	var q map[string]string
	if query != nil {
		q = query[0]
	}

	var s TaskStatus
	c.get("comments/analysis/" + taskID, q, &s)
	if s.Status != STATUS_RECEIVED {
		return ErrTaskNotReceive
	}
	return nil
}

const INTERVAL int = 1000

func (c *BosonNLPClient) CommentsStatus(taskID string, intervals ...int) error {
	interval := INTERVAL
	if intervals != nil {
		interval = intervals[0]
	}

	var s TaskStatus
	for {
		err := c.get("comments/status/" + taskID, nil, &s)
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

func (c *BosonNLPClient) CommentsResult(taskID string) (CommentsResponse, error) {
	var response CommentsResponse
	err := c.get("comments/result/" + taskID, nil, &response)
	return response, err
}
