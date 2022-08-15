package reportgen

import (
	"strconv"

	"github.com/xuri/excelize/v2"
)

// must start with 1
type Pos struct {
	Col uint
	Row uint
}

func (p Pos) ToExcel() string {
	cell, err := excelize.CoordinatesToCellName(int(p.Col), int(p.Row))
	if err != nil {
		panic(err)
	}
	return cell
}

func (p Pos) ExcelCol() string {
	col, err := excelize.ColumnNumberToName(int(p.Col))
	if err != nil {
		panic(err)
	}
	return col
}

func (p Pos) ExcelRow() string {
	return strconv.Itoa(int(p.Row))
}

func (p Pos) Add(col, row int) Pos {
	if int(p.Col)+col < 0 || int(p.Row)+row < 0 {
		panic("invalid pos modify")
	}

	return Pos{Col: uint(int(p.Col) + col), Row: uint(int(p.Row) + row)}
}
