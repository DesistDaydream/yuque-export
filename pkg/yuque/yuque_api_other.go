package yuque

type ExportsData struct {
	Data Export `json:"data"`
}

type Export struct {
	State string `json:"state"`
	URL   string `json:"url"`
}
