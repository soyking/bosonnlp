package bosonnlp

// _id 该典型意见的标示
// opinion 典型意见文本
// num 该典型意见类似的意见个数
// list 所有属于该典型意见的评论，其中str为意见，int为意见的来源评论ID
type CommentsResponse []struct {
	ID     int        `json:"_id"`
	Option string     `json:"option"`
	Num    int        `json:"num"`
	List   [][]string `json:"list"`
}

// query:
// alpha (0, 1] 默认 0.8 调节聚类最大cluster大小
// beta (0, 1) 默认 0.45 调节聚类平均cluster大小
// doc: http://docs.bosonnlp.com/comments.html
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
	return c.taskPush(task, "comments")
}

func (c *BosonNLPClient) CommentsAnalysis(taskID string, query ...map[string]string) error {
	return c.taskAnalysis(taskID, "comments", query...)
}

func (c *BosonNLPClient) CommentsStatus(taskID string, intervals ...int) error {
	return c.taskStatus(taskID, "comments", intervals...)
}

func (c *BosonNLPClient) CommentsResult(taskID string) (CommentsResponse, error) {
	var response CommentsResponse
	return response, c.taskResult(taskID, "comments", &response)
}
