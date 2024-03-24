package deck_test

import (
	"fmt"
	"testing"
	"toggl-card-game/internal/core/deck"

	"github.com/stretchr/testify/assert"
)

func TestNewCard(t *testing.T) {
	type args struct {
		suit deck.Suit
		rank deck.Rank
	}
	tests := []struct {
		name string
		args args
		want string // card code
	}{
		{
			name: "ace of spades",
			args: args{suit: deck.Spades, rank: deck.Ace},
			want: "AS",
		},
		{
			name: "king of hearts",
			args: args{suit: deck.Hearts, rank: deck.King},
			want: "KH",
		},
		{
			name: "two of diamonds",
			args: args{suit: deck.Diamonds, rank: deck.Two},
			want: "2D",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deck.NewCard(tt.args.suit, tt.args.rank)

			if got.Code() != tt.want {
				t.Errorf("NewCard() = %v, got code: %s, want code %s", got, got.Code(), tt.want)
			}
		})
	}
}

func TestBuildNewDeck(t *testing.T) {
	tests := []struct {
		name string
		args *deck.Builder
		want func() (*deck.Deck, error)
	}{
		{
			name: "sequenced full deck test",
			args: deck.NewBuilder(),
			want: func() (*deck.Deck, error) {
				return deck.NewBuilder().Build()
			},
		},
		{
			name: "shuffled full deck test",
			args: deck.NewBuilder().Shuffled(true),
			want: func() (*deck.Deck, error) {
				return deck.NewBuilder().Shuffled(true).Build()
			},
		},
		{
			name: "sequenced not full deck test",
			args: deck.NewBuilder().
				AddCard(deck.NewCard(deck.Spades, deck.Ace)).
				AddCard(deck.NewCard(deck.Hearts, deck.King)).
				AddCard(deck.NewCard(deck.Diamonds, deck.Two)),
			want: func() (*deck.Deck, error) {
				return deck.NewBuilder().
					AddCard(deck.NewCard(deck.Spades, deck.Ace)).
					AddCard(deck.NewCard(deck.Hearts, deck.King)).
					AddCard(deck.NewCard(deck.Diamonds, deck.Two)).
					Build()
			},
		},
		{
			name: "shuffled not full deck test",
			args: deck.NewBuilder().
				Shuffled(true).
				AddCard(deck.NewCard(deck.Spades, deck.Ace)).
				AddCard(deck.NewCard(deck.Hearts, deck.King)).
				AddCard(deck.NewCard(deck.Diamonds, deck.Two)),
			want: func() (*deck.Deck, error) {
				return deck.NewBuilder().
					Shuffled(true).
					AddCard(deck.NewCard(deck.Spades, deck.Ace)).
					AddCard(deck.NewCard(deck.Hearts, deck.King)).
					AddCard(deck.NewCard(deck.Diamonds, deck.Two)).
					Build()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected, err := tt.want()
			if err != nil {
				assert.FailNow(t, err.Error())
				return
			}

			actual, err := tt.args.Build()
			if err != nil {
				assert.FailNow(t, err.Error())
				return
			}

			if actual.Shuffled() {
				assert.NotEqual(t, expected, actual, fmt.Sprintf("expected: %v and actual %v are equal", expected, actual))
				assert.Equal(t, expected.Remaining(), actual.Remaining())
				assert.Equal(t, expected.Shuffled(), actual.Shuffled())
				return
			}

			assert.True(t, expected.Equals(actual), fmt.Sprintf("expected: %v and actual %v are not equal", expected, actual))
		})
	}
}
