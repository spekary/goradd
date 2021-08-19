package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)
import "github.com/elnormous/contenttype"

// DecodeJSONBody will read the json body and return it as a GO value.
//
// If it has a problem, it will panic with an appropriate http error.
//
// Credit goes to Alex Edwards for the ideas behind this code:
// https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body
func DecodeJSONBody(w http.ResponseWriter, r *http.Request, maxBytes int64) interface{} {
	if r.Header.Get("Content-Type") != "" {
		availableMediaTypes := []contenttype.MediaType{
			contenttype.NewMediaType("application/json"),
		}
		_, _, err := contenttype.GetAcceptableMediaType(r, availableMediaTypes)

		if err != nil {
			e := Error{errCode: http.StatusUnsupportedMediaType, message: "Content-Type header is not application/json"}
			panic(e)
		}
	}

	bodyReader := http.MaxBytesReader(w, r.Body, maxBytes)
	dec := json.NewDecoder(bodyReader)
	var j interface{}
	err := dec.Decode(&j)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			e := Error{errCode: http.StatusBadRequest, message: msg}
			panic (e)

		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			e := Error{errCode: http.StatusBadRequest, message: msg}
			panic (e)

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			e := Error{errCode: http.StatusBadRequest, message: msg}
			panic (e)

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			e := Error{errCode: http.StatusBadRequest, message: msg}
			panic (e)

		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			e := Error{errCode: http.StatusBadRequest, message: msg}
			panic (e)

		case err.Error() == "http: request body too large":
			msg := fmt.Sprintf("Request body must not be larger than %d bytes", maxBytes)
			e := Error{errCode: http.StatusRequestEntityTooLarge, message: msg}
			panic (e)

		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		msg := "Request body must only contain a single JSON object"
		e := Error{errCode: http.StatusBadRequest, message: msg}
		panic (e)
	}

	return j
}


