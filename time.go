package bosonnlp

// TimeType from TimeAnalysis
var TimeType = map[string]string{
	"timestamp":  "时间点",  // 时间点，ISO8601格式的时间字符串
	"timedelta":  "时间量",  // 时间量，格式为”xday,HH:MM:SS”或”HH:MM:SS”的字符串
	"timespan_0": "时间区间", // 表示时间点组成的时间区间结果，格式为[timestamp,timestamp]表示时间区间的起始和结束时间
	"timespan_1": "时间区间", // 时间区间结果，格式为[timedelta, timedelta]，表示时间区间的起始和结束时间
	"":           "N/A",  // 不能识别，返回空字符串
}

type TimeAnalysisResponse struct {
	Timestamp string `json:"timestamp" bson:"timestamp"`
	Type      string `json:"type" bson:"type"`
}

// doc: http://docs.bosonnlp.com/time.html
func (c *BosonNLPClient) TimeAnalysis(pattern string) (*TimeAnalysisResponse, error) {
	var response TimeAnalysisResponse
	err := c.post("time/analysis", map[string]string{"pattern": pattern}, nil, &response)
	return &response, err
}
