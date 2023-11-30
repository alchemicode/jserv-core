package core

import (
	"encoding/json"

	toml "github.com/pelletier/go-toml/v2"
	msg "github.com/vmihailenco/msgpack/v5"
	yaml "gopkg.in/yaml.v3"
)

// Represents a series of expressions that fetch documents meeting its conditions from the jServ database
type Query struct {
	// Collections to be queried
	Collections []string `json:"collection,omitempty" msgpack:"collection,omitempty" csv:"collection,omitempty" toml:"collection,omitempty" yaml:"collection,omitempty,flow"`
	// Attributes a document must have
	Has []string `json:"has,omitempty" msgpack:"has,omitempty" csv:"has,omitempty" toml:"has,omitempty" yaml:"has,omitempty,flow"`
	// Attributes that must be equal to a given value
	Equals map[string]interface{} `json:"equals,omitempty" msgpack:"equals,omitempty" csv:"equals,omitempty" toml:"equals,omitempty" yaml:"equals,omitempty,flow"`
	// Attributes that must not be equal to a given value
	NotEquals map[string]interface{} `json:"not-equals,omitempty" msgpack:"not-equals,omitempty" csv:"not-equals,omitempty" toml:"not-equals,omitempty" yaml:"not-equals,omitempty,flow"`
	// Attributes that must be less than a given value
	LessThan map[string]interface{} `json:"lt,omitempty" msgpack:"lt,omitempty" csv:"lt,omitempty" toml:"lt,omitempty" yaml:"lt,omitempty,flow"`
	// Attributes that must be less than or equal to a given value
	LessThanEqualTo map[string]interface{} `json:"lte,omitempty" msgpack:"lte,omitempty" csv:"lte,omitempty" toml:"lte,omitempty" yaml:"lt,omitempty,flowe"`
	// Attributes that must be greater than or equal to a given value
	GreaterThanEqualTo map[string]interface{} `json:"gte,omitempty" msgpack:"gte,omitempty" csv:"gte,omitempty" toml:"gte,omitempty" yaml:"gte,omitempty,flow"`
	// Attributes that must be greater than a given value
	GreaterThan map[string]interface{} `json:"gt,omitempty" msgpack:"gt,omitempty" csv:"gt,omitempty" toml:"gt,omitempty" yaml:"gt,omitempty,flow"`
	// Attributes that must be between two given values
	Between map[string]interface{} `json:"between,omitempty" msgpack:"between,omitempty" csv:"between,omitempty" toml:"between,omitempty" yaml:"between,omitempty,flow"`
}

// Creates query that will return all documents in the db
func (q *Query) AllQuery() {
	q.Has = []string{"_id"}
}

// Creates query that will return all documents in the given collections
func (q *Query) CollectionAllQuery(collections []string) {
	q.Collections = collections
	q.Has = []string{"_id"}
}

// Converts the Query into JSON
func (q Query) ToJson() string {
	js, _ := json.Marshal(q)
	return string(js)
}

// Converts the Query into MsgPack
func (q Query) ToMsgPack() string {
	m, _ := msg.Marshal(q)
	return string(m)
}

// Converts the Query into TOML
func (q Query) ToToml() string {
	t, _ := toml.Marshal(q)
	return string(t)
}

// Converts the Query into YAML
func (q Query) ToYaml() string {
	y, _ := yaml.Marshal(q)
	return string(y)
}

// JSON Constructor
// Creates the Query object from JSON
func (q *Query) FromJson(s string) {
	if err := json.Unmarshal([]byte(s), q); err != nil {
		panic(err)
	}
}

// MsgPack Constructor
// Creates the Query object from MsgPack
func (q *Query) FromMsgPack(s string) {
	if err := msg.Unmarshal([]byte(s), q); err != nil {
		panic(err)
	}
}

// TOML Constructor
// Creates the Query object from TOML
func (q *Query) FromToml(s string) {
	if err := toml.Unmarshal([]byte(s), q); err != nil {
		panic(err)
	}
}

// YAML Constructor
// Creates the Query object from YAML
func (q *Query) FromYaml(s string) {
	if err := yaml.Unmarshal([]byte(s), q); err != nil {
		panic(err)
	}
}
