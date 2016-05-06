package bosonnlp

type SummaryAnalysisRequest struct {
	Titile     string  `json:"title" bson:"title"`
	Content    string  `json:"content" bson:"content"`
	Percentage float64 `json:"percentage" bson:"percentage"`
	NotExceed  int     `json:"not_exceed" bson:"not_exceed"`
}

type SummaryAnalysisResponse string

// title 可以省略
//
// percentage： 字数限制 float or int, p <= 1，则认为是摘要字数与原文总字数的百分比；p > 1， 则认为是最终摘要的具体字数
//
// not_exceed： 0（default）1 是否开启严格字数限制开关，不开启，摘要字数大于等于字数限制以确保句子完整；开启，摘要字数小于等于字数限制
//
// doc： http://docs.bosonnlp.com/summary.html
func (c *BosonNLPClient) SummaryAnalysis(title, content string, percentage float64, notExceed bool) (SummaryAnalysisResponse, error) {
	var response SummaryAnalysisResponse
	request := SummaryAnalysisRequest{
		Titile:     title,
		Content:    content,
		Percentage: percentage,
	}
	if notExceed {
		request.NotExceed = 1
	} else {
		request.NotExceed = 0
	}

	err := c.post("summary/analysis", nil, request, &response)
	if err != nil {
		return "", err
	}
	return response, err
}
