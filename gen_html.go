package reportgen

import (
	"fmt"
	"strings"
)

func (report *ReportFile) ToHTML() string {
	builder := strings.Builder{}
	builder.WriteString(htmlTitle(report.Title, 1))
	for _, sheet := range report.Sheets {
		builder.WriteString(htmlTitle(sheet.Title, 2))
		for _, table := range sheet.Tables {
			builder.WriteString(htmlTitle(table.Title, 3))
			builder.WriteString(htmlTable(table))
		}
	}

	return builder.String()
}

func htmlTitle(s string, level uint) string {
	return fmt.Sprintf("<h%d>%s</h%d><br>", level, s, level)
}

func htmlTable(table Table) string {
	builder := strings.Builder{}
	builder.WriteString("<table border=\"0\">")

	builder.WriteString("<tr>")
	for _, colName := range table.Columns {
		builder.WriteString(fmt.Sprintf("<th colspan=%d>%s</th>", table.GetCol(colName).width(), colName))
	}
	builder.WriteString("</tr>")

	table.iterRows(func(row []Cell) bool {
		builder.WriteString("<tr>")
		for _, cell := range row {
			for _, s := range cell.Strings() {
				builder.WriteString(fmt.Sprintf("<td>%s</td>", s))
			}
		}
		builder.WriteString("</tr>")

		return true
	})
	builder.WriteString("</table>")

	return builder.String()
}
