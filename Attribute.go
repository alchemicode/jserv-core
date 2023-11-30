package core

import (
	"fmt"
)

// Represents an individual data value within a Document in the jServ database
type Attribute struct {
	// Unique key
	Key string `json:"key" msgpack:"key" csv:"key" toml:"key" yaml:"key"`
	// Data value
	Val interface{} `json:"value" msgpack:"value" csv:"value" toml:"value" yaml:"value"`
}

// Default Constructor
func (v *Attribute) New(key string, value interface{}) {
	v.Key = key
	v.Val = value
}

// Converts the Value container to a map
func (v Attribute) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["key"] = v.Key
	m["value"] = v.Val
	return m
}

// Converts the container to a string
func (v Attribute) String() string {
	return fmt.Sprintf("{ \" %s \" : %v  }", v.Key, v.Val)
}
