package bosonnlp

import "errors"

var (
	ErrWrongToken         = errors.New("err code:403 token is wrong")
	ErrTooManyArticles    = errors.New("err code:413 number of articles should less than 100")
	ErrExceedRequestLimit = errors.New("err code:429 exceed request frequency limit")
	ErrTaskNotFound       = errors.New("err code:404 task not found")
	ErrResponseType       = errors.New("response from boson with wrong type")
	ErrTaskError          = errors.New("task error")
	ErrTaskNotReceive     = errors.New("task not receive")
)
