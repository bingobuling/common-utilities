//author xinbing
//time 2018/10/16 10:35
// 基础算法套件
package algorithm

/*
	二分查找法的蛋疼实现，如果找到则返回数组下标，否则返回-1
	待查找的数组必须根据compare方法里所比较的元素进行排序，否则二分查找法是不会生效的
	length为待查找数组长度
	compare方法为比较方法，参数index为数组下标，方法的实现内容为用下标index的元素与待查找元素进行比较
*/
func BinarySearch(length int, compare func (index int) int) int {
	begin, last, mid := 0, length - 1, -1
	for begin <= last {
		mid = (begin + last) / 2
		rs := compare(mid)
		if rs == 0 {
			return mid
		} else if rs > 0 {
			last = mid -1
		} else {
			begin = mid + 1
		}
	}
	return -1
}

// 直接for循环查找，如果找到则返回数组下标，否则返回-1
func FlatSearch(length int, compare func(index int) bool) int {
	for i:=0; i<length; i++ {
		if compare(i) {
			return i
		}
	}
	return -1
}