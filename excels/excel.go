//author xinbing
//time 2018/9/14 9:34
//
package excels

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"io"
	"net/http"
	"strings"
)

type Excel struct {
	Sheets     []*Sheet
	cachedFile *excelize.File
}

func (p *Excel) Write(w io.Writer) error {
	if p.cachedFile != nil {
		return p.cachedFile.Write(w)
	}
	err := p.generateNewFile()
	if err != nil {
		return err
	}
	p.cachedFile.SetActiveSheet(0)
	return p.cachedFile.Write(w)
}

func (p *Excel) Download(header http.Header, w io.Writer, fileName string) error {
	if fileName == "" {
		fileName = "excel-exports.xlsx"
	} else if !strings.HasSuffix(fileName, ".xlsx") {
		fileName += ".xlsx"
	}
	header.Add("Content-Type", "multipart/form-data")
	header.Add("Content-Disposition", "attachment;fileName="+fileName)
	return p.Write(w)
}

func (p *Excel) generateNewFile() error {
	xlsxFile := excelize.NewFile()
	for _, sheet := range p.Sheets {
		err := sheet.fillFile(xlsxFile)
		if err != nil {
			return err
		}
	}
	p.cachedFile = xlsxFile
	return nil
}
