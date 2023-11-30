package core

import (
	"encoding/json"

	toml "github.com/pelletier/go-toml/v2"
	msg "github.com/vmihailenco/msgpack/v5"
	yaml "gopkg.in/yaml.v3"
)

// Represents an addition of a whole document, or a value within an existing document, to the jServ database
type Add struct {
	Collection string                 `json:"collection" msgpack:"collection" csv:"collection" toml:"collection" yaml:"collection"`
	Documents  []Document             `json:"documents,omitempty" msgpack:"documents,omitempty" csv:"documents,omitempty" toml:"documents,omitempty" yaml:"documents,omitempty,flow"`
	Values     map[string]interface{} `json:"values,omitempty" msgpack:"values,omitempty" csv:"values,omitempty" toml:"values,omitempty" yaml:"values,omitempty,flow"`
}

// Converts the Query into JSON
func (a Add) ToJson() string {
	js, _ := json.Marshal(a)
	return string(js)
}

// Converts the Query into MsgPack
func (a Add) ToMsgPack() string {
	m, _ := msg.Marshal(a)
	return string(m)
}

// Converts the Query into TOML
func (a Add) ToToml() string {
	t, _ := toml.Marshal(a)
	return string(t)
}

// Converts the Query into YAML
func (a Add) ToYaml() string {
	y, _ := yaml.Marshal(a)
	return string(y)
}

// JSON Constructor
// Creates the Query object from JSON
func (a *Add) FromJson(s string) {
	if err := json.Unmarshal([]byte(s), a); err != nil {
		panic(err)
	}
}

// MsgPack Constructor
// Creates the Query object from MsgPack
func (a *Add) FromMsgPack(s string) {
	if err := msg.Unmarshal([]byte(s), a); err != nil {
		panic(err)
	}
}

// TOML Constructor
// Creates the Query object from JSON
func (a *Add) FromToml(s string) {
	if err := toml.Unmarshal([]byte(s), a); err != nil {
		panic(err)
	}
}

// YAML Constructor
// Creates the Query object from YAML
func (a *Add) FromYaml(s string) {
	if err := yaml.Unmarshal([]byte(s), a); err != nil {
		panic(err)
	}
}
