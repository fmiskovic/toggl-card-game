package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"toggl-card-game/internal/core/deck"

	"github.com/google/uuid"
)

// MyHandlerFunc is a custom handler function type that returns an error.
type MyHandlerFunc func(http.ResponseWriter, *http.Request) error

// MakeHandler decorates a custom handler function.
func MakeHandler(fn MyHandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			switch e := err.(type) {
			case ApiError:
				writeJson(w, e.Code, e)
			default:
				writeJson(w, http.StatusInternalServerError, map[string]string{"errorMessage": err.Error(), "httpCode": "500"})
			}
		}
	})
}

func writeJson(w http.ResponseWriter, code int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	// Convert out to JSON bytes
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Write JSON response body
	_, err = w.Write(b)
	if err != nil {
		return err
	}

	return nil
}

// RequestFunc is a custom request parser function type.
// It takes an http.Request and returns a input value or an error.
type RequestParserFunc[In any] func(r *http.Request) (In, error)

// Handle is a generic handler function that takes a target function and returns a custom handler function.
func Handle[In any, Out any](reqPar RequestParserFunc[In], svcFunc deck.TargetFunc[In, Out]) MyHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		// Parse request
		in, err := reqPar(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return err
		}

		// Call service function
		out, err := svcFunc(r.Context(), in)
		if err != nil {
			switch e := err.(type) {
			case deck.SvcError:
				return ApiError{e.Error(), http.StatusBadRequest}
			default:
				return err
			}
		}

		return writeJson(w, http.StatusOK, out)
	}
}

func ParseCreateRequest(r *http.Request) (deck.CreateRequest, error) {
	var req deck.CreateRequest

	// Parse query parameters
	q := r.URL.Query()
	if q.Has("cards") {
		// TODO: validate card codes
		// codes := strings.Split(q.Get("cards"), ",")
		// for _, code := range codes {
		// 	card := deck.CardsMap[code]
		// 	if card == nil {
		// 		return req, ApiError{http.StatusBadRequest, "invalid card code"}
		// 	}
		// }
		req.Cards = strings.Split(q.Get("cards"), ",")
	}

	if q.Has("shuffled") {
		shuffled, err := strconv.ParseBool(q.Get("shuffled"))
		if err != nil {
			slog.Warn("unable to parse shuffled query parameter, using default false value", "error", err)
			shuffled = false
		}
		req.Shuffled = shuffled
	}

	return req, nil
}

func ParseOpenRequest(r *http.Request) (deck.OpenRequest, error) {
	id := r.PathValue("UUID")
	_, err := uuid.Parse(id)
	if err != nil {
		return deck.OpenRequest{}, err
	}

	return deck.OpenRequest{DeckId: id}, nil
}

func ParseDrawRequest(r *http.Request) (deck.DrawRequest, error) {
	req := new(deck.DrawRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return *req, err
	}

	return *req, nil
}
