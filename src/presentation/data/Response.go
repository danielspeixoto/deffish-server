package data

type Response struct {
	Status string `json:"status"`
	Data interface{} `json:"data"`
}
