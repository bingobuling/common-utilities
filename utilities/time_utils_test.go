//author xinbing
//time 2018/8/30 21:04
package utilities

import (
	"fmt"
	"testing"
	"time"
)

func TestT(t *testing.T) {
	fmt.Println("z:",time.Now().UnixNano() / int64(time.Millisecond))
	fmt.Println("b:", time.Now().Unix())
}