package helpers

func Contains(slice []string, element string) bool {
	for _, el := range slice {
		if el == element {
			return true
		}
	}

	return false
}
