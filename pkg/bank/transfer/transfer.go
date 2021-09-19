package transfer

func Total(amount int) int {
	bonus := 0.5
	total := float64(amount) + float64(amount) /100.0 * bonus
	return int(total)
}