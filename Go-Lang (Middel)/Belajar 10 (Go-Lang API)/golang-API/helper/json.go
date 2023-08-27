package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body) //baca json dulu
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(w http.ResponseWriter, response interface{}) { //writer
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	PanicIfError(err)
}
