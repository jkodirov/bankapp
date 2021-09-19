package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax()  {
	payments := []types.Payment{
		{
			ID: 101,
			Amount: 13,
		},
		{
			ID: 102,
			Amount: 10,
		},
		{
			ID: 103,
			Amount: 14,
		},
	}
	max := Max(payments)
	fmt.Println(max.ID)
	// Output:
	// 103
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