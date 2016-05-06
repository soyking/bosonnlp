package bosonnlp

// TagType from TagAnalysis
var TagType = map[string]string{
	"n":     "名词",
	"nr":    "人名",
	"nr1":   "中文姓氏",
	"nrf":   "音译人名",
	"ns":    "地名",
	"nt":    "组织机构名",
	"nz":    "其它专有名词",
	"nl":    "名词性惯用语",
	"t":     "时间词",
	"s":     "处所词",
	"f":     "方位词",
	"v":     "动词",
	"vd":    "副动词",
	"vshi":  "动词",
	"vyou":  "动词",
	"vi":    "不及物动词",
	"vl":    "动词性惯用语",
	"a":     "形容词",
	"ad":    "副形词",
	"an":    "名形词",
	"al":    "形容词性惯用语",
	"b":     "区别词",
	"bl":    "区别词性惯用语",
	"z":     "状态词",
	"r":     "代词",
	"m":     "数词",
	"q":     "量词",
	"d":     "副词",
	"dl":    "副词性惯用语",
	"p":     "介词",
	"pba":   "介词“把”",
	"pbei":  "介词“被”",
	"c":     "连词",
	"u":     "助词",
	"uzhe":  "助词“着”",
	"ule":   "助词“了”",
	"uguo":  "助词“过”",
	"ude":   "助词“的”、“地”、“得”",
	"usuo":  "助词“所”",
	"udeng": "助词“等”、“等等”",
	"uyy":   "助词“一样”、“似的”",
	"udh":   "助词“的话”",
	"uzhi":  "助词“之”",
	"ulian": "助词“连”",
	"y":     "语气词",
	"o":     "拟声词",
	"h":     "前缀",
	"k":     "后缀",
	"nx":    "字符串",
	"w":     "标点符号",
	"wkz":   "左括号",
	"wky":   "右括号",
	"wyz":   "左引号",
	"wyy":   "右引号",
	"wj":    "句号",
	"ww":    "问号",
	"wt":    "叹号",
	"wd":    "逗号",
	"wf":    "分号",
	"wn":    "顿号",
	"wm":    "冒号",
	"ws":    "省略号",
	"wp":    "破折号",
	"wb":    "百分号千分号",
	"wh":    "单位符号",
	"email": "电子邮件地址”",
	"tel":   "电话号码",
	"id":    "身份证号",
	"ip":    "ip地址",
	"url":   "网页链接”",
}

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
