package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"toggl-card-game/internal/core/deck"
	"toggl-card-game/internal/handlers"
	"toggl-card-game/internal/repo"

	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestHandleCreateDeck(t *testing.T) {
	svc := deck.NewService(repo.NewInMemoryRepo())
	handler := handlers.MakeHandler(handlers.Handle(handlers.ParseCreateRequest, svc.CreateDeck))

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/deck", handler)
	server := httptest.NewServer(mux)

	defer server.Close()

	tests := []struct {
		name     string
		route    string
		wantCode int
		verify   func(t *testing.T, res *http.Response)
	}{
		{
			name:     "create full sequenced deck test",
			route:    "/api/deck",
			wantCode: http.StatusOK,
			verify: func(t *testing.T, res *http.Response) {
				resBody := res.Body
				defer resBody.Close()

				createRes := new(deck.CreateResponse)
				err := json.NewDecoder(resBody).Decode(createRes)
				assert.Nil(t, err)
				assert.Equal(t, 52, createRes.Remaining)
				assert.Equal(t, false, createRes.Shuffled)
			},
		},
		{
			name:     "create full shuffled deck test",
			route:    "/api/deck?shuffled=true",
			wantCode: http.StatusOK,
			verify: func(t *testing.T, res *http.Response) {
				resBody := res.Body
				defer resBody.Close()

				createRes := new(deck.CreateResponse)
				err := json.NewDecoder(resBody).Decode(createRes)
				assert.Nil(t, err)
				assert.Equal(t, 52, createRes.Remaining)
				assert.Equal(t, true, createRes.Shuffled)
			},
		},
		{
			name:     "create partial sequenced deck test",
			route:    "/api/deck?cards=AS,10S,JS,QS,KS",
			wantCode: http.StatusOK,
			verify: func(t *testing.T, res *http.Response) {
				resBody := res.Body
				defer resBody.Close()

				createRes := new(deck.CreateResponse)
				err := json.NewDecoder(resBody).Decode(createRes)
				assert.Nil(t, err)
				assert.Equal(t, 5, createRes.Remaining)
				assert.Equal(t, false, createRes.Shuffled)
			},
		},
		{
			name:     "create partial shuffled deck test",
			route:    "/api/deck?cards=AS,10S,JS,QS,KS&shuffled=true",
			wantCode: http.StatusOK,
			verify: func(t *testing.T, res *http.Response) {
				resBody := res.Body
				defer resBody.Close()

				createRes := new(deck.CreateResponse)
				err := json.NewDecoder(resBody).Decode(createRes)
				assert.Nil(t, err)
				assert.Equal(t, 5, createRes.Remaining)
				assert.Equal(t, true, createRes.Shuffled)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.Post(server.URL+tt.route, "application/json", nil)

			if err != nil {
				t.Fatalf("error making request to server. Err: %v", err)
			}

			if resp.StatusCode != tt.wantCode {
				t.Fatalf("expected status code %d, got %d", tt.wantCode, resp.StatusCode)
			}

			tt.verify(t, resp)
		})
	}
}

