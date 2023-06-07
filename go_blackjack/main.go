package main

import (
	"fmt"
	go_deck "go_cards"
	"strings"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Hand []go_deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func (h Hand) DealerString() string {
	return h[0].String() + ", **HIDDEN**"
}
func (h Hand) DealerHidden() string {
	return "/home/tmwpwo/go_cardgame/go_blackjack/CARDS/HIDDEN.png"
}

func (h Hand) firstCard() string {
	return "/home/tmwpwo/go_cardgame/go_blackjack/CARDS/" + h[0].String() + ".png"
}
func (h Hand) secondCard() string {
	return "/home/tmwpwo/go_cardgame/go_blackjack/CARDS/" + h[1].String() + ".png"
}
func (h Hand) thirdCard() string {
	return "/home/tmwpwo/go_cardgame/go_blackjack/CARDS/" + h[2].String() + ".png"
}

func (h Hand) Score() int {
	minscore := h.MinScore()
	if minscore > 11 {
		return minscore
	}
	for _, c := range h {
		if c.Rank == go_deck.Ace {
			return minscore + 10
		}
	}
	return minscore
}

func (h Hand) MinScore() int {
	score := 0
	for _, c := range h {
		score += min(int(c.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	App := app.New()
	window := App.NewWindow("blackjack")

	cards := go_deck.New(go_deck.Deck(1), go_deck.Shuffle)
	var card go_deck.Card
	var player, dealer Hand
	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&player, &dealer} {
			card, cards = draw(cards)
			*hand = append(*hand, card)
		}
	}
	var input string

	imgPlayer1 := canvas.NewImageFromFile(player.firstCard())
	imgPlayer2 := canvas.NewImageFromFile(player.secondCard())
	imgDealer1 := canvas.NewImageFromFile(dealer.firstCard())
	imgDealer2 := canvas.NewImageFromFile(dealer.DealerHidden())

	imgPlayer1.FillMode = canvas.ImageFillOriginal
	imgPlayer2.FillMode = canvas.ImageFillOriginal
	imgDealer1.FillMode = canvas.ImageFillOriginal
	imgDealer2.FillMode = canvas.ImageFillOriginal

	button1 := widget.NewButton("stand", func() {
		imgDealer2.File = dealer.secondCard()
		imgDealer2.Refresh()
	})

	playerContainer := container.NewHBox(
		imgPlayer1,
		imgPlayer2,
	)

	dealerContainer := container.NewHBox(
		imgDealer1,
		imgDealer2,
	)

	bigContainer := container.NewVBox(
		playerContainer,
		dealerContainer,
		button1,
	)

	button2 := widget.NewButton("hit", func() {

		card, cards = draw(cards)
		player = append(player, card)
		newCard := canvas.NewImageFromFile(player.thirdCard())
		fmt.Println(player.thirdCard())
		newCard.FillMode = canvas.ImageFillOriginal
		bigContainer.Add(newCard)
		bigContainer.Refresh()

	})

	content := container.NewVBox(bigContainer, button2)

	window.SetContent(content)

	window.ShowAndRun()
	for input != "s" {
		fmt.Println("player: ", player)
		fmt.Println("dealer: ", dealer.DealerString())
		fmt.Println("What will you do? h for hit or s for stand")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			card, cards = draw(cards)
			player = append(player, card)
		}
	}
	for dealer.Score() <= 16 || (dealer.Score() == 17 && dealer.MinScore() != 17) {
		card, cards = draw(cards)
		dealer = append(dealer, card)
	}
	playerScore, dealerScore := player.Score(), dealer.Score()
	fmt.Println("==Final Hands==")
	fmt.Println("player: ", player, "\nScore: ", playerScore)
	fmt.Println("dealer: ", dealer, "\nScore: ", dealerScore)
	switch {
	case playerScore > 21:
		fmt.Println("you lost")
	case dealerScore > 21:
		fmt.Println("you won!")
	case playerScore > dealerScore:
		fmt.Println("you won!")
	case playerScore < dealerScore:
		fmt.Println("you lost!")
	case playerScore == dealerScore:
		fmt.Println("draw")
	}

}

func draw(cards []go_deck.Card) (go_deck.Card, []go_deck.Card) {
	return cards[0], cards[1:]
}
