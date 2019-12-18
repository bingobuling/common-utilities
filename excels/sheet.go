//author xinbing
//time 2018/9/14 9:35
//
package excels

import (
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize"
	"reflect"
	"strconv"
	"strings"
)

type Sheet struct {
	Name 			string
	BeginRow		int
	BeginCol		int
	ColumnDefines 	[]*ColumnDefine
	DataSource      interface{} //要导出的数据
	data       		[]interface{} //DataSource进行变换
	handlingRow		int
}

func (p *Sheet) fillFile(xlsxFile *excelize.File) error {
	err := p.check()
	if err != nil {
		return err
	}
	xlsxFile.NewSheet(p.Name)
	p.fillTitleRow(xlsxFile)
	for _,item := range p.data {
		values := make([]interface{}, len(p.ColumnDefines))
		for index, col := range p.ColumnDefines {
			values[index] = col.ValueExtractor(item)
		}
		p.fillRow(values, xlsxFile)
	}
	return nil
}

func (p *Sheet) check() error {
	if p.Name = strings.Trim(p.Name, " "); len(p.Name) == 0 {
		p.Name = "Sheet1"
	}
	if p.BeginCol < 0 {
		p.BeginCol = 0
	}
	if p.BeginRow <0 {
		p.BeginRow = 0
	}
	p.handlingRow = p.BeginRow
	dataSlice,err := p.convertToSlice(p.DataSource)
	if err != nil {
		return err
	}
	p.data = dataSlice
	if p.ColumnDefines == nil || len(p.ColumnDefines) == 0 {
		return errors.New("sheet ColumnDefine can not be empty")
	}
	for _,item := range p.ColumnDefines {
		if item.Width <= 0 {
			item.Width = 10
		}
	}
	return nil
}

func (p *Sheet) convertToSlice(data interface{}) ([]interface{}, error) {
	value := reflect.ValueOf(data)
	if value.Kind() != reflect.Slice {
		return nil, errors.New("sheet DataSource must be a slice")
	}
	l := value.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = value.Index(i).Interface()
	}
	return ret, nil
}

func (p *Sheet) fillTitleRow(xlsx *excelize.File) {
	p.handlingRow++ //因为handlingRow是从0开始的，而依赖的excelize从1开始的
	strRowNum := strconv.Itoa(p.handlingRow)
	for index,item := range p.ColumnDefines {
		col := string('A' + p.BeginCol + index)
		axis := col + strRowNum
		xlsx.SetColWidth(p.Name,col,col, item.Width)
		xlsx.SetCellValue(p.Name, axis, item.Title)
	}
}

func (p *Sheet) fillRow(values []interface{},xlsx *excelize.File) {
	p.handlingRow++
	strRowNum := strconv.Itoa(p.handlingRow)
	for index,value := range values {
		axis := string('A' + p.BeginCol + index) + strRowNum
		xlsx.SetCellValue(p.Name, axis, value)
	}
}