package middleware

import (
	customerror "app/internal/custom_error"
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/lib/pq"
)

type ErrorResponse struct {
	Messages []ErrorResponseItem `json:"messages"`
	Type     string              `json:"type"`
}

type ErrorResponseItem struct {
	Name   string `json:"name"`
	Reason string `json:"reason"`
}

func ErrorFormatter(w http.ResponseWriter, r *http.Request, err error) {
	// Extract the corresponding encoder from the Accept of Request Header
	code, body := parseError(err)

	w.Header().Set("Content-Type", "application/json")
	// Set HTTP Status Code
	w.WriteHeader(code)
	_, _ = w.Write(body)
}

func parseValidationError(err *customerror.ValidationError) (int, ErrorResponse) {
	return http.StatusBadRequest, ErrorResponse{
		Type:     "ValidationError",
		Messages: parseItemValidationError(err),
	}
}

func parsePQError(err *pq.Error) (int, ErrorResponse) {
	if err.Code == "23505" {
		return http.StatusConflict, ErrorResponse{
			Type: "DuplicateResource",
			Messages: []ErrorResponseItem{
				{
					Name:   "resource",
					Reason: "duplicate violation on unique constraint",
				},
			},
		}
	}

	return parseDefaultError(err)
}

func parseNoRowsError(err error) (int, ErrorResponse) {
	return http.StatusNotFound, ErrorResponse{
		Type: "RecordNotFound",
		Messages: []ErrorResponseItem{
			{
				Name:   "data",
				Reason: err.Error(),
			},
		},
	}
}

func parseDefaultError(err error) (int, ErrorResponse) {
	return http.StatusInternalServerError, ErrorResponse{
		Type: "InternalServerError",
		Messages: []ErrorResponseItem{
			{
				Name:   "UnexpectedError",
				Reason: err.Error(),
			},
		},
	}
}

func parseError(err error) (httpCode int, body []byte) {
	var errResponse ErrorResponse
	switch parsedError := err.(type) {
	case *customerror.ValidationError:
		httpCode, errResponse = parseValidationError(parsedError)
	case *pq.Error:
		httpCode, errResponse = parsePQError(parsedError)
	default:
		if errors.Is(parsedError, sql.ErrNoRows) {
			httpCode, errResponse = parseNoRowsError(parsedError)
		} else {
			httpCode, errResponse = parseDefaultError(parsedError)
		}
	}

	body, _ = json.Marshal(errResponse)

	return httpCode, body
}

func parseItemValidationError(err *customerror.ValidationError) []ErrorResponseItem {
	var errorItems []ErrorResponseItem

	for _, message := range strings.Split(err.Error(), ";") {
		messages := strings.Split(message, ": ")

		key := messages[0]
		for _, detailError := range strings.Split(messages[1], ",") {
			errorItems = append(errorItems, ErrorResponseItem{
				Name:   key,
				Reason: detailError,
			})
		}
	}
	return errorItems
}
