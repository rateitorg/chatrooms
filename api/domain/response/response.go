package response

type Response struct {
	Code int `json:"code"`
	Data interface{} `json:"data,omitempty"`
}