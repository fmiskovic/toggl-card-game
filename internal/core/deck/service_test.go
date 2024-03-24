package deck_test

import (
	"context"
	"errors"
	"testing"
	"toggl-card-game/internal/core/deck"
	mocks "toggl-card-game/mocks/internal_/core/deck"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_CreateDeck(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name    string
		given   deck.CreateRequest
		when    func() (*deck.Deck, error)
		want    *deck.CreateResponse
		wantErr bool
	}{
		{
			name:  "create full sequenced deck test",
			given: deck.CreateRequest{},
			when: func() (*deck.Deck, error) {
				return deck.NewBuilder().Build()
			},
			want: &deck.CreateResponse{
				DeckId:    uuid.NewString(),
				Shuffled:  false,
				Remaining: 52,
			},
			wantErr: false,
		},
		{
			name: "create partial sequenced deck test",
			given: deck.CreateRequest{
				Cards: []string{"AS", "2S", "3S"},
			},
			when: func() (*deck.Deck, error) {
				return deck.NewBuilder().
					AddCard(deck.CardsMap["AS"]).
					AddCard(deck.CardsMap["2S"]).
					AddCard(deck.CardsMap["3S"]).
					Build()
			},
			want: &deck.CreateResponse{
				DeckId:    uuid.NewString(),
				Shuffled:  false,
				Remaining: 3,
			},
			wantErr: false,
		},
		{
			name:  "create full shuffled deck test",
			given: deck.CreateRequest{},
			when: func() (*deck.Deck, error) {
				return deck.NewBuilder().Shuffled(true).Build()
			},
			want: &deck.CreateResponse{
				DeckId:    uuid.NewString(),
				Shuffled:  true,
				Remaining: 52,
			},
			wantErr: false,
		},
		{
			name: "create partial shuffled deck test",
			given: deck.CreateRequest{
				Cards: []string{"AS", "2S", "3S"},
			},
			when: func() (*deck.Deck, error) {
				return deck.NewBuilder().
					Shuffled(true).
					AddCard(deck.CardsMap["AS"]).
					AddCard(deck.CardsMap["2S"]).
					AddCard(deck.CardsMap["3S"]).
					Build()
			},
			want: &deck.CreateResponse{
				DeckId:    uuid.NewString(),
				Shuffled:  true,
				Remaining: 3,
			},
			wantErr: false,
		},
		{
			name: "repo returns an error test",
			given: deck.CreateRequest{
				Cards: []string{"AS", "2S", "3S"},
			},
			when: func() (*deck.Deck, error) {
				return nil, errors.New("repo error")
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock repo call
			d, err := tt.when()
			repoMock := mocks.NewRepo(t)
			repoMock.EXPECT().Create(ctx, mock.Anything).Return(d, err)

			// service under test
			svc := deck.NewService(repoMock)

			actual, err := svc.CreateDeck(ctx, tt.given)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}

			if err != nil {
				assert.FailNow(t, err.Error())
				return
			}

			assert.NotNil(t, actual.DeckId)
			assert.Equal(t, tt.want.Shuffled, actual.Shuffled)
			assert.Equal(t, tt.want.Remaining, actual.Remaining)
		})
	}
}

