package deck

import (
	"context"

	"github.com/google/uuid"
)

// Repo is the deck port that defines methods that any repository adapter must implement.
type Repo interface {
	Create(ctx context.Context, deck *Deck) (*Deck, error)
	Get(ctx context.Context, id uuid.UUID) (*Deck, error)
	Update(ctx context.Context, deck *Deck) (*Deck, error)
}

// TargetFunc is a generic function type that represents any service function
// and is being called by handlers.
type TargetFunc[In any, Out any] func(context.Context, In) (Out, error)
