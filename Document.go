package core

import (
	"encoding/json"
	"fmt"

	toml "github.com/pelletier/go-toml/v2"
	msg "github.com/vmihailenco/msgpack/v5"
	yaml "gopkg.in/yaml.v3"
)

// Represents a grouping of data under a unique identifier (_id) within the jServ database
type Document struct {
	Id   string                 `json:"_id" msgpack:"_id" csv:"_id" toml:"_id" yaml:"_id"`
	Data map[string]interface{} `json:"data" msgpack:"data" csv:"data" toml:"data" yaml:"data,flow"`
}

// Default Constructor
// Creates an empty Document with only an id
func (d *Document) WithoutData(id string) {
	d.Id = id
	d.Data = make(map[string]interface{})
}

// Map Constructor
// Creates an Document with given id and data map
func (d *Document) WithData(id string, data map[string]interface{}) {
	d.Id = id
	d.Data = data
}

// Values Constructor
// Creates a Document with a given id and a list of Value objects
func (d *Document) WithValues(id string, values []Value) {
	d.Id = id
	d.Data = make(map[string]interface{})
	for _, v := range values {
		d.Data[v.Key] = v.Val
	}
}

// JSON Constructor
// Creates a Document from given Json string
func (d *Document) FromJson(s string) {
	if err := json.Unmarshal([]byte(s), d); err != nil {
		panic(err)
	}
}
func (d *Document) FromMsgPack(s string) {
	if err := msg.Unmarshal([]byte(s), d); err != nil {
		panic(err)
	}
}
func (d *Document) FromToml(s string) {
	if err := toml.Unmarshal([]byte(s), d); err != nil {
		panic(err)
	}
}
func (d *Document) FromYaml(s string) {
	if err := yaml.Unmarshal([]byte(s), d); err != nil {
		panic(err)
	}
}

// Returns the Document as a map
func (d Document) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["id"] = d.Id
	m["data"] = d.Data
	return m
}

// Returns the Document in JSON format
func (d Document) ToJson() string {
	js, _ := json.Marshal(d)
	return string(js)
}
func (d Document) ToMsgPack() string {
	m, _ := msg.Marshal(d)
	return string(m)
}
func (d Document) ToToml() string {
	t, _ := json.Marshal(d)
	return string(t)
}
func (d Document) ToYaml() string {
	y, _ := json.Marshal(d)
	return string(y)
}

// Returns the Document as a string
func (d Document) String() string {
	return fmt.Sprintf(" \"id\" : %s\n \"data\" : %v", d.Id, d.Data)
}
