//author xinbing
//time 2018/9/14 9:35
//
package excels

type ColumnDefine struct {
	Title          string
	Width          float64
	ValueExtractor func(data interface{}) interface{}
}
