//author bing.xin
//time  2018/10/23 17:19
//desc 
package versions

import (
	"fmt"
	"os"
)

var (
	BuildVersion	string
	BuildTime		string
)

func init() {
	args := os.Args
	if nil == args || len(args) < 2 {
		return
	}
	if "version" == args[1] {
		fmt.Printf("version:%v,buildTime:%v\n", BuildVersion, BuildTime)
	}
}