func TestService_OpenDeck(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name    string
		args    deck.OpenRequest
		when    func() (*deck.Deck, error)
		want    *deck.OpenResponse
		wantErr bool
	}{
		{
			name: "open full sequenced deck test",
			args: deck.OpenRequest{DeckId: uuid.NewString()},
			when: func() (*deck.Deck, error) {
				return deck.NewBuilder().Build()
			},
			want: &deck.OpenResponse{
				DeckId:    uuid.NewString(),
				Shuffled:  false,
				Remaining: 52,
				Cards:     dtos,
			},
			wantErr: false,
		},
		{
			name: "open partial sequenced deck test",
			args: deck.OpenRequest{DeckId: uuid.NewString()},
			when: func() (*deck.Deck, error) {
				return deck.NewBuilder().
					AddCard(deck.CardsMap["AS"]).
					AddCard(deck.CardsMap["2S"]).
					AddCard(deck.CardsMap["10S"]).
					Build()
			},
			want: &deck.OpenResponse{
				DeckId:    uuid.NewString(),
				Shuffled:  false,
				Remaining: 3,
				Cards: []deck.CardDto{
					deck.CardsMap["AS"].ToDto(),
					deck.CardsMap["2S"].ToDto(),
					deck.CardsMap["10S"].ToDto(),
				},
			},
			wantErr: false,
		},
		{
			name: "open full shuffled deck test",
			args: deck.OpenRequest{DeckId: uuid.NewString()},
			when: func() (*deck.Deck, error) {
				return deck.NewBuilder().Shuffled(true).Build()
			},
			want: &deck.OpenResponse{
				DeckId:    uuid.NewString(),
				Shuffled:  true,
				Remaining: 52,
				Cards:     dtos,
			},
			wantErr: false,
		},
		{
			name: "open partial shuffled deck test",
			args: deck.OpenRequest{DeckId: uuid.NewString()},
			when: func() (*deck.Deck, error) {
				return deck.NewBuilder().
					Shuffled(true).
					AddCard(deck.CardsMap["AS"]).
					AddCard(deck.CardsMap["2S"]).
					AddCard(deck.CardsMap["3S"]).
					Build()
			},
			want: &deck.OpenResponse{
				DeckId:    uuid.NewString(),
				Shuffled:  true,
				Remaining: 3,
				Cards: []deck.CardDto{
					deck.CardsMap["AS"].ToDto(),
					deck.CardsMap["2S"].ToDto(),
					deck.CardsMap["3S"].ToDto(),
				},
			},
			wantErr: false,
		},
		{
			name: "open deck invalid id test",
			args: deck.OpenRequest{DeckId: "invalid-id"},
			when: func() (*deck.Deck, error) {
				return deck.NewBuilder().Build()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "open deck repo returns an error test",
			args: deck.OpenRequest{DeckId: uuid.NewString()},
			when: func() (*deck.Deck, error) {
				return nil, errors.New("repo error")
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock repo call
			d, err := tt.when()
			repoMock := mocks.NewRepo(t)
			repoMock.On("Get", ctx, mock.Anything).Return(d, err).Maybe()

			// service under test
			svc := deck.NewService(repoMock)

			actual, err := svc.OpenDeck(ctx, tt.args)

			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}

			if err != nil {
				assert.FailNow(t, err.Error())
				return
			}

			assert.NotNil(t, actual.DeckId)
			assert.Equal(t, tt.want.Shuffled, actual.Shuffled)
			assert.Equal(t, tt.want.Remaining, actual.Remaining)
			assert.True(t, len(actual.Cards) == len(tt.want.Cards))

			if !tt.want.Shuffled {
				assert.Equal(t, actual.Cards, tt.want.Cards)
			}
		})
	}
}

func TestService_DrawCards(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name    string
		args    deck.DrawRequest
		when    func() (*deck.Deck, error)
		want    *deck.DrawResponse
		wantErr bool
	}{
		{
			name: "draw 3 cards from full sequenced deck test",
			args: deck.DrawRequest{DeckId: uuid.NewString(), Count: 3},
			when: func() (*deck.Deck, error) {
				return deck.NewBuilder().Build()
			},
			want: &deck.DrawResponse{
				Cards: dtos[:3],
			},
			wantErr: false,
		},
		{
			name: "draw 2 cards from partial sequenced deck test",
			args: deck.DrawRequest{DeckId: uuid.NewString(), Count: 2},
			when: func() (*deck.Deck, error) {
				return deck.NewBuilder().
					AddCard(deck.CardsMap["AS"]).
					AddCard(deck.CardsMap["2S"]).
					AddCard(deck.CardsMap["3S"]).
					Build()
			},
			want: &deck.DrawResponse{
				Cards: []deck.CardDto{
					deck.CardsMap["AS"].ToDto(),
					deck.CardsMap["2S"].ToDto(),
				},
			},
			wantErr: false,
		},
		{
			name: "open deck invalid id test",
			args: deck.DrawRequest{DeckId: "invalid-id"},
			when: func() (*deck.Deck, error) {
				return deck.NewBuilder().Build()
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "open deck repo returns an error test",
			args: deck.DrawRequest{DeckId: uuid.NewString(), Count: 3},
			when: func() (*deck.Deck, error) {
				return nil, errors.New("repo error")
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock repo call
			d, err := tt.when()
			repoMock := mocks.NewRepo(t)
			repoMock.On("Get", ctx, mock.Anything).Return(d, err).Maybe()
			repoMock.On("Update", ctx, mock.Anything).Return(d, err).Maybe()

			// service under test
			svc := deck.NewService(repoMock)

			actual, err := svc.DrawCards(ctx, tt.args)

			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}

			assert.Equal(t, tt.want, actual)
		})
	}
}

