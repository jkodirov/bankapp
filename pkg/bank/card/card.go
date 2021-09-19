package card

import "bank/pkg/bank/types"

func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card {
		ID: 1000,
		PAN: "5058 xxxx xxxx 0001",
		Balance: 0,
		Currency: currency,
		Color: color,
		Name: name,
		Active: true,
	}
	return card
}

func Withdraw(card types.Card, amount types.Money) types.Card {
	if !card.Active {
		return card
	}
	if card.Balance < amount {
		return card
	}
	if amount < 0 {
		return card
	}

	if amount > 2_000_000 {
		return card
	}

	card.Balance = card.Balance - amount
	
	return card
}

const limit = 50_000_00
func Deposit(card *types.Card, amount types.Money) {
	if !(*card).Active {
		return
	}
	if amount > limit{
		return
	}

	if amount < 0{
		return
	}
	(*card).Balance += amount
}

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if !(*card).Active {
		return
	}

	if (*card).Balance < 0{
		return
	}

	bonus := int((*card).MinBalance) * percent * daysInMonth / daysInYear /100
	if bonus <= 5_000_00 {
		(*card).Balance += types.Money(bonus)
	}
}

func Total(cards []types.Card) types.Money {
	total := types.Money(0)
	for _, card := range cards {
		total += card.Balance
	}
	return total
}

func PaymentSources(cards []types.Card) []types.PaymentSource {
	sources := []types.PaymentSource{}
	for _, card := range cards {
		if card.Balance > 0 && card.Active {
			sources = append(sources, types.PaymentSource{Type: "card", Number: string(card.PAN), Balance: card.Balance})
		}
	}

	return sources
}