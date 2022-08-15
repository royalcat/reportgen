package reportgen

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type CellStyle struct {
	Border       BordersStyle
	Font         FontStyle
	Fill         FillStyle
	TextAlign    TextAlign
	NumberFormat NumberFormat

	excelCache *excelize.Style
}

type BordersStyle struct {
	Top, Right, Bottom, Left BorderStyle
}

type BorderPattern uint

const (
	BorderPatternNone BorderPattern = iota
	BorderPatternBold
)

type BorderStyle struct {
	Pattern BorderPattern
	Color   RGB
}

func NewBorderStyleAll(color RGB, pattern BorderPattern) BordersStyle {
	return BordersStyle{
		Top:    BorderStyle{Color: color, Pattern: pattern},
		Right:  BorderStyle{Color: color, Pattern: pattern},
		Bottom: BorderStyle{Color: color, Pattern: pattern},
		Left:   BorderStyle{Color: color, Pattern: pattern},
	}
}

type Alignment uint

const (
	AlignmentNone Alignment = iota
	AlignmentStart
	AlignmentCenter
	AlignmentEnd
)

func (a *Alignment) ToExcelHorizontal() string {
	switch *a {
	case AlignmentNone:
		return ""
	case AlignmentStart:
		return "left"
	case AlignmentCenter:
		return "center"
	case AlignmentEnd:
		return "right"
	}
	return ""
}

func (a *Alignment) ToExcelVertical() string {
	switch *a {
	case AlignmentNone:
		return ""
	case AlignmentStart:
		return "top"
	case AlignmentCenter:
		return "center"
	case AlignmentEnd:
		return "bottom"
	}
	return ""
}

type TextAlign struct {
	Horizontal, Vertical Alignment
}

type FontStyle struct {
	Bold, Italic bool
	Size         float64
	Color        RGB
}

type FillPattern uint

const (
	FillPatternNone FillPattern = iota
	FillPatternSolid
)

type FillStyle struct {
	Color   RGB
	Pattern FillPattern
}

func (s *CellStyle) ToExcel() *excelize.Style {
	if s.excelCache != nil {
		return s.excelCache
	}

	return &excelize.Style{
		Border:    s.Border.ToExcel(),
		Font:      s.Font.ToExcel(),
		Fill:      s.Fill.ToExcel(),
		Alignment: s.TextAlign.ToExcel(),
		NumFmt:    s.NumberFormat.ToExcel(),
	}
}

func (a *TextAlign) ToExcel() *excelize.Alignment {
	if a == nil {
		return nil
	}

	return &excelize.Alignment{
		Horizontal: a.Horizontal.ToExcelHorizontal(),
		Vertical:   a.Vertical.ToExcelVertical(),
		WrapText:   true,
	}
}

func (b *BordersStyle) ToExcel() []excelize.Border {
	return []excelize.Border{
		{Type: "top", Color: b.Top.Color.ToHex(), Style: int(b.Top.Pattern)},
		{Type: "right", Color: b.Right.Color.ToHex(), Style: int(b.Right.Pattern)},
		{Type: "bottom", Color: b.Bottom.Color.ToHex(), Style: int(b.Bottom.Pattern)},
		{Type: "left", Color: b.Left.Color.ToHex(), Style: int(b.Left.Pattern)},
	}
}

func (b *FontStyle) ToExcel() *excelize.Font {
	return &excelize.Font{
		Bold:   b.Bold,
		Italic: b.Italic,
		Size:   b.Size,
		Color:  b.Color.ToHex(),
	}
}

func (f *FillStyle) ToExcel() excelize.Fill {
	switch f.Pattern {
	case FillPatternNone:
		return excelize.Fill{}
	case FillPatternSolid:
		return excelize.Fill{
			Type:    "pattern",
			Pattern: int(f.Pattern),
			Color:   []string{f.Color.ToHex()},
		}
	}
	return excelize.Fill{}
}

type RGB struct {
	R, G, B uint8
}

func (r *RGB) ToHex() string {
	return fmt.Sprintf("%02x%02x%02x", r.R, r.G, r.B)
}

type NumberFormat uint

const (
	NumberFormatGeneral NumberFormat = iota
	NumberFormatDecimal
	NumberFormatFloat
	NumberFormatPercentDecimal
	NumberFormatPercentFloat
	NumberFormatTimeHMS
)

func (n NumberFormat) ToExcel() int {
	switch n {
	case NumberFormatGeneral:
		return 0
	case NumberFormatDecimal:
		return 1
	case NumberFormatFloat:
		return 2
	case NumberFormatPercentDecimal:
		return 9
	case NumberFormatPercentFloat:
		return 10
	case NumberFormatTimeHMS:
		return 21

	}
	return 0
}
