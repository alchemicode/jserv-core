package core

import (
	"fmt"
)

// Represents an individual value within a Document in the jServ database
type Value struct {
	Key string      `json:"key" msgpack:"key" csv:"key" toml:"key" yaml:"key"`
	Val interface{} `json:"value" msgpack:"value" csv:"value" toml:"value" yaml:"value"`
}

// Default Constructor
func (v *Value) New(key string, value interface{}) {
	v.Key = key
	v.Val = value
}

// Converts the Value container to a map
func (v Value) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["key"] = v.Key
	m["value"] = v.Val
	return m
}

// Converts the container to a string
func (v Value) String() string {
	return fmt.Sprintf("{ \" %s \" : %v  }", v.Key, v.Val)
}
