package bosonnlp

type KeywordsAnalysisRawResponse [][][]interface{}

type KeyWord struct {
	Weight float64
	Word   string
}

type KeywordsAnalysisResponse [][]KeyWord

// doc: http://docs.bosonnlp.com/keywords.html
func (c *BosonNLPClient) KeywordsAnalysis(content []string, query ...map[string]string) (KeywordsAnalysisResponse, error) {
	var rawResponse KeywordsAnalysisRawResponse
	var q map[string]string
	if query != nil {
		q = query[0]
	}
	err := c.post("keywords/analysis", q, content, &rawResponse)
	if err != nil {
		return nil, err
	}

	response := make(KeywordsAnalysisResponse, len(rawResponse))
	for i := range rawResponse {
		response[i] = make([]KeyWord, len(rawResponse[i]))
		for j := range response[i] {
			if len(rawResponse[i][j]) != 2 {
				return nil, ErrResponseType
			}
			if weight, ok := rawResponse[i][j][0].(float64); ok {
				response[i][j].Weight = weight
			} else {
				return nil, err
			}

			if word, ok := rawResponse[i][j][1].(string); ok {
				response[i][j].Word = word
			} else {
				return nil, err
			}
		}

	}

	return response, err
}
