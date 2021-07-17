package main

import (
	"fmt"
	"log"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/jaffee/commandeer"
)

type Main struct {
	Xlsx      string `help:"要翻译内容xlsx文件路径" flag:"path"`
	SheetName string `help:"要翻译的表单" flag:"sheet"`
	Lang      string `help:"原来语言" flag:"lang"`
	Target    string `help:"翻译语言" flag:"target"`
}

func NewMain() *Main {
	return &Main{
		Xlsx:      "",
		SheetName: "",
		Lang:      "zh-CN",
		Target:    "en",
	}
}

func (m *Main) Run() error {
	if m.Xlsx == "" {
		return fmt.Errorf("确保参数正确")
	}
	return nil
}

func main() {

	conf := NewMain()
	err := commandeer.Run(conf)
	if err != nil {
		log.Fatal(err)
	}

	f, err := excelize.OpenFile(conf.Xlsx)
	if err != nil {
		log.Fatal(err)
		return
	}

	if conf.SheetName != "" {
		dealSheet(f, conf.SheetName, conf)
	} else {
		sheets := f.GetSheetList()
		for _, name := range sheets {
			dealSheet(f, name, conf)
		}
	}

	fmt.Println("正在导出表格...")

	pwd, _ := os.Getwd()
	outputPath := pwd + "/output.xlsx"

	err = f.SaveAs(outputPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("已导出表格在" + outputPath)
}

func dealSheet(f *excelize.File, sheetName string, conf *Main) {

	fmt.Println("开始处理表格：" + sheetName)

	rows, err := f.Rows(sheetName)
	if err != nil {
		log.Fatal(err)
	}

	row := 0
	for rows.Next() {
		cols, err := rows.Columns()
		if err != nil {
			log.Fatal(err)
		}

		for index, colCell := range cols {
			if colCell != "" {
				name, err := excelize.CoordinatesToCellName(index+1, row+1)
				if err == nil {
					translated, err := translate(colCell, conf.Lang, conf.Target)
					if err == nil && translated != "" {
						_ = f.SetCellValue(sheetName, name, translated)
					}

				}
			}
		}
		row++
	}

}
