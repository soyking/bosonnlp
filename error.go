package bosonnlp

import "errors"

var (
	ErrWrongToken         = errors.New("err code:403 token is wrong")
	ErrTooManyArticles    = errors.New("err code:413 number of articles should less than 100")
	ErrExceedRequestLimit = errors.New("err code:429 exceed request frequency limit")
	ErrResponseType       = errors.New("response from boson with wrong type")
)
