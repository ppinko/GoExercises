package main

import (
	"fmt"
)

// https://exercism.org/tracks/go/exercises/expenses

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	records := make([]Record, 0)
	for _, val := range in {
		if predicate(val) {
			records = append(records, val)
		}
	}
	return records
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	f := func(r Record) bool {
		return r.Day >= p.From && r.Day <= p.To
	}
	return f
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise
func ByCategory(c string) func(Record) bool {
	f := func(r Record) bool {
		return r.Category == c
	}
	return f
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	filtered := Filter(in, ByDaysPeriod(p))
	totalAmount := 0.0
	for _, v := range filtered {
		totalAmount += v.Amount
	}
	return totalAmount
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	anyCategory := Filter(in, ByCategory(c))
	if len(anyCategory) == 0 {
		return 0.0, fmt.Errorf("error(unknown category %s)", c)
	}
	filtered := Filter(anyCategory, ByDaysPeriod(p))
	totalAmount := 0.0
	for _, v := range filtered {
		totalAmount += v.Amount
	}
	return totalAmount, nil
}

// Day1Records only returns true for records that are from day 1
func Day1Records(r Record) bool {
	return r.Day == 1
}

func main() {
	fmt.Println()

	records := []Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
	}

	fmt.Println(Filter(records, Day1Records))

	records2 := []Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
		{Day: 26, Amount: 300, Category: "university"},
		{Day: 28, Amount: 1300, Category: "rent"},
	}

	period := DaysPeriod{From: 1, To: 15}

	fmt.Println(Filter(records2, ByDaysPeriod(period)))
	// =>
	// [
	//   {Day: 1, Amount: 15, Category: "groceries"},
	//   {Day: 11, Amount: 300, Category: "utility-bills"},
	//   {Day: 12, Amount: 28, Category: "groceries"},
	// ]

	records3 := []Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
		{Day: 28, Amount: 1300, Category: "rent"},
	}

	fmt.Println(Filter(records3, ByCategory("groceries")))
	// =>
	// [
	//   {Day: 1, Amount: 15, Category: "groceries"},
	//   {Day: 12, Amount: 28, Category: "groceries"},
	// ]

	records4 := []Record{
		{Day: 15, Amount: 16, Category: "entertainment"},
		{Day: 32, Amount: 20, Category: "groceries"},
		{Day: 40, Amount: 30, Category: "entertainment"},
	}

	p1 := DaysPeriod{From: 1, To: 30}
	p2 := DaysPeriod{From: 31, To: 60}

	fmt.Println(TotalByPeriod(records4, p1))
	// => 16

	fmt.Println(TotalByPeriod(records4, p2))
	// => 50

	p3 := DaysPeriod{From: 1, To: 30}
	p4 := DaysPeriod{From: 31, To: 60}

	records5 := []Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
		{Day: 26, Amount: 300, Category: "university"},
		{Day: 28, Amount: 1300, Category: "rent"},
	}

	fmt.Println(CategoryExpenses(records5, p3, "entertainment"))
	// => 0, error(unknown category entertainment)

	fmt.Println(CategoryExpenses(records5, p3, "rent"))
	// => 1300, nil

	fmt.Println(CategoryExpenses(records5, p4, "rent"))
	// => 0, nil
}
