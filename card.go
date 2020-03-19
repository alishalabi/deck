//go:generate stringer -type=Suit,Rank

package deck

import "fmt"

// Suit corresponds to clubs, spades, hearts, diamonds, and jokers
type Suit uint8
const (
  Club Suit = iota // Iota increments each const
  Spade
  Heart
  Diamond
  Joker
)

// Create array of suits (small number of items)
var suits = [...]Suit{Club, Spade, Heart, Diamond}

// Ranks correspond to card's value (2, 7, King, etc)
type Rank uint8
const (
  _ Rank = iota // Offsetting values by 1 (to correspond to card value)
  Ace
  Two
  Three
  Four
  Five
  Six
  Seven
  Eight
  Nine
  Ten
  Jack
  Queen
  King
)

const (
  minRank = Ace
  maxRank = King
)


type Card struct {
  Suit
  Rank
}

// Return a string corresponding to card (Ex: King of Spades)
func (c Card) String() string {
  if c.Suit == Joker {
    return c.Suit.String()
  }
  return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func New() []Card {
  var cards []Card
  for _, suit := range suits {
    for rank := minRank; rank <= maxRank; rank++ {
      cards = append(cards, Card{ Suit: suit, Rank: rank})
    }
  }
  return cards
}
