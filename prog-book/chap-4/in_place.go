package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func noemptyInPlace(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	// the others in underlying array still exist
	return strings[:i]
}

func noemptyInPlaceWithAppend(strings []string) []string {
	out := strings[:0] // make a slice pointer point to the head of strings slice, len = 0
	for _, s := range strings {
		if s != "" {
			out = append(out, s) // cost update len and cap of out slice
		}
	}
	return out
}

func removeKeepOrder(slice []int, i int) []int {
	copy(slice[:i], slice[i+1:])
	return slice[:len(slice)-1]
}

func removeNotKeepOrder(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1] // write the last ele direct to the removed place
	return slice[:len(slice)-1]
}

func noAdjacentDupStrings(strings []string) []string {
	if len(strings) < 2 {
		return strings
	}
	var r, w int
	for r, w = 1, 0; r < len(strings); r++ {
		if strings[r] != strings[w] {
			w++
			strings[w] = strings[r]
		}
	}
	return strings[:w+1]
}

func Squash(s string) string {
	var w, spaces int
	runes := []rune(s)
	for r := 0; r < len(runes); r++ {
		if unicode.IsSpace(runes[r]) {
			if spaces == 0 {
				runes[w] = ' '
				w++
			}
			spaces++
		} else {
			runes[w] = runes[r]
			w++
			spaces = 0
		}
	}
	return string(runes[:w])
}

// Reverse reverses a byte array representing an unicode encoded string.
// Trick: convert string to []byte and we have byte rune as an interface
func Reverse(b []byte) []byte {
	r := bytes.Runes(b)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return []byte(string(r))
}

func main() {
	tests := []string{"abc", "a", "a", "a", "abc", "a", "a"}
	fmt.Println(noAdjacentDupStrings(tests))
}
