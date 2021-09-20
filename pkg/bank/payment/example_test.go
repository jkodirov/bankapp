package payment

import (
	"github.com/jkodirov/bankapp/pkg/bank/types"
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
