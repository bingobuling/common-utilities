//author xinbing
//time 2018/9/5 13:54
package utilities

import (
	"fmt"
	"testing"
)

func TestValidPhone(t *testing.T) {
	fmt.Println(ValidPhone("17417771777"))
}

func TestValidEmail(t *testing.T) {
	fmt.Println(ValidEmail("bin-g.xin@cnlaunch.com-a.cn.bb"))
}

func TestValidUrl(t *testing.T) {
	fmt.Println(ValidUrl("http://f_1-1.f121.f1212"))
	fmt.Println(ValidUrl("https://f1.f"))
	fmt.Println(ValidUrl("ws://f1.f"))
	fmt.Println(ValidUrl("ws://12.1"))
	fmt.Println(ValidUrl("ws://11.2"))
	fmt.Println(ValidUrl("ws://11.2"))
}
func TestPwdGrade(t *testing.T) {
}
