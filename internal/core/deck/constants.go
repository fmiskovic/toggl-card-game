package deck

const (
	Spades   Suit = "SPADES"   // ♠
	Diamonds Suit = "DIAMONDS" // ♦
	Clubs    Suit = "CLUBS"    // ♣
	Hearts   Suit = "HEARTS"   // ♥
)

const (
	Ace   Rank = "ACE"
	Two   Rank = "2"
	Three Rank = "3"
	Four  Rank = "4"
	Five  Rank = "5"
	Six   Rank = "6"
	Seven Rank = "7"
	Eight Rank = "8"
	Nine  Rank = "9"
	Ten   Rank = "10"
	Jack  Rank = "JACK"
	Queen Rank = "QUEEN"
	King  Rank = "KING"
)

// CardsMap maps card codes to Card objects, useful for quick lookup.
var CardsMap = map[string]Card{
	"AS":  NewCard(Spades, Ace),
	"2S":  NewCard(Spades, Two),
	"3S":  NewCard(Spades, Three),
	"4S":  NewCard(Spades, Four),
	"5S":  NewCard(Spades, Five),
	"6S":  NewCard(Spades, Six),
	"7S":  NewCard(Spades, Seven),
	"8S":  NewCard(Spades, Eight),
	"9S":  NewCard(Spades, Nine),
	"10S": NewCard(Spades, Ten),
	"JS":  NewCard(Spades, Jack),
	"QS":  NewCard(Spades, Queen),
	"KS":  NewCard(Spades, King),
	"AD":  NewCard(Diamonds, Ace),
	"2D":  NewCard(Diamonds, Two),
	"3D":  NewCard(Diamonds, Three),
	"4D":  NewCard(Diamonds, Four),
	"5D":  NewCard(Diamonds, Five),
	"6D":  NewCard(Diamonds, Six),
	"7D":  NewCard(Diamonds, Seven),
	"8D":  NewCard(Diamonds, Eight),
	"9D":  NewCard(Diamonds, Nine),
	"10D": NewCard(Diamonds, Ten),
	"JD":  NewCard(Diamonds, Jack),
	"QD":  NewCard(Diamonds, Queen),
	"KD":  NewCard(Diamonds, King),
	"AC":  NewCard(Clubs, Ace),
	"2C":  NewCard(Clubs, Two),
	"3C":  NewCard(Clubs, Three),
	"4C":  NewCard(Clubs, Four),
	"5C":  NewCard(Clubs, Five),
	"6C":  NewCard(Clubs, Six),
	"7C":  NewCard(Clubs, Seven),
	"8C":  NewCard(Clubs, Eight),
	"9C":  NewCard(Clubs, Nine),
	"10C": NewCard(Clubs, Ten),
	"JC":  NewCard(Clubs, Jack),
	"QC":  NewCard(Clubs, Queen),
	"KC":  NewCard(Clubs, King),
	"AH":  NewCard(Hearts, Ace),
	"2H":  NewCard(Hearts, Two),
	"3H":  NewCard(Hearts, Three),
	"4H":  NewCard(Hearts, Four),
	"5H":  NewCard(Hearts, Five),
	"6H":  NewCard(Hearts, Six),
	"7H":  NewCard(Hearts, Seven),
	"8H":  NewCard(Hearts, Eight),
	"9H":  NewCard(Hearts, Nine),
	"10H": NewCard(Hearts, Ten),
	"JH":  NewCard(Hearts, Jack),
	"QH":  NewCard(Hearts, Queen),
	"KH":  NewCard(Hearts, King),
}
