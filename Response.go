package core

import "encoding/json"

//Represents a server response after receiving and processing HTTP requests
type Response struct {
	// Whether the http request executed successfully ("ok" or "error")
	Status string `json:"status"`
	// Message explaining the status
	Message string `json:"message"`
	// JSON data containing anything pertinent to explaining the status
	Data string `json:"data"`
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
func (r Response) DataToMap() (map[string]interface{}, error) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(r.Data), &data); err != nil {
		return map[string]interface{}{}, err
	}
	return data, nil
}

//Converts the Response object into JSON
func (r Response) ToJson() string {
	data := make(map[string]interface{})
	data["status"] = r.Status
	data["message"] = r.Message
	if r.Data != "" {
		dat, err := r.DataToMap()
		if err != nil {
			data["data"] = dat
		}
	}
	js, _ := json.Marshal(data)
	return string(js)
}
