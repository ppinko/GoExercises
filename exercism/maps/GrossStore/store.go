package main

import "fmt"

// https://exercism.org/tracks/go/exercises/gross-store

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{"quarter_of_a_dozen": 3,
		"half_of_a_dozen": 6,
		"dozen":           12,
		"small_gross":     120,
		"gross":           144,
		"great_gross":     1728,
	}
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	val, exists := units[unit]
	if !exists {
		return false
	}
	bill[item] = bill[item] + val
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	val, exists := units[unit]
	if !exists {
		return false
	}
	_, exists2 := bill[item]
	if !exists2 {
		return false
	}
	if bill[item]-val < 0 {
		return false
	}
	bill[item] = bill[item] - val
	if bill[item] == 0 {
		delete(bill, item)
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	val, exists := bill[item]
	return val, exists
}

func main() {
	units := Units()
	fmt.Println(units)

	bill := NewBill()
	fmt.Println(bill)

	ok := AddItem(bill, units, "carrot", "dozen")
	fmt.Println(ok)

	ok = RemoveItem(bill, units, "carrot", "dozen")
	fmt.Println(ok)

	bill2 := map[string]int{"carrot": 12, "grapes": 3}
	qty, ok2 := GetItem(bill2, "carrot")
	fmt.Println(qty)
	// Output: 12
	fmt.Println(ok2)
	// Output: true
}
