package bosonnlp

func (c *BosonNLPClient) Cluster(tasks []Task, taskIDs ...string) error {
	var taskID string
	if len(taskID) != 0 {
		taskID = taskIDs[0]
	} else {
		taskID = GenerateID()
	}

	var pushResponse TaskPushResponse
	err := c.post("cluster/push/"+taskID, nil, tasks, &pushResponse)
	if err != nil {
		return err
	}

	if pushResponse.Count != len(tasks) || pushResponse.TaskID != taskID {
		return ErrResponseType
	}

	var status string
	//for status != "DONE" {
	c.get("cluster/status/"+taskID, nil, &status)
	println(status)
	//}

	return nil
}
