package bosonnlp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const host = "http://api.bosonnlp.com/"

type BosonNLPClient struct {
	Token      string
	httpClient *http.Client
}

func NewBosonNLPClient(token string) *BosonNLPClient {
	return &BosonNLPClient{Token: token, httpClient: http.DefaultClient}
}

func (c *BosonNLPClient) post(endpoint string, query map[string]string, body interface{}, response interface{}) error {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil
	}

	url := host + endpoint
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}
	request.Header.Set("X-token", c.Token)
	if query != nil {
		values := request.URL.Query()
		for k, v := range query {
			values.Add(k, v)
		}
		request.URL.RawQuery = values.Encode()
	}
	println(request.URL.String())

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	println(string(respBytes))

	// err check
	var errResponse struct {
		Status int `json:"status" bson:"status"`
	}
	json.Unmarshal(respBytes, &errResponse)
	if errResponse.Status == 403 {
		return ErrWrongToken
	} else if errResponse.Status == 413 {
		return ErrTooManyArticles
	} else if errResponse.Status == 429 {
		return ErrExceedRequestLimit
	}

	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return err
	}

	return nil
}
