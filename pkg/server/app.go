package server

type App interface {
	Run() error
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *Response) GetStatus() int {
	return r.Status
}
