package core

import (
	"encoding/json"

	toml "github.com/pelletier/go-toml/v2"
	msg "github.com/vmihailenco/msgpack/v5"
	yaml "gopkg.in/yaml.v3"
)

// Represents an addition of a whole document, or a value within an existing document, to the jServ database
type Add struct {
	// Collection to be added to
	Collection string `json:"collection" msgpack:"collection" csv:"collection" toml:"collection" yaml:"collection"`
	// Whole documents to add to the collection
	Documents []Document `json:"documents,omitempty" msgpack:"documents,omitempty" csv:"documents,omitempty" toml:"documents,omitempty" yaml:"documents,omitempty,flow"`
	// Attributes to be added to speciic documents organized by doc_id
	Values map[string]interface{} `json:"values,omitempty" msgpack:"values,omitempty" csv:"values,omitempty" toml:"values,omitempty" yaml:"values,omitempty,flow"`
}

// Converts the Add into JSON
func (a Add) ToJson() string {
	js, _ := json.Marshal(a)
	return string(js)
}

// Converts the Add into MsgPack
func (a Add) ToMsgPack() string {
	m, _ := msg.Marshal(a)
	return string(m)
}

// Converts the Add into TOML
func (a Add) ToToml() string {
	t, _ := toml.Marshal(a)
	return string(t)
}

// Converts the Add into YAML
func (a Add) ToYaml() string {
	y, _ := yaml.Marshal(a)
	return string(y)
}

// JSON Constructor
// Creates the Add object from JSON
func (a *Add) FromJson(s string) {
	if err := json.Unmarshal([]byte(s), a); err != nil {
		panic(err)
	}
}

// MsgPack Constructor
// Creates the Add object from MsgPack
func (a *Add) FromMsgPack(s string) {
	if err := msg.Unmarshal([]byte(s), a); err != nil {
		panic(err)
	}
}

// TOML Constructor
// Creates the Add object from JSON
func (a *Add) FromToml(s string) {
	if err := toml.Unmarshal([]byte(s), a); err != nil {
		panic(err)
	}
}

// YAML Constructor
// Creates the Add object from YAML
func (a *Add) FromYaml(s string) {
	if err := yaml.Unmarshal([]byte(s), a); err != nil {
		panic(err)
	}
}
