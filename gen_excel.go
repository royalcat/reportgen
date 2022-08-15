package reportgen

import (
	"github.com/royalcat/gobatteries"
	"github.com/xuri/excelize/v2"
)

func (r *ReportFile) ToExcel() *excelize.File {

	g := excelGen{
		F:         excelize.NewFile(),
		StylesMap: make(map[*excelize.Style]int),
	}
	f := g.F

	props := &excelize.DocProperties{
		Title:   r.Title,
		Creator: r.Creator,
	}

	f.SetDocProps(props)
	nameStyle := g.getExcelStyle(
		&excelize.Style{
			Alignment: &excelize.Alignment{
				Horizontal: "center",
				Vertical:   "center",
			},
			Font: &excelize.Font{
				Bold: true,
			},
		})

	for _, sheet := range r.Sheets {
		f.SetDocProps(props)
		f.NewSheet(sheet.Name)

		// Шапка отчета
		f.MergeCell(sheet.Name, "A1", "R1")
		f.SetRowHeight(sheet.Name, 1, 30)
		f.SetCellStr(sheet.Name, "A1", sheet.Title)
		f.SetCellStyle(sheet.Name, "A1", "R1", nameStyle)

		rowI := 3

		for _, table := range sheet.Tables {
			pos := Pos{Col: 1, Row: uint(rowI)}

			// заголовок таблицы
			f.MergeCell(sheet.Name, pos.ToExcel(), pos.Add(int(table.Width()-1), 0).ToExcel())
			f.SetCellStyle(sheet.Name, pos.ToExcel(), pos.Add(int(table.Width()-1), 0).ToExcel(), nameStyle)
			f.SetRowHeight(sheet.Name, int(pos.Row), 30)
			f.SetCellStr(sheet.Name, pos.ToExcel(), table.Title)
			rowI++

			g.writeExcelTable(sheet.Name, pos.Add(0, 1), table)
			rowI += int(table.Height())

			rowI += 1 + sheet.TableOffset
		}
	}

	f.DeleteSheet("Sheet1")

	return f
}

type excelGen struct {
	F         *excelize.File
	StylesMap map[*excelize.Style]int
}

func (g *excelGen) writeExcelTable(sheet string, startPos Pos, t Table) {
	f := g.F

	f.SetRowHeight(sheet, int(startPos.Row), 40)

	colOffset := startPos.Col - 1
	colums := t.Columns
	if len(colums) == 0 || colums == nil {
		colums = gobatteries.KeysOfMap(t.data)
	}

	for _, name := range colums {
		col, ok := t.data[name]
		if !ok {
			col = &Column{}
		}

		colCellWidth := 1
		if col.width() > 1 {
			colCellWidth = int(col.width())
		} else if int(col.Style.HeaderCellWidth) > colCellWidth {
			colCellWidth = int(col.Style.HeaderCellWidth)
		}

		headerPos := startPos.Add(int(colOffset), 0)

		if style := gobatteries.FirstNotNil(col.Style.HeaderStyle, t.DefaultHeaderStyle); style != nil {
			f.SetCellStyle(sheet, headerPos.ToExcel(), headerPos.ToExcel(), g.getExcelStyle(style.ToExcel()))
		}
		if colCellWidth != 1 {
			f.MergeCell(sheet, headerPos.ToExcel(), headerPos.Add(colCellWidth-1, 0).ToExcel())
		}
		if col.Style.Width != 0 {
			f.SetColWidth(sheet, headerPos.ExcelCol(), headerPos.Add(colCellWidth-1, 0).ExcelCol(), col.Style.Width)
		}

		f.SetCellValue(sheet, headerPos.ToExcel(), name)

		rowStartPos := headerPos.Add(0, 1)

		for i, cell := range col.Data {
			rowPos := rowStartPos.Add(0, i)
			if len(cell.Data) == 1 && colCellWidth != 1 {
				f.MergeCell(sheet, rowPos.ToExcel(), rowPos.Add(colCellWidth-1, 0).ToExcel())
			}

			for j, c := range cell.Data {
				cellPos := rowPos.Add(j, 0)

				if style := gobatteries.FirstNotNil(cell.Style, col.Style.DefaultCellStyle, t.DefaultCellStyle); style != nil {
					f.SetCellStyle(sheet, cellPos.ToExcel(), cellPos.ToExcel(), g.getExcelStyle(style.ToExcel()))
				}

				f.SetCellValue(sheet, cellPos.ToExcel(), c)
			}

		}

		colOffset += uint(colCellWidth)
	}

	//endCell, _ := excelize.CoordinatesToCellName(int(t.Width()), int(startPos.Row+t.Height())-1)

	//f.AddTable(sheet, startPos.ToExcel(), endCell, "")
}

func (g *excelGen) getExcelStyle(style *excelize.Style) int {
	if style == nil {
		return 0
	}

	if styleId, ok := g.StylesMap[style]; ok {
		return styleId
	}

	styleId, err := g.F.NewStyle(style)
	if err != nil {
		panic(err)
	}
	g.StylesMap[style] = styleId

	return styleId
}
