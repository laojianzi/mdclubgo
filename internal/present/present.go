package present

// Data for api response
type Data struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// Build return a response data
func Build(data interface{}) *Data {
	return &Data{
		Code: Normal,
		Data: data,
	}
}
