package bosonnlp

// _id 该cluster最具代表性的文档
// num 该cluster包含的文档数目
// list 所有属于该cluster的文档 _id
type ClusterResponse []struct {
	ID   string   `json:"_id" bson:"_id"`
	Num  int      `json:"num" bson:"num"`
	List []string `json:"list" bson:"list"`
}

// query:
// alpha (0, 1] 默认 0.8 调节聚类最大cluster大小
// beta (0, 1) 默认 0.45 调节聚类平均cluster大小
// doc: http://docs.bosonnlp.com/cluster.html
func (c *BosonNLPClient) Cluster(task *CommonTask, query ...map[string]string) (ClusterResponse, error) {
	taskID, err := c.ClusterPush(task)
	if err != nil {
		return nil, err
	}

	err = c.ClusterAnalysis(taskID, query...)
	if err != nil {
		return nil, err
	}

	err = c.ClusterStatus(taskID)
	if err != nil {
		return nil, err
	}

	return c.ClusterResult(taskID)
}

func (c *BosonNLPClient) ClusterPush(task *CommonTask) (string, error) {
	return c.taskPush(task, "cluster")
}

func (c *BosonNLPClient) ClusterAnalysis(taskID string, query ...map[string]string) error {
	return c.taskAnalysis(taskID, "cluster", query...)
}

func (c *BosonNLPClient) ClusterStatus(taskID string, intervals ...int) error {
	return c.taskStatus(taskID, "cluster", intervals...)
}

func (c *BosonNLPClient) ClusterResult(taskID string) (ClusterResponse, error) {
	var response ClusterResponse
	return response, c.taskResult(taskID, "cluster", &response)
}
