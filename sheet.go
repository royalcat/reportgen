package reportgen

type ReportSheet struct {
	Name  string
	Title string

	TableOffset int
	Tables      []Table
}

func (s *ReportSheet) AddTable(table ...Table) {
	s.Tables = append(s.Tables, table...)
}
