package bosonnlp

// EntityType from NerAnalysis
var EntityType = map[string]string{
	"time":         "时间",
	"location":     "地点",
	"person_name":  "人名",
	"org_name":     "组织名",
	"company_name": "公司名",
	"product_name": "产品名",
	"job_title":    "职位",
}

type Entity struct {
	Start      int
	End        int
	EntityType string
}
type NerAnalysisResponse []struct {
	// json unmarshal data
	RawEntity [][]interface{} `json:"entity"`
	// converted data
	Entity []Entity
	Tag    []string `json:"tag"`
	Word   []string `json:"word"`
}

// doc: http://docs.bosonnlp.com/ner.html
func (c *BosonNLPClient) NerAnalysis(content []string) (NerAnalysisResponse, error) {
	var response NerAnalysisResponse
	err := c.post("ner/analysis", nil, content, &response)
	if err != nil {
		return nil, err
	}

	for i := range response {
		response[i].Entity = make([]Entity, len(response[i].RawEntity))
		for j := range response[i].RawEntity {
			rawEntity := response[i].RawEntity[j]
			if len(rawEntity) != 3 {
				return nil, ErrResponseType
			}

			if start, ok := rawEntity[0].(float64); ok {
				response[i].Entity[j].Start = int(start)
			} else {
				return nil, ErrResponseType
			}

			if end, ok := rawEntity[1].(float64); ok {
				response[i].Entity[j].End = int(end)
			} else {
				return nil, ErrResponseType
			}

			if entityType, ok := rawEntity[2].(string); ok {
				response[i].Entity[j].EntityType = entityType
			} else {
				return nil, ErrResponseType
			}
		}
	}

	return response, nil
}
