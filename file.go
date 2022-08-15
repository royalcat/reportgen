package reportgen

type ReportFile struct {
	Title   string
	Creator string
	Sheets  []ReportSheet

	//stylesMap map[*excelize.Style]int

	//f *excelize.File
}

func NewFile() *ReportFile {
	return &ReportFile{
		//stylesMap: make(map[*excelize.Style]int),
		//f:         excelize.NewFile(),
	}
}

func (r *ReportFile) AddSheet(s ...ReportSheet) {
	r.Sheets = append(r.Sheets, s...)
}

func (r *ReportFile) AddFirstSheet(s ...ReportSheet) {
	r.Sheets = append(s, r.Sheets...)
}
