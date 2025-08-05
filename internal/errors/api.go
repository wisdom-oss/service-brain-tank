package errors

import (
	"net/http"

	"github.com/wisdom-oss/common-go/v3/types"
)

var ErrMissingParameter = types.ServiceError{
	Type:   "https://www.rfc-editor.org/rfc/rfc9110#section-15.5.1",
	Status: http.StatusBadRequest,
	Title:  "Request Missing Parameter",
	Detail: "The request is missing a required parameter. Check the error field for more information",
}
