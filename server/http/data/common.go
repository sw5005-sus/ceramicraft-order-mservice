package data

const (
	ServiceName = "order-ms"
)

type BaseResponse struct {
	Code   int         `json:"code"`
	ErrMsg string      `json:"err_msg,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}
