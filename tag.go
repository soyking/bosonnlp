package bosonnlp

// TagType from TagAnalysis
var TagType = map[string]string{}

type TagAnalysisResponse []struct {
	Tag  []string `json:"tag" bson:"tag"`
	Word []string `json:"word" bson:"word"`
}

// doc: http://docs.bosonnlp.com/tag.html
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
