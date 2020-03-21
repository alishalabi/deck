# Deck Of Playing Card
The purpose of this project is to create a reusable API to generate a deck of playing cards. This deck can easily be utilized in any card game projects (Poker, Go Fish, Blackjack).

## Usage
- Package can be imported into any Golang project by using `go get https://github.com/alishalabi/deck`
- Package can be cloned locally using `git clone https://github.com/alishalabi/deck.git`

## Public methods
- `New(options ...func([]Card) []Card) []Card`
  - Creates a new deck, pass in any number of options
- `DefaultSort(cards []Card) []Card`
  - Reset cards to default order
- `Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card`
  - Allow user to implement custom sorting options
- `Shuffle(cards []Card) []Card`
  - Assigns random order to each card in deck (seed based on time/date)
- `Jokers(n int) func([]Card) []Card`
  - Adds a given number of jokers to deck
- `Filter(f func(card Card) bool) func([]Card)`
  - Filters out given suits or ranks from deck
- `Deck(n int) func([]Card) []Card`
  - Add multiple decks

## Iterations
- [x] Create standard deck struct
- [x] Create method to generate standard deck
- [x] Create method to sort deck
- [x] Create method to shuffle deck
- [x] Create method to add jokers
- [x] Create method to filter out cards
- [x] Create method to multiple decks
