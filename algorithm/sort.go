//author xin-bing
//time 10/20/2018 15:09
//排序算法，go自带的排序算法需要实现一个接口，感觉不是很便利，所以自己实现些算法
package algorithm

const (
	insertSortThreshold = 47  //数据量小于这个，那么会使用插入排序
	quickSortThreshold = 286 //如果数据量小于改常量，那么会选择快速排序
)
type sorter struct {
	length int
	compare *func(i, j int) int
	swap *func(i, j int)
}
//Sort方法，该方法会判断数据量的大小，状况而选择相应的算法进行排序，该方法接受3个参数
//length为待排数据长度
//compare是一个对数据进行比较的函数，该方法接受两个int类型的参数，即数据的下标，即待比较的两个元素下标
//如果要升序排序，如果下标i的元素比下标j的元素小则返回小于0的值
//如果要降序排序，如果下标i的元素比下标j的元素小则返回大于0的值
//swap是交换顺序
func Sort(length int, compare func(i, j int) int, swap func(i, j int)) {
	s := &sorter{
		length: length,
		compare: &compare,
		swap: &swap,
	}
	if length < insertSortThreshold {
		insertSort(1, length, compare, swap)
	}else if length < quickSortThreshold {
		quickSort(s)
	}else {
		quickSort(s)
		//mergeSort(s)
	}
}

// 归并排序
func mergeSort(s *sorter) {
}

// 快排
func quickSort(s *sorter) {
	_quickSort(0, s.length, maxDepth(s.length), s,)
}

func _quickSort(low, high, maxDepth int, s *sorter) {
	for high-low > 12 {
		if maxDepth == 0 {
			heapSort(low, high, s)
			return
		}
		maxDepth--
		mlo, mhi := choosePivot(low, high, s)
		if mlo-low < high-mhi {
			_quickSort(low, mlo, maxDepth, s)
			low = mhi // i.e., quickSort(data, mhi, b)
		}else {
			_quickSort(mhi, high, maxDepth, s)
			high = mlo // i.e., quickSort(data, a, mlo)
		}
	}
	if high-low > 1 {
		// Do ShellSort pass with gap 6
		// It could be written in this simplified form cause b-a <= 12
		for i := low + 6; i < high; i++ {
			if (*s.compare)(i, i-6) < 0 {
				(*s.swap)(i, i-6)
			}
		}
		insertSort(low, high, *s.compare, *s.swap)
	}
}

func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}

func choosePivot(lo, hi int, s *sorter) (int, int) {
	m := int(uint(lo+hi) >> 1)
	if hi-lo > 40 {
		// Tukey's ``Ninther,'' median of three medians of three.
		magic := (hi - lo) / 8
		medianOfThree(lo, lo+magic, lo+2*magic, s)
		medianOfThree(m, m-magic, m+magic, s)
		medianOfThree(hi-1, hi-1-magic, hi-1-2*magic, s)
	}
	medianOfThree(lo, m, hi-1, s)
	// Invariants are:
	//	data[lo] = pivot (set up by ChoosePivot)
	//	data[lo < i < a] < pivot
	//	data[a <= i < b] <= pivot
	//	data[b <= i < c] unexamined
	//	data[c <= i < hi-1] > pivot
	//	data[hi-1] >= pivot
	pivot := lo
	a, c := lo+1, hi-1

	for ; a < c && (*s.compare)(a, pivot) < 0; a++ {
	}
	b := a
	for {
		for ; b < c && (*s.compare)(pivot, b) >= 0; b++ { // data[b] <= pivot
		}
		for ; b < c && (*s.compare)(pivot, c-1) < 0; c-- { // data[c-1] > pivot
		}
		if b >= c {
			break
		}
		(*s.swap)(b, c-1)
		b++
		c--
	}
	// If hi-c<3 then there are duplicates (by property of median of nine).
	// Let be a bit more conservative, and set border to 5.
	protect := hi-c < 5
	if !protect && hi-c < (hi-lo)/4 {
		// Lets test some points for equality to pivot
		dups := 0
		if (*s.compare)(pivot, hi-1) >= 0 { // data[hi-1] = pivot
			(*s.swap)(c, hi-1)
			c++
			dups++
		}
		if (*s.compare)(b-1, pivot) >= 0 { // data[b-1] = pivot
			b--
			dups++
		}
		// m-lo = (hi-lo)/2 > 6
		// b-lo > (hi-lo)*3/4-1 > 8
		// ==> m < b ==> data[m] <= pivot
		if (*s.compare)(m, pivot) >= 0 { // data[m] = pivot
			(*s.swap)(m, b-1)
			b--
			dups++
		}
		// if at least 2 points are equal to pivot, assume skewed distribution
		protect = dups > 1
	}
	if protect {
		// Protect against a lot of duplicates
		// Add invariant:
		//	data[a <= i < b] unexamined
		//	data[b <= i < c] = pivot
		for {
			for ; a < b && (*s.compare)(b-1, pivot) >= 0; b-- { // data[b] == pivot
			}
			for ; a < b && (*s.compare)(a, pivot) < 0; a++ { // data[a] < pivot
			}
			if a >= b {
				break
			}
			// data[a] == pivot; data[b-1] < pivot
			(*s.swap)(a, b-1)
			a++
			b--
		}
	}
	// Swap pivot into middle
	(*s.swap)(pivot, b-1)
	return b - 1, c
}

func medianOfThree(m1, m0, m2 int, s *sorter) {
	// sort 3 elements
	if (*s.compare)(m1, m0) < 0 {
		(*s.swap)(m1, m0)
	}
	// data[m0] <= data[m1]
	if (*s.compare)(m2, m1) < 0 {
		(*s.swap)(m2, m1)
		if (*s.compare)(m1, m0) < 0 {
			(*s.swap)(m1, m0)
		}
	}
	// now data[m0] <= data[m1] <= data[m2]
}

// siftDown implements the heap property on data[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDown( lo, hi, first int, s *sorter) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && (*s.compare)(first+child, first+child+1) < 0 {
			child++
		}
		if (*s.compare)(first+root, first+child) >= 0 {
			return
		}
		(*s.swap)(first+root, first+child)
		root = child
	}
}
func heapSort(a,b int, s *sorter) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(i, hi, first, s)
	}

	// Pop elements, largest first, into end of data.
	for i := hi - 1; i >= 0; i-- {
		(*s.swap)(first, first+i)
		siftDown(lo, i, first, s)
	}
}

// 直接插入排序，适合为数据量小，基本有续的待排数据
func insertSort(low, high int, compare func(i, j int) int, swap func(i, j int)) {
	for i := low; i < high; i++ {
		for j := i; j>0 && compare(j, j-1) < 0; j-- {
			swap(j, j-1)
		}
	}
}

// 倒转序列
func reverseArr(length int, swap func(i, j int)) {
	left, right := 0, length - 1
	for left < right {
		swap(left, right)
		left++
		right--
	}
}