func TestHandleOpenDeck(t *testing.T) {
	memoryRepo := repo.NewInMemoryRepo()

	// seed with decks
	deck1, _ := deck.NewBuilder().Build()
	deck2, _ := deck.NewBuilder().
		AddCard(deck.CardsMap["AS"]).
		AddCard(deck.CardsMap["AD"]).
		AddCard(deck.CardsMap["AH"]).
		Build()

	memoryRepo.Create(context.Background(), deck1)
	memoryRepo.Create(context.Background(), deck2)

	svc := deck.NewService(memoryRepo)
	handler := handlers.MakeHandler(handlers.Handle(handlers.ParseOpenRequest, svc.OpenDeck))

	mux := http.NewServeMux()
	mux.HandleFunc("/api/deck/{UUID}", handler)
	server := httptest.NewServer(mux)

	defer server.Close()

	tests := []struct {
		name     string
		route    string
		wantCode int
		verify   func(t *testing.T, res *http.Response)
	}{
		{
			name:     "open full sequenced deck test",
			route:    fmt.Sprintf("/api/deck/%s", deck1.Id().String()),
			wantCode: http.StatusOK,
			verify: func(t *testing.T, res *http.Response) {
				resBody := res.Body
				defer resBody.Close()

				openRes := new(deck.OpenResponse)
				err := json.NewDecoder(resBody).Decode(openRes)
				assert.Nil(t, err)
				assert.Equal(t, 52, openRes.Remaining)
			},
		},
		{
			name:     "open partial sequenced deck test",
			route:    fmt.Sprintf("/api/deck/%s", deck2.Id().String()),
			wantCode: http.StatusOK,
			verify: func(t *testing.T, res *http.Response) {
				resBody := res.Body
				defer resBody.Close()

				openRes := new(deck.OpenResponse)
				err := json.NewDecoder(resBody).Decode(openRes)
				assert.Nil(t, err)
				assert.Equal(t, 3, openRes.Remaining)
			},
		},
		{
			name:     "open non-existent deck test",
			route:    fmt.Sprintf("/api/deck/%s", uuid.New().String()),
			wantCode: http.StatusBadRequest,
			verify: func(t *testing.T, res *http.Response) {
				resBody := res.Body
				defer resBody.Close()

				apiErr := new(handlers.ApiError)
				err := json.NewDecoder(resBody).Decode(apiErr)
				assert.Nil(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.Get(server.URL + tt.route)
			if err != nil {
				t.Fatalf("error making request to server. Err: %v", err)
			}

			if resp.StatusCode != tt.wantCode {
				t.Fatalf("expected status code %d, got %d", tt.wantCode, resp.StatusCode)
			}

			tt.verify(t, resp)
		})
	}
}

func TestHandleDrawCards(t *testing.T) {
	memoryRepo := repo.NewInMemoryRepo()
	// seed decks
	deck1, _ := deck.NewBuilder().Build()
	deck2, _ := deck.NewBuilder().
		AddCard(deck.CardsMap["AS"]).
		AddCard(deck.CardsMap["AD"]).
		AddCard(deck.CardsMap["AH"]).
		Build()

	memoryRepo.Create(context.Background(), deck1)
	memoryRepo.Create(context.Background(), deck2)

	svc := deck.NewService(memoryRepo)
	handler := handlers.MakeHandler(handlers.Handle(handlers.ParseDrawRequest, svc.DrawCards))

	mux := http.NewServeMux()
	mux.HandleFunc("PUT /api/deck", handler)
	server := httptest.NewServer(mux)

	defer server.Close()

	tests := []struct {
		name     string
		route    string
		args     deck.DrawRequest
		wantCode int
		verify   func(t *testing.T, res *http.Response)
	}{
		{
			name:  "draw from full sequenced deck test",
			route: "/api/deck",
			args: deck.DrawRequest{
				DeckId: deck1.Id().String(),
				Count:  3,
			},
			wantCode: http.StatusOK,
			verify: func(t *testing.T, res *http.Response) {
				resBody := res.Body
				defer resBody.Close()

				drawRes := new(deck.DrawResponse)
				err := json.NewDecoder(resBody).Decode(drawRes)
				assert.Nil(t, err)
				assert.Equal(t, 3, len(drawRes.Cards))
				assert.Equal(t, []deck.CardDto{
					{Value: "ACE", Suit: "SPADES", Code: "AS"},
					{Value: "2", Suit: "SPADES", Code: "2S"},
					{Value: "3", Suit: "SPADES", Code: "3S"},
				}, drawRes.Cards)
			},
		},
		{
			name:  "draw from partial sequenced deck test",
			route: "/api/deck",
			args: deck.DrawRequest{
				DeckId: deck2.Id().String(),
				Count:  2,
			},
			wantCode: http.StatusOK,
			verify: func(t *testing.T, res *http.Response) {
				resBody := res.Body
				defer resBody.Close()

				drawRes := new(deck.DrawResponse)
				err := json.NewDecoder(resBody).Decode(drawRes)
				assert.Nil(t, err)
				assert.Equal(t, []deck.CardDto{
					{Value: "ACE", Suit: "SPADES", Code: "AS"},
					{Value: "ACE", Suit: "DIAMONDS", Code: "AD"},
				}, drawRes.Cards)
			},
		},
		{
			name:  "draw from non-existent deck test",
			route: "/api/deck",
			args: deck.DrawRequest{
				DeckId: uuid.NewString(),
				Count:  3,
			},
			wantCode: http.StatusBadRequest,
			verify: func(t *testing.T, res *http.Response) {
				resBody := res.Body
				defer resBody.Close()

				apiErr := new(handlers.ApiError)
				err := json.NewDecoder(resBody).Decode(apiErr)
				assert.Nil(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.args)
			if err != nil {
				t.Fatalf("error marshalling request body. Err: %v", err)
			}
			req, _ := http.NewRequest("PUT", server.URL+tt.route, bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("error making request to server. Err: %v", err)
			}

			if resp.StatusCode != tt.wantCode {
				t.Fatalf("expected status code %d, got %d", tt.wantCode, resp.StatusCode)
			}

			tt.verify(t, resp)
		})
	}
}
