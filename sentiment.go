package bosonnlp

type SentimentAnalysisResponse [][]float32

// 通用
// doc: http://docs.bosonnlp.com/sentiment.html
func (c *BosonNLPClient) SentimentAnalysis(content []string, industry ...string) (SentimentAnalysisResponse, error) {
	var response SentimentAnalysisResponse
	endpoint := "sentiment/analysis"
	if industry != nil {
		endpoint = endpoint + "?" + industry[0]
	}
	err := c.post(endpoint, nil, content, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

// 汽车
// doc: http://docs.bosonnlp.com/sentiment.html
func (c *BosonNLPClient) AutoSentimentAnalysis(content []string) (SentimentAnalysisResponse, error) {
	return c.SentimentAnalysis(content, "auto")
}

// 厨具
// doc: http://docs.bosonnlp.com/sentiment.html
func (c *BosonNLPClient) KitchenSentimentAnalysis(content []string) (SentimentAnalysisResponse, error) {
	return c.SentimentAnalysis(content, "kitchen")
}

// 餐饮
// doc: http://docs.bosonnlp.com/sentiment.html
func (c *BosonNLPClient) FoodSentimentAnalysis(content []string) (SentimentAnalysisResponse, error) {
	return c.SentimentAnalysis(content, "food")
}

// 新闻
// doc: http://docs.bosonnlp.com/sentiment.html
func (c *BosonNLPClient) NewsSentimentAnalysis(content []string) (SentimentAnalysisResponse, error) {
	return c.SentimentAnalysis(content, "news")
}

// 微博
// doc: http://docs.bosonnlp.com/sentiment.html
func (c *BosonNLPClient) WeiboAutoSentimentAnalysis(content []string) (SentimentAnalysisResponse, error) {
	return c.SentimentAnalysis(content, "weibo")
}
