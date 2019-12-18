//author xinbing
//time 2018/9/14 9:01
//
package excels

import (
	"fmt"
	"github.com/bingobuling/common-utilities/common_models"
	"reflect"
	"testing"
)

func TestF(t *testing.T) {
	resps := []*common_models.Resp{
		{
			Code: 0,
		},
		{
			Code: 1,
		},
	}
	zz := D(resps)
	fmt.Println(zz[0].(*common_models.Resp))
}

func D(list interface{}) []interface{} {
	value := reflect.ValueOf(list)
	if value.Kind() != reflect.Slice {
		panic("to slice arr not slice")
	}
	l := value.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = value.Index(i).Interface()
	}
	return ret
}
