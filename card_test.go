package deck

import (
  "fmt"
  "testing"
  )

func ExampleCard() {
  fmt.Println(Card{Rank:King, Suit:Spade})
  fmt.Println(Card{Rank:Ace, Suit:Diamond})
  fmt.Println(Card{Rank:Seven, Suit:Heart})
  fmt.Println(Card{Rank:Four, Suit:Spade})
  fmt.Println(Card{Rank:King, Suit:Club})
  fmt.Println(Card{Suit:Joker})


  // Output:
  // King of Spades
  // Ace of Diamonds
  // Seven of Hearts
  // Four of Spades
  // King of Clubs
  // Joker
}

func TestNewDeck(t *testing.T) {
  cards := New()
  if len(cards) != 52  { // Full deck of cards
      t.Error("Incorrect number of cards in new deck")
    }
}

func TestDefaultSort(t *testing.T) {
  cards := New(DefaultSort)
  expected := Card{Rank: Ace, Suit: Club}
  if cards[0] != expected {
    t.Error("Cards not default sorted, expected Ace of Clubs as first card")
  }
}

func TestSort(t *testing.T) {
  cards := New(Sort(Less))
  expected := Card{Rank: Ace, Suit: Club}
  if cards[0] != expected {
    t.Error("Cards not default sorted, expected Ace of Clubs as first card")
  }
}

func TestJokers(t *testing.T) {
  cards := New(Jokers(3))
  count := 0
  for _, c := range cards {
    if c.Suit == Joker {
      count++
    }
  }
  if count != 3 {
    t.Error("Expected 3 jokers. Received:", count)
  }
}
