package bosonnlp

// NewsType id from ClassifyAnalysis
var NewsType = map[int]string{
	0:  "体育",
	1:  "教育",
	2:  "财经",
	3:  "社会",
	4:  "娱乐",
	5:  "军事",
	6:  "国内",
	7:  "科技",
	8:  "互联网",
	9:  "房产",
	10: "国际",
	11: "女人",
	12: "汽车",
	13: "游戏",
}

type ClassifyAnalysisResponse []int

// doc: http://docs.bosonnlp.com/classify.html
func (c *BosonNLPClient) ClassifyAnalysis(content []string) (ClassifyAnalysisResponse, error) {
	var response ClassifyAnalysisResponse
	err := c.post("classify/analysis", nil, content, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}
