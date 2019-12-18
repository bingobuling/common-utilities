//author xin-bing
//time 10/20/2018 15:09
//sort 测试
package algorithm

import (
	"testing"
	"math/rand"
	"time"
	"fmt"
)
type orderType int
const (
	orderASC  orderType = 0
	orderDESC orderType = 1
)
func TestSort(t *testing.T) {
	for i := 0; i< 1000; i++ {
		arrPointer := generateArr(2000)
		arr := *arrPointer
		Sort(len(arr), func(index1, index2 int) int {
			return arr[index2] - arr[index1]
		},
		func(index1, index2 int) {
			arr[index1], arr[index2] = arr[index2], arr[index1]
		})
		if !checkSort(arrPointer, orderDESC) { //更改了上面的排序算法，需要修改这个验证的参数
			t.Error("test sort failed")
			return
		}
	}
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
func generateArr(maxLen int) *[]int {
	length := r.Intn(maxLen) + 1
	arr := make([]int, length)
	for i:=0; i<length;i++ {
		arr[i] = r.Intn(1000)
	}
	return &arr
}

func checkSort(arr *[]int, orderType orderType) bool{
	rArr := *arr
	if orderType == orderASC {
		for i:=0; i < len(rArr) - 1; i++ {
			if rArr[i] > rArr[i + 1] {
				fmt.Println(i)
				return false
			}
		}
	} else {
		for i:=0; i < len(rArr) - 1; i++ {
			if rArr[i] < rArr[i + 1] {
				fmt.Println(i)
				return false
			}
		}
	}
	return true
}
