package entities

type Options struct {
	Method  string `json:"method"`
	Data    string `json:"data"`
	URL     string `json:"url"`
	Header  string `json:"header,omitempty"`
	Persist string `json:"persist,omitempty"`
}
