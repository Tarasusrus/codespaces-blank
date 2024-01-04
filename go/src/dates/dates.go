package dates

const DaysInWeek int = 7

func weeksToDays(weeks int) int{
	return weeks * DaysInWeek
}

func daysToWeek(days int) float64 {
	return float64(days) / float64(DaysInWeek)
}
