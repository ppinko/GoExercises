package main

import "fmt"

// https://exercism.org/tracks/go/exercises/proverb/edit

/*
For want of a nail the shoe was lost.
For want of a shoe the horse was lost.
For want of a horse the rider was lost.
For want of a rider the message was lost.
For want of a message the battle was lost.
For want of a battle the kingdom was lost.
And all for the want of a nail.
*/

func Proverb(rhyme []string) []string {
	var ret []string
	for i := 0; i < len(rhyme); i++ {
		if i != len(rhyme)-1 {
			line := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
			ret = append(ret, line)
		} else {
			line := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
			ret = append(ret, line)
		}
	}
	return ret
}

func main() {
	fmt.Println()

	test := []string{"nail", "shoe", "horse", "rider", "message", "battle", "kingdom"}
	proverb := Proverb(test)
	fmt.Println(proverb)
	fmt.Println(len(proverb))
}
