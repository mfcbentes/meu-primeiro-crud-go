package rest_err

import "net/http"

type RestErr interface {
	Message() string
	Err() string
	Code() int
	Causes() []Causes
	Error() string
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type restErr struct {
	ErrMessage string   `json:"message"`
	ErrError   string   `json:"error"`
	ErrCode    int      `json:"code"`
	ErrCauses  []Causes `json:"causes,omitempty"`
}

func (r *restErr) Message() string {
	return r.ErrMessage
}

func (r *restErr) Err() string {
	return r.ErrError
}

func (r *restErr) Code() int {
	return r.ErrCode
}

func (r *restErr) Causes() []Causes {
	return r.ErrCauses
}

func (r *restErr) Error() string {
	return r.ErrMessage
}

func NewRestErr(message, err string, code int, causes []Causes) RestErr {
	return &restErr{
		ErrMessage: message,
		ErrError:   err,
		ErrCode:    code,
		ErrCauses:  causes,
	}
}

func NewBadRequestError(message string) RestErr {
	return &restErr{
		ErrMessage: message,
		ErrError:   "bad_request",
		ErrCode:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) RestErr {
	return &restErr{
		ErrMessage: message,
		ErrError:   "bad_request",
		ErrCode:    http.StatusBadRequest,
		ErrCauses:  causes,
	}
}

func NewInternalServerError(message string) RestErr {
	return &restErr{
		ErrMessage: message,
		ErrError:   "internal_server_error",
		ErrCode:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) RestErr {
	return &restErr{
		ErrMessage: message,
		ErrError:   "not_found",
		ErrCode:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) RestErr {
	return &restErr{
		ErrMessage: message,
		ErrError:   "forbidden",
		ErrCode:    http.StatusForbidden,
	}
}
