package deck

import (
	"context"
	"sync"

	"github.com/google/uuid"
)

// Service holds the deck use cases - business logic.
type Service struct {
	repo Repo
	lock sync.Mutex
}

// NewService creates a new deck service.
func NewService(repo Repo) *Service {
	return &Service{
		repo: repo,
		lock: sync.Mutex{},
	}
}

// CreateDeck creates a new deck of cards.
func (s *Service) CreateDeck(ctx context.Context, req CreateRequest) (*CreateResponse, error) {
	deck, err := NewBuilder().Cards(ToCards(req.Cards)).Shuffled(req.Shuffled).Build()
	if err != nil {
		return nil, err
	}

	deck, err = s.repo.Create(ctx, deck)
	if err != nil {
		return nil, NewSvcError(err, ErrCreateDeck)
	}

	return &CreateResponse{
		DeckId:    deck.id.String(),
		Shuffled:  deck.shuffled,
		Remaining: deck.remaining,
	}, nil
}

// OpenDeck opens a deck of cards.
func (s *Service) OpenDeck(ctx context.Context, req OpenRequest) (*OpenResponse, error) {
	id, err := uuid.Parse(req.DeckId)
	if err != nil {
		return nil, err
	}

	deck, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, NewSvcError(err, ErrDeckNotFound)
	}

	cards := make([]CardDto, 0, len(deck.cards))
	for _, c := range deck.cards {
		cards = append(cards, CardDto{
			Value: string(c.rank),
			Suit:  string(c.suit),
			Code:  c.code,
		})
	}

	return &OpenResponse{
		DeckId:    deck.id.String(),
		Shuffled:  deck.shuffled,
		Remaining: deck.remaining,
		Cards:     cards,
	}, nil
}

// DrawCards draws cards from the deck.
func (s *Service) DrawCards(ctx context.Context, req DrawRequest) (*DrawResponse, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	deckID, err := uuid.Parse(req.DeckId)
	if err != nil {
		return nil, err
	}

	deck, err := s.repo.Get(ctx, deckID)
	if err != nil {
		return nil, NewSvcError(err, ErrUpdateDeck)
	}

	// draw cards from the deck
	cards := make([]Card, 0, req.Count)
	for i := 0; i < req.Count; i++ {
		if deck.remaining == 0 {
			break
		}

		cards = append(cards, deck.cards[i])
		deck.remaining--
	}
	deck.cards = deck.cards[req.Count:]

	// update the deck
	s.repo.Update(ctx, deck)

	return &DrawResponse{Cards: ToDtos(cards)}, nil
}
