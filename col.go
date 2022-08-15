package reportgen

type Column struct {
	Data []Cell

	Style ColumnStyle
}

type ColumnStyle struct {
	Width            float64
	HeaderCellWidth  uint // only for excel
	HeaderStyle      *CellStyle
	DefaultCellStyle *CellStyle
}

func (c *Column) width() uint {
	maxWith := uint(0)

	for _, cell := range c.Data {
		if cell.Width() > maxWith {
			maxWith = cell.Width()
		}
	}

	return maxWith
}

func (c *Column) Height() uint {
	return uint(len(c.Data))
}

func (c *Column) AddCell(cell Cell) {
	c.Data = append(c.Data, cell)
}
