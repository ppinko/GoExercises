package main

// https://exercism.org/tracks/go/exercises/allergies

func KnownAllergies() ([]uint, []string) {
	code := make([]uint, 8, 8)
	var bit uint = 1
	for i, _ := range code {
		code[i] = bit
		bit = bit << 1
	}
	allergies := []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}
	return code, allergies
}

func Allergies(allergies uint) []string {
	codes, types := KnownAllergies()
	var haveAllergies []string
	for i, v := range codes {
		if allergies&v == v {
			haveAllergies = append(haveAllergies, types[i])
		}
	}
	return haveAllergies
}

func AllergicTo(allergies uint, allergen string) bool {
	codes, types := KnownAllergies()
	for i, v := range types {
		if allergen == v && allergies&codes[i] == codes[i] {
			return true
		}
	}
	return false
}

func main() {

}