// full deck of sequenced cards
var dtos = []deck.CardDto{
	{Value: "ACE", Suit: "SPADES", Code: "AS"},
	{Value: "2", Suit: "SPADES", Code: "2S"},
	{Value: "3", Suit: "SPADES", Code: "3S"},
	{Value: "4", Suit: "SPADES", Code: "4S"},
	{Value: "5", Suit: "SPADES", Code: "5S"},
	{Value: "6", Suit: "SPADES", Code: "6S"},
	{Value: "7", Suit: "SPADES", Code: "7S"},
	{Value: "8", Suit: "SPADES", Code: "8S"},
	{Value: "9", Suit: "SPADES", Code: "9S"},
	{Value: "10", Suit: "SPADES", Code: "10S"},
	{Value: "JACK", Suit: "SPADES", Code: "JS"},
	{Value: "QUEEN", Suit: "SPADES", Code: "QS"},
	{Value: "KING", Suit: "SPADES", Code: "KS"},
	{Value: "ACE", Suit: "DIAMONDS", Code: "AD"},
	{Value: "2", Suit: "DIAMONDS", Code: "2D"},
	{Value: "3", Suit: "DIAMONDS", Code: "3D"},
	{Value: "4", Suit: "DIAMONDS", Code: "4D"},
	{Value: "5", Suit: "DIAMONDS", Code: "5D"},
	{Value: "6", Suit: "DIAMONDS", Code: "6D"},
	{Value: "7", Suit: "DIAMONDS", Code: "7D"},
	{Value: "8", Suit: "DIAMONDS", Code: "8D"},
	{Value: "9", Suit: "DIAMONDS", Code: "9D"},
	{Value: "10", Suit: "DIAMONDS", Code: "10D"},
	{Value: "JACK", Suit: "DIAMONDS", Code: "JD"},
	{Value: "QUEEN", Suit: "DIAMONDS", Code: "QD"},
	{Value: "KING", Suit: "DIAMONDS", Code: "KD"},
	{Value: "ACE", Suit: "CLUBS", Code: "AC"},
	{Value: "2", Suit: "CLUBS", Code: "2C"},
	{Value: "3", Suit: "CLUBS", Code: "3C"},
	{Value: "4", Suit: "CLUBS", Code: "4C"},
	{Value: "5", Suit: "CLUBS", Code: "5C"},
	{Value: "6", Suit: "CLUBS", Code: "6C"},
	{Value: "7", Suit: "CLUBS", Code: "7C"},
	{Value: "8", Suit: "CLUBS", Code: "8C"},
	{Value: "9", Suit: "CLUBS", Code: "9C"},
	{Value: "10", Suit: "CLUBS", Code: "10C"},
	{Value: "JACK", Suit: "CLUBS", Code: "JC"},
	{Value: "QUEEN", Suit: "CLUBS", Code: "QC"},
	{Value: "KING", Suit: "CLUBS", Code: "KC"},
	{Value: "ACE", Suit: "HEARTS", Code: "AH"},
	{Value: "2", Suit: "HEARTS", Code: "2H"},
	{Value: "3", Suit: "HEARTS", Code: "3H"},
	{Value: "4", Suit: "HEARTS", Code: "4H"},
	{Value: "5", Suit: "HEARTS", Code: "5H"},
	{Value: "6", Suit: "HEARTS", Code: "6H"},
	{Value: "7", Suit: "HEARTS", Code: "7H"},
	{Value: "8", Suit: "HEARTS", Code: "8H"},
	{Value: "9", Suit: "HEARTS", Code: "9H"},
	{Value: "10", Suit: "HEARTS", Code: "10H"},
	{Value: "JACK", Suit: "HEARTS", Code: "JH"},
	{Value: "QUEEN", Suit: "HEARTS", Code: "QH"},
	{Value: "KING", Suit: "HEARTS", Code: "KH"},
}
