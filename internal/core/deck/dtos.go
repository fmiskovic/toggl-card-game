package deck

// CardDto represents a data transfer object for a card.
type CardDto struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

// ToDto converts a Card entity to a CardDto.
func (c Card) ToDto() CardDto {
	return CardDto{
		Value: string(c.rank),
		Suit:  string(c.suit),
		Code:  c.code,
	}
}

// ToDtos converts a slice of Card entities to a slice of CardDto.
func ToDtos(cards []Card) []CardDto {
	dtos := make([]CardDto, 0, len(cards))
	for _, c := range cards {
		dtos = append(dtos, c.ToDto())
	}
	return dtos
}

// ToCards converts a slice of card codes to a slice of Card entities.
func ToCards(codes []string) []Card {
	cards := make([]Card, 0, len(codes))
	for _, code := range codes {
		cards = append(cards, CardsMap[code])
	}
	return cards
}

// CreateRequest represents a request to create a deck.
type CreateRequest struct {
	Shuffled bool
	Cards    []string
}

// CreateResponse represents a response for creating a deck.
type CreateResponse struct {
	DeckId    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

// OpenRequest represents a request to open a deck.
type OpenRequest struct {
	DeckId string
}

// OpenResponse represents a response for opening a deck.
type OpenResponse struct {
	DeckId    string    `json:"deck_id"`
	Shuffled  bool      `json:"shuffled"`
	Remaining int       `json:"remaining"`
	Cards     []CardDto `json:"cards"`
}

// DrawRequest represents a request to draw cards from a deck.
type DrawRequest struct {
	DeckId string `json:"deck_id"`
	Count  int    `json:"count"`
}

// DrawResponse represents a response for drawing cards from a deck.
type DrawResponse struct {
	Cards []CardDto `json:"cards"`
}
