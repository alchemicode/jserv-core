package core

import (
	"encoding/json"

	toml "github.com/pelletier/go-toml/v2"
	msg "github.com/vmihailenco/msgpack/v5"
	yaml "gopkg.in/yaml.v3"
)

// Represents a server response after receiving and processing HTTP requests
type Response struct {
	// Whether the http request executed successfully ("ok" or "error")
	Status string `json:"status" msgpack:"status" csv:"status" toml:"status" yaml:"status"`
	// Message explaining the status
	Message string `json:"message" msgpack:"message" csv:"message" toml:"message" yaml:"message"`
	// Data containing anything pertinent to the status
	Data map[string]interface{} `json:"data,omitempty" msgpack:"data,omitempty" csv:"data,omitempty" toml:"data,omitempty" yaml:"data,omitempty,flow"`
}

// Constructs a response without the inclusion of data
func (r *Response) WithoutData(status string, message string) {
	r.Status = status
	r.Message = message
	r.Data = map[string]interface{}{}
}

// Constructs a response with the inclusion of data
func (r *Response) WithData(status string, message string, data map[string]interface{}) {
	r.Status = status
	r.Message = message
	r.Data = data
}

// Converts the Response object into JSON
func (r Response) ToJson() string {
	js, _ := json.Marshal(r)
	return string(js)
}

// Converts the Mod into MsgPack
func (r Response) ToMsgPack() string {
	mp, _ := msg.Marshal(r)
	return string(mp)
}

// Converts the Mod into TOML
func (r Response) ToToml() string {
	t, _ := toml.Marshal(r)
	return string(t)
}

// Converts the Mod into YAML
func (r Response) ToYaml() string {
	y, _ := yaml.Marshal(r)
	return string(y)
}

// JSON Constructor
// Creates the Mod object from JSON
func (r *Response) FromJson(s string) {
	if err := json.Unmarshal([]byte(s), r); err != nil {
		panic(err)
	}
}

// MsgPack Constructor
// Creates the Mod object from MsgPack
func (r *Response) FromMsgPack(s string) {
	if err := msg.Unmarshal([]byte(s), r); err != nil {
		panic(err)
	}
}

// TOML Constructor
// Creates the Mod object from TOML
func (r *Response) FromToml(s string) {
	if err := toml.Unmarshal([]byte(s), r); err != nil {
		panic(err)
	}
}

// YAML Constructor
// Creates the Mod object from YAML
func (r *Response) FromYaml(s string) {
	if err := yaml.Unmarshal([]byte(s), r); err != nil {
		panic(err)
	}
}
