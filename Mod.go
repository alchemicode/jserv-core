package core

import (
	"encoding/json"

	toml "github.com/pelletier/go-toml/v2"
	msg "github.com/vmihailenco/msgpack/v5"
	yaml "gopkg.in/yaml.v3"
)

// Represents a modification to a specific document within the jServ database
type Mod struct {
	// Collection that contains the document
	Collection string `json:"collection" msgpack:"collection" csv:"collection" toml:"collection" yaml:"collection"`
	// Document that contains the values
	Document string `json:"document" msgpack:"document" csv:"document" toml:"document" yaml:"document"`
	// Attributes to be modified, and their new values
	Values map[string]interface{} `json:"values" msgpack:"values" csv:"values" toml:"values" yaml:"values,flow"`
}

// Converts the Mod into JSON
func (m Mod) ToJson() string {
	js, _ := json.Marshal(m)
	return string(js)
}

// Converts the Mod into MsgPack
func (m Mod) ToMsgPack() string {
	mp, _ := msg.Marshal(m)
	return string(mp)
}

// Converts the Mod into TOML
func (m Mod) ToToml() string {
	t, _ := toml.Marshal(m)
	return string(t)
}

// Converts the Mod into YAML
func (m Mod) ToYaml() string {
	y, _ := yaml.Marshal(m)
	return string(y)
}

// JSON Constructor
// Creates the Mod object from JSON
func (m *Mod) FromJson(s string) {
	if err := json.Unmarshal([]byte(s), m); err != nil {
		panic(err)
	}
}

// MsgPack Constructor
// Creates the Mod object from MsgPack
func (m *Mod) FromMsgPack(s string) {
	if err := msg.Unmarshal([]byte(s), m); err != nil {
		panic(err)
	}
}

// TOML Constructor
// Creates the Mod object from TOML
func (m *Mod) FromToml(s string) {
	if err := toml.Unmarshal([]byte(s), m); err != nil {
		panic(err)
	}
}

// YAML Constructor
// Creates the Mod object from YAML
func (m *Mod) FromYaml(s string) {
	if err := yaml.Unmarshal([]byte(s), m); err != nil {
		panic(err)
	}
}
