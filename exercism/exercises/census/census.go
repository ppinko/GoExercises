package main

import "fmt"

// https://exercism.org/tracks/go/exercises/census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{Name: name, Age: age, Address: address}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	return r.Address["street"] != "" && r.Name != ""
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Address = nil
	r.Name = ""
	r.Age = 0
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	counter := 0
	for _, r := range residents {
		if r.HasRequiredInfo() {
			counter++
		}
	}
	return counter
}

func main() {
	fmt.Println()

	name := "Matthew Sanabria"
	age := 29
	address := map[string]string{"street": "Main St."}

	NewResident(name, age, address)
	// => &{Matthew Sanabria 29 map[street:Main St.]}

	resident := NewResident(name, age, address)

	resident.HasRequiredInfo()
	// => true

	resident.Delete()

	fmt.Println(resident)
}
