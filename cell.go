package reportgen

import (
	"fmt"
	"time"
)

type Cell struct {
	Data []CellData

	Style *CellStyle
}

func NewCell(data ...CellData) Cell {
	return Cell{Data: data}
}

func (c Cell) Strings() []string {
	out := make([]string, 0, len(c.Data))
	for _, value := range c.Data {
		switch v := value.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			out = append(out, fmt.Sprintf("%d", v))
		case float32, float64:
			out = append(out, fmt.Sprintf("%.2f", v)) // MAYBE всегда оставляет два знака после запятой
		case string:
			out = append(out, v)
		case []byte:
			out = append(out, string(v))
		case time.Duration:
			out = append(out, v.String())
		case time.Time:
			out = append(out, v.String())
		case bool:
			out = append(out, fmt.Sprintf("%t", v))
		case nil:
			out = append(out, "")
		default:
			out = append(out, fmt.Sprint(value))
		}
	}

	return out
}

func (c *Cell) Width() uint {
	return uint(len(c.Data))
}

type CellData interface {
	cellDataInt | cellDataFloat | cellDataString | any
}

type cellDataInt interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

type cellDataFloat interface {
	float32 | float64
}

type cellDataString interface {
	string | []byte
}

type cellDataTime interface {
	time.Time
}

type cellDataMulti interface {
	[]CellData
}
