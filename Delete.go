package core

import (
	"encoding/json"

	toml "github.com/pelletier/go-toml/v2"
	msg "github.com/vmihailenco/msgpack/v5"
	yaml "gopkg.in/yaml.v3"
)

// Represents a request to delete specific attributes from a specific document
type Delete struct {
	// Collection that contains the document
	Collection string `json:"collection" msgpack:"collection" csv:"collection" toml:"collection" yaml:"collection"`
	// Document that contains the attributes
	Document string `json:"document" msgpack:"document" csv:"document" toml:"document" yaml:"document"`
	// Attributes to be deleted
	Values []string `json:"values" msgpack:"values" csv:"values" toml:"values" yaml:"values,flow"`
}

// Converts the Delete into JSON
func (d Delete) ToJson() string {
	js, _ := json.Marshal(d)
	return string(js)
}

// Converts the Delete into MsgPack
func (d Delete) ToMsgPack() string {
	mp, _ := msg.Marshal(d)
	return string(mp)
}

// Converts the Delete into TOML
func (d Delete) ToToml() string {
	t, _ := toml.Marshal(d)
	return string(t)
}

// Converts the Delete into YAML
func (d Delete) ToYaml() string {
	y, _ := yaml.Marshal(d)
	return string(y)
}

// JSON Constructor
// Creates the Delete object from JSON
func (d *Delete) FromJson(s string) {
	if err := json.Unmarshal([]byte(s), d); err != nil {
		panic(err)
	}
}

// MsgPack Constructor
// Creates the Delete object from MsgPack
func (d *Delete) FromMsgPack(s string) {
	if err := msg.Unmarshal([]byte(s), d); err != nil {
		panic(err)
	}
}

// TOML Constructor
// Creates the Delete object from TOML
func (d *Delete) FromToml(s string) {
	if err := toml.Unmarshal([]byte(s), d); err != nil {
		panic(err)
	}
}

// YAML Constructor
// Creates the Delete object from YAML
func (d *Delete) FromYaml(s string) {
	if err := yaml.Unmarshal([]byte(s), d); err != nil {
		panic(err)
	}
}
