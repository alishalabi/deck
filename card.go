//go:generate stringer -type=Suit,Rank

package deck

import (
  "fmt"
  "sort"
  "math/rand"
  "time"
)

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

// String return a string corresponding to card (Ex: King of Spades)
func (c Card) String() string {
  if c.Suit == Joker {
    return c.Suit.String()
  }
  return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

// New creates a new deck, allow options
func New(options ...func([]Card) []Card) []Card {
  var cards []Card
  for _, suit := range suits {
    for rank := minRank; rank <= maxRank; rank++ {
      cards = append(cards, Card{ Suit: suit, Rank: rank})
    }
  }
  // Add each option
  for _, option := range options {
    cards = option(cards)
  }
  return cards
}

// DefaultSort is a sorting option that can be passed into the New function
func DefaultSort(cards []Card) []Card{
  sort.Slice(cards, Less(cards))
  return cards
}

// Sort is a sorting option for custom sorting options
func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
  return func(cards []Card) []Card {
    sort.Slice(cards, less(cards))
    return cards
  }
}

// Less is a helper method for default sorting options
func Less(cards []Card) func(i, j int) bool {
  return func(i, j int) bool {
    return absoluteRank(cards[i]) < absoluteRank(cards[j])
  }
}
// absoluteRank assigns absolute int for each card
func absoluteRank(c Card) int {
  return int(c.Suit) * int(maxRank) + int(c.Rank)
}

// Shuffle assigns random order to deck
func Shuffle(cards []Card) []Card {
  output := make([]Card, len(cards))
  r := rand.New(rand.NewSource(time.Now().Unix())) // Random source for rand package
  permutation := r.Perm(len(cards))
  for i, j := range permutation {
    output[i] = cards[j]
  }
  return output
}

// Jokers add a given number of jokers to a new deck
func Jokers(n int) func([]Card) []Card {
  return func(cards []Card) []Card {
    for i := 0; i < n; i++ {
      cards = append(cards, Card{
        Rank: Rank(i), // Gives us a differentiated value to each Joker
        Suit: Joker,
      })
    }
    return cards
  }
}

// Filter removed designated cards from a deck
func Filter(f func(card Card) bool) func([]Card) []Card {
  return func(cards []Card) []Card {
    var output []Card
    for _, c := range cards {
      if !f(c) { // If card is not in forbidden slice, append
        output = append(output, c)
      }
    }
    return output
  }
}
