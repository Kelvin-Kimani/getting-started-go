package helper

func GetFirstNames(bookings []map[string]string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}

	return firstNames
}
