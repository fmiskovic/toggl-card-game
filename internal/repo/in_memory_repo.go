package repo

import (
	"context"
	"fmt"
	"toggl-card-game/internal/core/deck"

	"github.com/google/uuid"
)

// InMemoryRepo implements deck.Repo interface.
// Map is not thread-safe, so we need to use a sync mechanism like mutex to protect it or to use sync.Map.
// Since this is just an example, I will leave it as it is.
// In a real-world application, I would implement CQRS pattern.
type InMemoryRepo struct {
	decks map[uuid.UUID]*deck.Deck
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		decks: make(map[uuid.UUID]*deck.Deck, 52),
	}
}

func (r *InMemoryRepo) Create(ctx context.Context, deck *deck.Deck) (*deck.Deck, error) {
	r.decks[deck.Id()] = deck
	return deck, nil
}

func (r *InMemoryRepo) Get(ctx context.Context, id uuid.UUID) (*deck.Deck, error) {
	deck, ok := r.decks[id]
	if !ok {
		return nil, fmt.Errorf("deck with ID [%s] was not found", id.String())
	}
	return deck, nil
}

func (r *InMemoryRepo) Update(ctx context.Context, deck *deck.Deck) (*deck.Deck, error) {
	r.decks[deck.Id()] = deck
	return deck, nil
}
