package days

func Days(dayInt int) (day string) {
	switch dayInt {
	case 1:
		return "Monday"
	case 2:
		return "TuesDay"
	case 3:
		return "WednesDay"
	case 4:
		return "ThursDay"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	default:
		return "Sunday"

	}
}
