package models

// book/X/export
type BookExport struct {
	Data BookExportData `json:"data"`
}

type BookExportData struct {
	State string `json:"state"`
	URL   string `json:"url"`
}
