package main

import "fmt"

func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	return fmt.Sprintf("%s\nYou have been assigned to table %.3d. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.", Welcome(name), table, direction, distance, neighbor)
}

func main() {
	fmt.Println(Welcome("Christiane"))
	fmt.Println()
	fmt.Println(HappyBirthday("Frank", 58))
	fmt.Println()
	fmt.Println(AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298))
}
