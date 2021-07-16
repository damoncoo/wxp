package main

import (
	"fmt"
	"log"

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
	if m.SheetName == "" || m.Xlsx == "" {
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

	rows, err := f.Rows(conf.SheetName)
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
					if err != nil && translated != "" {
						_ = f.SetCellValue(conf.SheetName, name, translated)
					}

				}
			}
		}
		row++
	}

	err = f.SaveAs("output.xlsx")
	if err != nil {
		fmt.Println(err)
	}

}

// func translate(text string) string {

// 	fmt.Println("正在翻译：" + text)

// 	translated, err := gt.Translate(text, "auto", "en")
// 	fmt.Println(translated)

// 	if err != nil {
// 		return ""
// 	}

// 	return translated
// }
