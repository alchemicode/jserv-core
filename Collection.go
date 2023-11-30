package core

import (
	"fmt"

	m "github.com/vmihailenco/msgpack/v5"
)

// Represents a grouping of documents within the jServ database
type Collection struct {
	Name string
	List []*Document
}

// Named Constructor Constructor
// Creates an empty Collection object with just a name
// and reads its file
func (c *Collection) New(name string) {
	c.Name = name
	c.List = make([]*Document, 0)
}

func (c *Collection) FromMsgPack(b []byte) bool {
	//Generates map data from the file
	var dat []map[string]interface{}
	if err := m.Unmarshal(b, &dat); err != nil {
		fmt.Println("Error when generating data for " + c.Name + ".dat")
		return false
	} else {
		//Reads each object in the generated data from the file
		//and populates the collection's list
		for i := 0; i < len(dat); i++ {
			obj := new(Document)
			obj.WithData(dat[i]["id"].(string), dat[i]["data"].(map[string]interface{}))
			c.List = append(c.List, obj)
		}
		//Channel returns true if the read was successful
		return true
	}
}

func (c Collection) ToMsgPack() ([]byte, bool) {
	var vals []Document
	for _, d := range c.List {
		vals = append(vals, *d)
	}
	bytes, err := m.Marshal(vals)
	return bytes, (err == nil)

}
