package reportgen

import (
	"fmt"

	"github.com/mandolyte/mdtopdf"
)

func (report *ReportFile) ToPDF(output string) error {
	pf := mdtopdf.NewPdfRenderer("", "", output, "trace.log")
	pf.Pdf.AddFont("Helvetica-1251", "", "helvetica_1251.json")
	pf.Pdf.SetFont("Helvetica-1251", "", 12)
	// get the unicode translator
	tr := pf.Pdf.UnicodeTranslatorFromDescriptor("cp1251")
	pf.Normal = mdtopdf.Styler{
		Font: "Helvetica-1251", Style: "",
		Size: 10, Spacing: 2,
		TextColor: mdtopdf.Color{Red: 0, Green: 0, Blue: 0},
		FillColor: mdtopdf.Color{Red: 255, Green: 255, Blue: 255},
	}

	err := pf.Process([]byte(tr(report.ToMD())))
	if err != nil {
		return fmt.Errorf("pdf.OutputFileAndClose() error:%v", err)
	}

	return nil
}

// func genPDF(report *ReportFile) pdf.Maroto {
// 	m := pdf.NewMaroto(consts.Landscape, consts.Letter)

// 	// TODO to config
// 	m.AddUTF8Font("DroidSans", consts.Normal, "fonts/droid/DroidSans.ttf")
// 	m.AddUTF8Font("DroidSans", consts.Italic, "fonts/droid/DroidSans.ttf")
// 	m.AddUTF8Font("DroidSans", consts.Bold, "fonts/droid/DroidSans-Bold.ttf")
// 	m.AddUTF8Font("DroidSans", consts.BoldItalic, "fonts/droid/DroidSans-Bold.ttf")
// 	m.SetDefaultFontFamily("DroidSans")

// 	pageWidth, _ := m.GetPageSize()
// 	maxW := uint(pageWidth)

// 	for _, sheet := range report.Sheets {
// 		m.Row(8, func() {
// 			m.Col(maxW, func() {
// 				m.Text(sheet.Title)
// 			})
// 		})
// 		m.Line(10)

// 		for _, table := range sheet.Tables {
// 			m.Row(10, func() {
// 				m.Col(1, func() {
// 					m.Text(table.Title)
// 				})
// 			})

// 			m.SetBorder(true)

// 			m.Row(10, func() {
// 				for _, colName := range table.Columns {
// 					m.Col(1, func() {
// 						m.Text(colName)
// 					})
// 				}
// 			})

// 			table.iterRows(func(row []Cell) bool {
// 				m.Row(10, func() {
// 					for _, c := range row {
// 						for _, s := range c.Strings() {
// 							m.Col(1, func() {
// 								m.Text(s)
// 							})
// 						}
// 					}

// 				})
// 				return true
// 			})
// 			m.SetBorder(false)
// 		}
// 		m.AddPage()
// 	}

// 	return m
// }
