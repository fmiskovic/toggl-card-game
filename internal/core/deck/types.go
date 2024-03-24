package deck

import (
	"strconv"

	"github.com/google/uuid"
)

// Suit represents a playing card suit.
type Suit string

// Rank represents a playing card value.
type Rank string

// Card represents a playing card.
type Card struct {
	suit Suit
	rank Rank
	code string
}

// NewCard creates a new card with the given suit and rank.
func NewCard(suit Suit, rank Rank) Card {
	// code is first letter of rank and first letter of suit
	// e.g. Ace of Spades is AS
	// e.g. Two of Diamonds is 2D
	// e.g. Ten of Clubs is 10C
	var code string

	//check if rank is integer
	_, err := strconv.Atoi(string(rank))
	if err == nil {
		code = string(rank) + string(suit[0])
	} else {
		code = string(rank[0]) + string(suit[0])
	}

	return Card{suit: suit, rank: rank, code: code}
}

// Code returns the card code.
func (c Card) Code() string {
	return c.code
}

// Deck represents a deck of cards.
type Deck struct {
	id        uuid.UUID
	shuffled  bool
	remaining int
	cards     []Card
}

// Id returns the deck ID.
func (d *Deck) Id() uuid.UUID {
	return d.id
}

// Shuffled returns true if the deck is shuffled.
func (d *Deck) Shuffled() bool {
	return d.shuffled
}

// Remaining returns the number of remaining cards in the deck.
func (d *Deck) Remaining() int {
	return d.remaining
}

// Equals receiver purpose is to compare two decks with out ID.
func (d *Deck) Equals(other *Deck) bool {
	if d.shuffled != other.shuffled {
		return false
	}
	if d.remaining != other.remaining {
		return false
	}
	if len(d.cards) != len(other.cards) {
		return false
	}
	for i, c := range d.cards {
		if c != other.cards[i] {
			return false
		}
	}
	return true
}
