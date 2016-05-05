package bosonnlp

import "errors"

var (
	ErrWrongToken         = errors.New("token is wrong")
	ErrTooManyArticles    = errors.New("number of articles should less than 100")
	ErrExceedRequestLimit = errors.New("exceed request frequency limit")
	ErrResponseType       = errors.New("response from boson with wrong type")
)
