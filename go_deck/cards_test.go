package go_deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Suit: Heart, Rank: Ace})
	fmt.Println(Card{Suit: Spade, Rank: Three})
	fmt.Println(Card{Suit: Club, Rank: Ace})
	fmt.Println(Card{Suit: Diamond, Rank: Two})
	fmt.Println(Card{Suit: Joker})

	//Output:
	//Ace of Hearts
	//Three of Spades
	//Ace of Clubs
	//Two of Diamonds
	//Joker
}

func TestAmount(t *testing.T) {
	cards := New()
	if len(cards) != 52 {
		t.Fatal("nope")
	}
}

func TestSort(t *testing.T) {
	cards := New(DefaultSort)
	for _, card := range cards {
		fmt.Println(card)
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
		t.Error("expected 3 jokers, received: ", count)
	}
}
func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))

	fmt.Println(cards)
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Expected all two and threes to be filtered out")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	//13 ranks * 4 suits * 3 decks
	if len(cards) != 13*4*3 {
		t.Errorf("expected %d cards, reveived %d cards", 13*4*3, len(cards))
	}
}
