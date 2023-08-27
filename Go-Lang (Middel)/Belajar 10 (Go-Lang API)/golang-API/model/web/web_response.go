package web

type WebRespons struct {
	//untuk data yang ada di json
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
