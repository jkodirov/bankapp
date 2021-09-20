package payment

import "github.com/jkodirov/bankapp/pkg/bank/types"

func Max(payments []types.Payment) types.Payment {
	max := payments[0]
	for _, payment := range payments {
		if max.Amount < payment.Amount {
			max = payment
		}
	}
	return max
}
