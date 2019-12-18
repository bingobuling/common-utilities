//author xinbing
//time 2018/9/4 14:04
package http_utils

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestPost(t *testing.T) {
	var m *map[string]string = nil
	fmt.Println((*m)["12344"])
}

type zz struct {
	H 	 string `json:"h"`
	Body []byte `json:"body"`
}
func TestBody(t *testing.T) {
	f := &f{
		TT: "1",
		CC: 2,
	}
	byt,_ := json.Marshal(f)
	z := &zz{
		H: "111",
		Body: byt,
	}
	byt,_ = json.Marshal(z)
	nz := &zz{}
	json.Unmarshal(byt, nz)
	fmt.Println(string(nz.Body))
}

type f struct {
	TT string `json:"tt"`
	CC int `json:"cc"`
}

func TestQuery(t *testing.T) {
	zf := f{
		TT: "12",
		CC: 2,
	}
	byt,err := json.Marshal(&zf)
	fmt.Println(byt,err)
	m := make(map[string]interface{})
	json.Unmarshal(byt, &m)
	fmt.Println(m)
	func(b []byte){
		fmt.Println(len(b))
	}(nil)
}

func TestGetQueryStr(t *testing.T) {
	fmt.Println(GetQueryStr("http://fdsfs.com?fdsfsfd"))
}

