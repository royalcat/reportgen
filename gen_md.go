package reportgen

import (
	"fmt"
	"strings"
)

func (report *ReportFile) ToMD() string {
	builder := strings.Builder{}
	builder.WriteString(mdTitle(report.Title))
	for _, sheet := range report.Sheets {
		builder.WriteString(mdTitle2(sheet.Title))
		for _, table := range sheet.Tables {
			builder.WriteString(mdTitle3(table.Title))
			builder.WriteString(mdWriteTable(table))
		}
	}

	return builder.String()
}

func mdTitle(s string) string {
	s = sanitizeLine(s)
	return fmt.Sprintf("# %s\n\n", s)
}

func mdTitle2(s string) string {
	s = sanitizeLine(s)
	return fmt.Sprintf("## %s\n\n", s)
}

func mdTitle3(s string) string {
	s = sanitizeLine(s)
	return fmt.Sprintf("### %s\n\n", s)
}

func mdWriteTable(table Table) string {
	builder := strings.Builder{}
	for _, col := range table.Columns {
		builder.WriteString(fmt.Sprintf("| %s ", sanitizeLine(col)))
	}
	builder.WriteString("|\n")
	for i := 0; i < len(table.Columns); i++ {
		builder.WriteString("| --- ")
	}
	builder.WriteString("|\n")

	table.iterRows(func(row []Cell) bool {
		for _, cell := range row {
			builder.WriteString(fmt.Sprintf("| %s ", sanitizeLine(strings.Join(cell.Strings(), " / "))))
		}
		builder.WriteString("|\n")

		return true
	})
	return builder.String()
}

func sanitizeLine(s string) string {
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\r", "", -1)
	s = strings.Replace(s, "\t", "", -1)
	s = strings.Replace(s, "\v", "", -1)
	s = strings.Replace(s, "\f", "", -1)
	return s
}
