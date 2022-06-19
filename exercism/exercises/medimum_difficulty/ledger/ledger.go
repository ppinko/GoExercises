package main

// https://exercism.org/tracks/go/exercises/ledger

import (
	"errors"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	// validate input
	if ((locale != "nl-NL") && (locale != "en-US")) ||
		((currency != "USD") && (currency != "EUR")) {
		return "", errors.New("invalid input")
	}

	// make copy of input slice
	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)

	// sort slice by several factors
	sort.Slice(entriesCopy, func(i, j int) bool {
		return (entriesCopy[i].Date < entriesCopy[j].Date) ||
			((entriesCopy[i].Date == entriesCopy[j].Date) && (entriesCopy[i].Description < entriesCopy[j].Description)) ||
			((entriesCopy[i].Date == entriesCopy[j].Date) && (entriesCopy[i].Description == entriesCopy[j].Description) && (entriesCopy[i].Change < entriesCopy[j].Change))
	})

	var w1, w2, w3 string
	switch locale {
	case "nl-NL":
		{
			w1 = "Datum"
			w2 = "Omschrijving"
			w3 = " | Verandering\n"
		}
	case "en-US":
		{
			w1 = "Date"
			w2 = "Description"
			w3 = " | Change\n"
		}
	}
	s := w1 + strings.Repeat(" ", 10-len(w1)) + " | " + w2 + strings.Repeat(" ", 25-len(w2)) + w3

	for _, et := range entriesCopy {
		if len(et.Date) != 10 {
			return "", errors.New("invalid input")
		}
		d1, d2, d3, d4, d5 := et.Date[0:4], et.Date[4], et.Date[5:7], et.Date[7], et.Date[8:]
		if (d2 != '-') || (d4 != '-') {
			return "", errors.New("invalid input")
		}

		de := et.Description
		if len(de) > 25 {
			de = de[:22] + "..."
		} else {
			de = de + strings.Repeat(" ", 25-len(de))
		}

		var d string
		if locale == "nl-NL" {
			d = d5 + "-" + d3 + "-" + d1
		} else if locale == "en-US" {
			d = d3 + "/" + d5 + "/" + d1
		}

		negative := false
		cents := et.Change
		if cents < 0 {
			cents = cents * -1
			negative = true
		}

		centsStr := strconv.Itoa(cents)
		switch len(centsStr) {
		case 1:
			centsStr = "00" + centsStr
		case 2:
			centsStr = "0" + centsStr
		}
		rest := centsStr[:len(centsStr)-2]

		var parts []string
		for len(rest) > 3 {
			parts = append(parts, rest[len(rest)-3:])
			rest = rest[:len(rest)-3]
		}
		if len(rest) > 0 {
			parts = append(parts, rest)
		}

		beginning, end, del1, del2, neg := "", "", "", "", ""
		if locale == "nl-NL" {
			end, del1, del2, neg = " ", ".", ",", "-"
		} else {
			beginning, del1, del2, neg = "(", ",", ".", ")"
		}

		var a string
		if negative {
			a += beginning
		}
		if currency == "EUR" {
			a += "€" + end
		} else {
			a += "$" + end
		}

		for i := len(parts) - 1; i >= 0; i-- {
			a += parts[i] + del1
		}

		a = a[:len(a)-1]
		a += del2
		a += centsStr[len(centsStr)-2:]
		if negative {
			a += neg
		} else {
			a += " "
		}

		al := utf8.RuneCountInString(a)

		s += d + strings.Repeat(" ", 10-len(d)) + " | " + de + " | " + strings.Repeat(" ", 13-al) + a + "\n"
	}

	return s, nil
}

func main() {}
