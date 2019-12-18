//author xinbing
//time 2018/8/28 14:18
//数字工具
package utilities

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var fmtStrings = []string{
	//第一个是补位
	"%0.0f", "%0.1f", "%0.2f", "%0.3f", "%0.4f", "%0.5f", "%0.6f", "%0.7f", "%0.8f", "%0.9f", "%0.10f",
}

// 四舍五入
func Round(f float64, precision int) float64 {
	if precision <= 0 {
		return math.Round(f)
	}
	var fmtStr string
	if precision <= 10 {
		fmtStr = fmtStrings[precision]
	} else {
		fmtStr = "%0." + strconv.Itoa(precision) + "f"
	}
	s := fmt.Sprintf(fmtStr, f)
	nf, _ := strconv.ParseFloat(s, 64)
	return nf
}

// 向上舍入
func Floor(f float64, precision int) float64 {
	if precision <= 0 {
		return math.Floor(f)
	}
	pow := math.Pow10(precision)
	nf := f * pow
	nf = math.Floor(nf) / pow
	return Round(nf, precision)
}

// 向下舍去
func Ceil(f float64, precision int) float64 {
	if precision <= 0 {
		return math.Ceil(f)
	}
	pow := math.Pow10(precision)
	nf := f * pow
	nf = math.Ceil(nf) / pow
	return Round(nf, precision)
}

// 截取
func Trunc(f float64, precision int) float64 {
	if precision <= 0 {
		return math.Trunc(f)
	}
	pow := math.Pow10(precision)
	nf := f * pow
	nf = math.Trunc(nf) / pow
	return Round(nf, precision)
}

const (
	limit = 0.000000001
)

// compare
func Compare(f1 float64, f2 float64) int {
	r := math.Abs(f1 - f2)
	if r <= limit {
		return 0
	} else if f1 < f2 {
		return -1
	} else {
		return 1
	}
}

//指定比较的精度
func CompareWithScale(f1, f2 float64, scale int) int {
	if scale > 0 {
		scale = -scale
	}
	limit := math.Pow10(scale - 1)
	r := f1 - f2
	if math.Abs(r) <= limit {
		return 0
	} else if r > limit {
		return 1
	} else {
		return -1
	}
}

// 保留两位小数的浮点数
func TruncateTwoDecimals(s float64) (float64, error) {
	sTmp := strings.Split(fmt.Sprintf("%f", s), ".")
	if len(sTmp) != 2 {
		return 0, fmt.Errorf("it's not float string")
	}

	out, err := strconv.ParseFloat(sTmp[0]+"."+sTmp[1][0:2], 64)
	if err != nil {
		return 0, err
	}
	return out, nil
}
