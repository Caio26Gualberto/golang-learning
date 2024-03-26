package adresses

import "strings"

func AdressType(adress string) string {
	validTypes := []string{"street", "avenue", "road", "highway"}

	firstWordAdress := strings.Split(strings.ToLower(adress), " ")[0]

	adressHasValidType := false

	for _, typee := range validTypes {
		if typee == firstWordAdress {
			adressHasValidType = true
		}
	}

	if adressHasValidType {
		return strings.Title(firstWordAdress)
	}

	return "Invalid type!"
}
