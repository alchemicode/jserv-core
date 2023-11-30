package core

import (
	"encoding/json"

	toml "github.com/pelletier/go-toml/v2"
	msg "github.com/vmihailenco/msgpack/v5"
	yaml "gopkg.in/yaml.v3"
)

// Represents a modification to a specific document within the jServ database
type Mod struct {
	Collection string                 `json:"collection" msgpack:"collection" csv:"collection" toml:"collection" yaml:"collection"`
	Document   string                 `json:"document" msgpack:"document" csv:"document" toml:"document" yaml:"document"`
	Values     map[string]interface{} `json:"values" msgpack:"values" csv:"values" toml:"values" yaml:"values,flow"`
}

// Converts the Query into JSON
func (m Mod) ToJson() string {
	js, _ := json.Marshal(m)
	return string(js)
}

// Converts the Query into MsgPack
func (m Mod) ToMsgPack() string {
	mp, _ := msg.Marshal(m)
	return string(mp)
}

// Converts the Query into TOML
func (m Mod) ToToml() string {
	t, _ := toml.Marshal(m)
	return string(t)
}

// Converts the Query into YAML
func (m Mod) ToYaml() string {
	y, _ := yaml.Marshal(m)
	return string(y)
}

// JSON Constructor
// Creates the Query object from JSON
func (m *Mod) FromJson(s string) {
	if err := json.Unmarshal([]byte(s), m); err != nil {
		panic(err)
	}
}

// MsgPack Constructor
// Creates the Query object from MsgPack
func (m *Mod) FromMsgPack(s string) {
	if err := msg.Unmarshal([]byte(s), m); err != nil {
		panic(err)
	}
}

// TOML Constructor
// Creates the Query object from TOML
func (m *Mod) FromToml(s string) {
	if err := toml.Unmarshal([]byte(s), m); err != nil {
		panic(err)
	}
}

// YAML Constructor
// Creates the Query object from YAML
func (m *Mod) FromYaml(s string) {
	if err := yaml.Unmarshal([]byte(s), m); err != nil {
		panic(err)
	}
}
