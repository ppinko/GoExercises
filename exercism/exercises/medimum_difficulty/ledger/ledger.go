package main

// https://exercism.org/tracks/go/exercises/ledger

import (
	"errors"
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
	// make copy of input slice
	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)

	if len(entries) == 0 {
		if _, err := FormatLedger(currency, "en-US", []Entry{{Date: "2014-01-01", Description: "", Change: 0}}); err != nil {
			return "", err
		}
	}

	m1 := map[bool]int{true: 0, false: 1}
	m2 := map[bool]int{true: -1, false: 1}
	es := entriesCopy

	for len(es) > 1 {
		first, rest := es[0], es[1:]
		for i, e := range rest {
			if (m1[e.Date == first.Date]*m2[e.Date < first.Date]*4 +
				m1[e.Description == first.Description]*m2[e.Description < first.Description]*2 +
				m1[e.Change == first.Change]*m2[e.Change < first.Change]*1) < 0 {
				es[0], es[i+1] = es[i+1], es[0]
			}
		}
		es = es[1:]
	}

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
	default:
		{
			return "", errors.New("")
		}
	}
	s := w1 + strings.Repeat(" ", 10-len(w1)) + " | " + w2 + strings.Repeat(" ", 25-len(w2)) + w3

	type intermediate struct {
		i int
		s string
		e error
	}

	co := make(chan intermediate)
	for i, et := range entriesCopy {
		go func(i int, entry Entry) {
			if len(entry.Date) != 10 {
				co <- intermediate{e: errors.New("")}
			}
			d1, d2, d3, d4, d5 := entry.Date[0:4], entry.Date[4], entry.Date[5:7], entry.Date[7], entry.Date[8:]
			if (d2 != '-') || (d4 != '-') {
				co <- intermediate{e: errors.New("")}
			}

			de := entry.Description
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
			cents := entry.Change
			if cents < 0 {
				cents = cents * -1
				negative = true
			}

			var a string
			if locale == "nl-NL" {
				switch currency {
				case "EUR":
					{
						a += "€ "
					}
				case "USD":
					{
						a += "$ "
					}
				default:
					{
						co <- intermediate{e: errors.New("")}
					}
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
				for i := len(parts) - 1; i >= 0; i-- {
					a += parts[i] + "."
				}
				a = a[:len(a)-1]
				a += ","
				a += centsStr[len(centsStr)-2:]
				if negative {
					a += "-"
				} else {
					a += " "
				}
			} else if locale == "en-US" {
				if negative {
					a += "("
				}
				if currency == "EUR" {
					a += "€"
				} else if currency == "USD" {
					a += "$"
				} else {
					co <- intermediate{e: errors.New("")}
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
				for i := len(parts) - 1; i >= 0; i-- {
					a += parts[i] + ","
				}
				a = a[:len(a)-1]
				a += "."
				a += centsStr[len(centsStr)-2:]
				if negative {
					a += ")"
				} else {
					a += " "
				}
			} else {
				co <- intermediate{e: errors.New("")}
			}

			al := utf8.RuneCountInString(a)

			co <- intermediate{i, d + strings.Repeat(" ", 10-len(d)) + " | " + de + " | " + strings.Repeat(" ", 13-al) + a + "\n", nil}
		}(i, et)
	}

	ss := make([]string, len(entriesCopy))
	for range entriesCopy {
		v := <-co
		if v.e != nil {
			return "", v.e
		}
		ss[v.i] = v.s
	}

	for i := 0; i < len(entriesCopy); i++ {
		s += ss[i]
	}

	return s, nil
}

func main() {}
