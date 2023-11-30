package core

import (
	msg "github.com/vmihailenco/msgpack/v5"
)

// Represents a grouping of documents within the jServ database
type Collection struct {
	// Collection id and filename
	Name string
	// Pointers to Documents in memory
	List []*Document
}

// Named Constructor
// Creates an empty Collection object with just a name
func (c *Collection) New(name string) {
	c.Name = name
	c.List = make([]*Document, 0)
}

// Reads an entire collection from MsgPack
func (c *Collection) FromMsgPack(b []byte) bool {
	//Generates map data from the file
	var dat []map[interface{}]interface{}
	if err := msg.Unmarshal(b, &dat); err != nil {
		return false
	} else {
		//Reads each document in the generated data from the file
		//and populates the collection's list
		for _, m := range dat {
			doc := new(Document)
			doc.Id = m["_id"].(string)
			doc.Data = m["data"].(map[string]interface{})
			c.List = append(c.List, doc)
		}
		//Channel returns true if the read was successful
		return true
	}
}

// Converts the collection in full to MsgPack
func (c Collection) ToMsgPack() ([]byte, bool) {
	var vals []Document
	for _, d := range c.List {
		vals = append(vals, *d)
	}
	bytes, err := msg.Marshal(vals)
	return bytes, (err == nil)

}
