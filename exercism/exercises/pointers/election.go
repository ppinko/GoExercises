package main

import "fmt"

// https://exercism.org/tracks/go/exercises/election-day

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	ptr := &initialVotes
	return ptr
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	} else {
		return *counter
	}
}

// IncrementVoteCount increments the value in a vote counter
func IncrementVoteCount(counter *int, increment int) {
	if counter != nil {
		*counter += increment
	}
}

type ElectionResult struct {
	Name  string
	Votes int
}

// NewElectionResult creates a new election result
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	return &ElectionResult{Name: candidateName, Votes: votes}
}

// DisplayResult creates a message with the result to be displayed
func DisplayResult(result *ElectionResult) string {
	return fmt.Sprintf("%s (%d)", result.Name, result.Votes)
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	results[candidate] -= 1
}

func main() {
	fmt.Println()

	var initialVotes int
	initialVotes = 2

	var counter *int
	counter = NewVoteCounter(initialVotes)
	fmt.Println(*counter == initialVotes) // true

	var votes int
	votes = 3

	var voteCounter *int
	voteCounter = &votes

	VoteCount(voteCounter)
	// Output: 3

	var nilVoteCounter *int
	VoteCount(nilVoteCounter)
	// Output: 0

	IncrementVoteCount(voteCounter, 2)

	fmt.Println(votes == 5)        // true
	fmt.Println(*voteCounter == 5) // true

	var result *ElectionResult
	result = NewElectionResult("Peter", 3)

	fmt.Println(result.Name == "Peter") // true
	fmt.Println(result.Votes == 3)      // true

	var result2 *ElectionResult
	result2 = &ElectionResult{
		Name:  "John",
		Votes: 32,
	}

	DisplayResult(result2)
	// Output: John (32)

	var finalResults = map[string]int{
		"Mary": 10,
		"John": 51,
	}

	DecrementVotesOfCandidate(finalResults, "Mary")

	fmt.Println(finalResults["Mary"])
	// Output: 9
}
