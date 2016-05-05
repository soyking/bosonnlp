package bosonnlp

type TagAnalysisResponse []struct {
	Tag  []string `json:"tag" bson:"tag"`
	Word []string `json:"word" bson:"word"`
}

func (c *BosonNLPClient) TagAnalysis(content []string, query ...map[string]string) (TagAnalysisResponse, error) {
	var response TagAnalysisResponse
	var q map[string]string
	if query != nil {
		q = query[0]
	}
	err := c.post("tag/analysis", q, content, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}
