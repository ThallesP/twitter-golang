package usecase

type Exception struct {
	error
	ErrorName  string `json:"errname"`
	Message    string `json:"message"`
	StatusCode int
}

func NewException(message string, statuscode int, errname string) *Exception {
	return &Exception{
		ErrorName:  errname,
		Message:    message,
		StatusCode: statuscode,
	}
}
