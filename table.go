package reportgen

import (
	"github.com/royalcat/gobatteries"
)

type Table struct {
	Title   string
	Columns []string

	DefaultHeaderStyle *CellStyle
	DefaultCellStyle   *CellStyle
	ColumnsStyles      map[string]ColumnStyle

	data map[string]*Column
}

type tableData map[string]*Column

func (t *Table) GetCol(name string) *Column {
	if c, ok := t.data[name]; ok {
		return c
	}
	return &Column{}
}

func (t *Table) Height() uint {
	maxHeight := uint(0)

	for _, col := range t.data {
		if col.Height() > maxHeight {
			maxHeight = col.Height()
		}
	}
	return maxHeight
}

func (t *Table) Width() uint {
	width := uint(0)
	for _, col := range t.data {
		if col.width() > col.Style.HeaderCellWidth {
			width += col.width()
		} else if col.Style.HeaderCellWidth != 0 {
			width += col.Style.HeaderCellWidth
		}
	}

	if int(width) < len(t.Columns) {
		width = uint(len(t.Columns))
	}

	return width
}

func (t *Table) AddRow(row map[string]Cell) {
	if t.data == nil {
		t.data = make(map[string]*Column)
	}

	for k, v := range row {
		if _, ok := t.data[k]; !ok {
			if gobatteries.InSlice(t.Columns, k) {
				style, _ := t.ColumnsStyles[k]

				t.data[k] = &Column{Style: style}
			} else {
				panic("unknown column: " + k)
			}
		}
		t.data[k].AddCell(v)
	}
}

func (t *Table) SetCol(name string, row []Cell) {
	style, _ := t.ColumnsStyles[name]

	t.data[name] = &Column{Data: row, Style: style}
}

func (t *Table) GetStyle(name string) ColumnStyle {
	if style, ok := t.ColumnsStyles[name]; ok {
		return style
	}

	return ColumnStyle{}
}

func (t *Table) iterRows(iter func(row []Cell) bool) {
	if t.data == nil {
		t.data = make(map[string]*Column)
	}

	h := t.Height()
	for i := uint(0); i < h; i++ {
		row := make([]Cell, 0, len(t.Columns))
		for _, colName := range t.Columns {
			row = append(row, t.data[colName].Data[i])
		}
		if !iter(row) {
			return
		}
	}
}
