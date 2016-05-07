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
	return c.handleRequest(request, query, response)
}

func (c *BosonNLPClient) get(endpoint string, query map[string]string, response interface{}) error {
	url := host + endpoint
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	return c.handleRequest(request, query, response)
}

func (c *BosonNLPClient) handleRequest(request *http.Request, query map[string]string, response interface{}) error {
	if query != nil {
		values := request.URL.Query()
		for k, v := range query {
			values.Add(k, v)
		}
		request.URL.RawQuery = values.Encode()
	}
	request.Header.Set("X-token", c.Token)
	println(request.URL.String())

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	err = c.errCheck(respBytes)
	if err != nil {
		return err
	}
	println(string(respBytes))

	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return err
	}
	return nil
}

type BosonNLPErrResponse struct {
	Status int `json:"status"`
}

func (c *BosonNLPClient) errCheck(respBytes []byte) error {
	var errResponse BosonNLPErrResponse
	err := json.Unmarshal(respBytes, &errResponse)
	if err == nil {
		if errResponse.Status == 403 {
			return ErrWrongToken
		} else if errResponse.Status == 413 {
			return ErrTooManyArticles
		} else if errResponse.Status == 429 {
			return ErrExceedRequestLimit
		} else if errResponse.Status == 404 {
			return ErrTaskNotFound
		}
	}
	return nil
}
