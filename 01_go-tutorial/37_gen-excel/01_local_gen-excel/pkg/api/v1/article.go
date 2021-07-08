package v1

import (
	"example/gen-excel/pkg/export"
	"example/gen-excel/pkg/file"
	"fmt"
	"github.com/tealeg/xlsx"
	"strconv"
	time2 "time"
)

type Tag struct {
	ID       int
	Name     string
	CreateBy string
}

func (t *Tag) GetAll() []Tag {
	return []Tag{
		{
			ID:       1,
			Name:     "电视剧",
			CreateBy: "gao",
		},
		{
			ID:       2,
			Name:     "电影",
			CreateBy: "gao",
		},
	}
}

func (t *Tag) Export() (string, error) {
	tags := t.GetAll()

	fileXlsx := xlsx.NewFile()
	sheet, err := fileXlsx.AddSheet("标签信息")
	if err != nil {
		return "", err
	}

	titles := []string{"ID", "名称", "创建人"}
	row := sheet.AddRow()

	var cell *xlsx.Cell
	for _, title := range titles {
		cell = row.AddCell()
		cell.Value = title
	}

	for _, v := range tags {
		values := []string{
			strconv.Itoa(v.ID),
			v.Name,
			v.CreateBy,
		}

		row = sheet.AddRow()

		for _, value := range values {
			cell = row.AddCell()
			cell.Value = value
		}
	}

	time := strconv.Itoa(int(time2.Now().Unix()))
	filename := "tags-" + time + ".xlsx"

	//创建目录
	err = file.MustOpen(export.GetExcelFullPath())
	if err != nil {
		fmt.Println("err")
	}

	fullPath := export.GetExcelFullPath() + filename

	err = fileXlsx.Save(fullPath)
	if err != nil {
		return "", err
	}

	return filename, nil
}
