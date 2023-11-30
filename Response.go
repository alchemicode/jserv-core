package core

import "encoding/json"

//Represents a server response after receiving and processing HTTP requests
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

// Constructs a response without the inclusion of JSON data
func (r *Response) WithoutData(status string, message string) {
	r.Status = status
	r.Message = message
	r.Data = ""
}

// Constructs a response with the inclusion of JSON data
func (r *Response) WithData(status string, message string, data string) {
	r.Status = status
	r.Message = message
	r.Data = data
}

//Converts the JSON string field Data to a map
func (r Response) DataToMap() map[string]interface{} {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(r.Data), &data); err != nil {
		panic(err)
	}
	return data
}

//Converts the Response object into JSON
func (r Response) ToJson() string {
	data := make(map[string]interface{})
	data["status"] = r.Status
	data["message"] = r.Message
	data["data"] = r.DataToMap()
	js, _ := json.Marshal(data)
	return string(js)
}
