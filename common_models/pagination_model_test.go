//author xinbing
//time 2018/10/13 11:39
//
package common_models

import (
	"fmt"
	"testing"
)

type testPagination struct {
	Z string
	Pagination
}

func TestPagination_Offset(t *testing.T) {
	z := testPagination{}
	z.Z = "123"
	z.Offset()
	//z.CalCurrCapacity(10)
	fmt.Println(z)
	b := BuildPagination(1, 10, 101)
	fmt.Println(b)
}
