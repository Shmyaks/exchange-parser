// Package schemes is json schemes of Okx
package schemes

type spotInfoJSONScheme struct {
	AskPx  string `json:"askPx"`
	BidPx  string `json:"bidPx"`
	InstIO string `json:"instId"`
	Vol24h string `json:"vol24h"`
}

// SpotJSONScheme json scheme
type SpotJSONScheme struct {
	Data []spotInfoJSONScheme `json:"data"`
}
