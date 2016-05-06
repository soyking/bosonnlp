package bosonnlp

type SuggestAnalysisRawResponse [][]interface{}

type Suggest struct {
	Score float64
	Word  string
}

type SuggestAnalysisResponse []Suggest

// SuggestAnalysis can only analyse one word
// doc: http://docs.bosonnlp.com/suggest.html
func (c *BosonNLPClient) SuggestAnalysis(content string, query ...map[string]string) (SuggestAnalysisResponse, error) {
	var rawResponse SuggestAnalysisRawResponse
	var q map[string]string
	if query != nil {
		q = query[0]
	}
	err := c.post("suggest/analysis", q, content, &rawResponse)
	if err != nil {
		return nil, err
	}

	response := make(SuggestAnalysisResponse, len(rawResponse))
	for i := range rawResponse {
		if len(rawResponse[i]) != 2 {
			return nil, ErrResponseType
		}

		if score, ok := rawResponse[i][0].(float64); ok {
			response[i].Score = score
		} else {
			return nil, err
		}

		if word, ok := rawResponse[i][1].(string); ok {
			response[i].Word = word
		} else {
			return nil, err
		}
	}

	return response, nil
}
