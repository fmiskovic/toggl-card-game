package deck

import (
	"math/rand"

	"github.com/google/uuid"
)

type Builder struct {
	id       uuid.UUID
	shuffled bool
	cards    []Card
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) Id(id uuid.UUID) *Builder {
	b.id = id
	return b
}

func (b *Builder) Shuffled(shuffled bool) *Builder {
	b.shuffled = shuffled
	return b
}

func (b *Builder) Cards(cards []Card) *Builder {
	b.cards = cards
	return b
}

func (b *Builder) AddCard(card Card) *Builder {
	b.cards = append(b.cards, card)
	return b
}

func (b *Builder) Build() (*Deck, error) {
	deck := &Deck{}

	// if b has properties and set them to d OR set default values
	if b.id != uuid.Nil {
		deck.id = b.id
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}
		deck.id = id
	}

	if len(b.cards) == 0 {
		deck.cards = initAllCards()
	} else {
		deck.cards = b.cards
	}

	deck.remaining = len(deck.cards)

	if b.shuffled {
		shuffled := deck.cards
		rand.Shuffle(len(shuffled), func(i, j int) {
			shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
		})
		deck.cards = shuffled
		deck.shuffled = true
	}

	return deck, nil
}

func initAllCards() []Card {
	cards := make([]Card, 0, 52)
	for _, suit := range []Suit{Spades, Diamonds, Clubs, Hearts} {
		for _, rank := range []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King} {
			cards = append(cards, NewCard(suit, rank))
		}
	}
	return cards
}
