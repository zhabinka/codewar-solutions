package kata

func Hero(bullets, dragons int) bool {
	return int(bullets / dragons) >= 2
}