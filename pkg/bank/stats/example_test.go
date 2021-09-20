package stats

import (
	"fmt"

	"github.com/jkodirov/bankapp/pkg/bank/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID: 1,
			Amount: 120,
		},
		{
			ID: 2,
			Amount: 100,
		},
		{
			ID: 3,
			Amount: 80,
		},
	}
	fmt.Println(Avg(payments))
	//Output:
	// 100
}