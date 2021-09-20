package stats

import "github.com/jkodirov/bankapp/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {
	sum := 0
	for _, payment := range payments {
		sum += int(payment.Amount)
	}
	avg := types.Money(sum / len(payments))
	return avg
}