package bosonnlp

// RoleType from DepparserAnalysis
var RoleType = map[string]string{
	"ROOT": "核心词",
	"SBJ":  "主语成分",
	"OBJ":  "宾语成分",
	"PU":   "标点符号",
	"TMP":  "时间成分",
	"LOC":  "位置成分",
	"MNR":  "方式成分",
	"POBJ": "介宾成分",
	"PMOD": "介词修饰",
	"NMOD": "名词修饰",
	"VMOD": "动词修饰",
	"VRD":  "动结式（第二动词为第一动词结果）",
	"DEG":  "连接词“的”结构",
	"DEV":  "“地”结构",
	"LC":   "位置词结构",
	"M":    "量词结构",
	"AMOD": "副词修饰",
	"PRN":  "括号成分",
	"VC":   "动词“是”修饰",
	"COOR": "并列关系",
	"CS":   "从属连词成分",
	"DEC":  "关系从句“的”",
}

type DepparserAnalysisResponse []struct {
	Role []string `json:"role" bson:"role"`
	Head []int    `json:"head" bson:"head"`
	Tag  []string `json:"tag" bson:"tag"`
	Word []string `json:"word" bson:"word"`
}

// doc: http://docs.bosonnlp.com/depparser.html
func (c *BosonNLPClient) DepparserAnalysis(content []string) (DepparserAnalysisResponse, error) {
	var response DepparserAnalysisResponse
	err := c.post("depparser/analysis", nil, content, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}
