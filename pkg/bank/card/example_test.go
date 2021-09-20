package card

import (
	"github.com/jkodirov/bankapp/pkg/bank/types"
	"fmt"
)

func ExampleWithdraw_positive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 1000000
}

func ExampleWithdraw_noMoney() {
	result := Withdraw(types.Card{Balance: 0, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 0
}

func ExampleWithdraw_inactive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: false}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 2000000
}

func ExampleWithdraw_limit() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 30_000_00)
	fmt.Println(result.Balance)
	// Output: 2000000
}

func ExampleDeposit_ok() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output: 3000000
}

func ExampleDeposit_active() {
	card := types.Card{Balance: 20_000_00, Active: false}
	Deposit(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleDeposit_overlimit() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&card, 50_000_01)
	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleAddBonus_ok() {
	card := types.Card{Balance: 20_000_00, Active: true, MinBalance: 10_000_00}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 2002465
}

func ExampleAddBonus_inactive() {
	card := types.Card{Balance: 20_000_00, Active: false, MinBalance: 10_000_00}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleAddBonus_negative() {
	card := types.Card{Balance: -20_000_00, Active: true, MinBalance: 10_000_00}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: -2000000
}

func ExampleAddBonus_overlimit() {
	card := types.Card{Balance: 1, Active: true, MinBalance: 10_000_00_00}
	AddBonus(&card, 200, 30, 365)
	fmt.Println(card.Balance)
	// Output: 1
}

func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active: true,
		},
		{
			Balance: 10_000_00,
			Active: true,
		},
		{
			Balance: 10_000_00,
			Active: true,
		},
	}
	fmt.Println(Total(cards))
	//Output:
	//3000000
}

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			PAN: "5555 **** **** 8888",
			Balance: 10_000_00,
			Active: true,
		},
		{
			PAN: "6666 **** **** 8888",
			Balance: 10_000_00,
			Active: true,
		},
		{
			PAN: "7777 **** **** 8888",
			Balance: 10_000_00,
			Active: false,
		},
	}
	sources := PaymentSources(cards)
	for _, source := range sources {
		fmt.Println(source.Number)
	}
	//Output:
	// 5555 **** **** 8888
	// 6666 **** **** 8888
}