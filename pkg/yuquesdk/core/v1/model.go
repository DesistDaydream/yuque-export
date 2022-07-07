package v1

var (
	// BaseAPI address
	BaseAPI = "https://www.yuque.com/api/"
	//EmptyRO empty options
	EmptyRO = new(RequestOption)
)

// book/X/export
type BookExport struct {
	Data BookExportData `json:"data"`
}

type BookExportData struct {
	State string `json:"state"`
	URL   string `json:"url"`
}
