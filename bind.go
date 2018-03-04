package bind

import (
	"errors"
	"io"
	"net/http"
)

// Bind bind the request content into a struct you specify
// it find the type itself
func Bind(req *http.Request, uStruct interface{}) error {
	if contentType := req.Header.Get("Content-Type"); contentType != "application/json" {
		return errors.New("Request body has no header json (only supported for now)")
	}

	return JSON(req, uStruct)
}

// JSON deserialize and map the field  into a struct you specify
func JSON(req *http.Request, uStruct interface{}) error {
	var errs Errors

	if req.Body != nil {
		defer req.Body.Close()
		err := DecodeJSON(req.Body, uStruct)
		if err != nil && err != io.EOF {
			errs.Add([]string{}, DeserializationError, err.Error())
			return errs
		}
	} else {
		errs.Add([]string{}, DeserializationError, "Empty request body")
		return errs
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
