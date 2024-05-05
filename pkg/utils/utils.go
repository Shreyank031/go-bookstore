package utils

//When we start writing code in book-controller, we need data that is `unmarshelled`.
//We will recieve request from the user and it will be in `json`. We need to unmarshell in order to use
//it inside our controllers. Thats what we are going to do here.

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {

	if body, err := io.ReadAll(r.Body); err == nil {

		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
