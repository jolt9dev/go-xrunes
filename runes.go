package xrunes

import (
	"slices"
	"unicode"
)

// Contains checks if the slice of runes `r` is present within the slice of runes `s`.
// It returns true if `r` is found within `s`, otherwise it returns false.
func Contains(s []rune, r []rune) bool {
	return Index(s, r) > -1
}

// ContainsFold reports whether the slice of runes `r` is within the slice of runes `s`,
// using Unicode case-folding to compare runes. It returns true if `r` is found within `s`,
// otherwise it returns false.
func ContainsFold(s []rune, r []rune) bool {
	return IndexFold(s, r) > -1
}

// Equal reports whether two slices of runes are equal.
// It returns true if both slices contain the same elements in the same order.
func Equal(x []rune, y []rune) bool {
	return slices.Equal(x, y)
}

// EqualFold reports whether two slices of runes are equal under Unicode case-folding,
// which is a more general form of case-insensitivity. It returns true if the slices
// have the same length and each corresponding rune pair is either the same or
// considered equal under Unicode case-folding rules. If the lengths of the slices
// differ, it returns false.
func EqualFold(x []rune, y []rune) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			if unicode.IsLetter(x[i]) {
				if equalFoldRune(x[i], y[i]) {
					continue
				}
			}

			return false
		}
	}

	return true
}

// HasPrefix reports whether the slice of runes s begins with prefix.
// It returns true if s starts with prefix, and false otherwise.
//
// Parameters:
//   - s: The slice of runes to be checked.
//   - prefix: The slice of runes to be looked for at the start of s.
//
// Returns:
//   - bool: true if s starts with prefix, false otherwise.
func HasPrefix(s []rune, prefix []rune) bool {
	if len(s) < len(prefix) {
		return false
	}

	return slices.Equal(s[:len(prefix)], prefix)
}

// HasPrefixFold reports whether the slice of runes s begins with prefix,
// using Unicode case-folding to compare runes. It returns true if s starts
// with prefix, and false otherwise.
func HasPrefixFold(s []rune, prefix []rune) bool {
	sl := len(s)
	rl := len(prefix)

	if rl == 0 {
		return true
	}

	if sl < rl {
		return false
	}

	return EqualFold(s[:rl], prefix)
}

// HasSuffix reports whether the slice of runes s ends with suffix.
// It returns true if s ends with suffix, and false otherwise.
func HasSuffix(s []rune, suffix []rune) bool {
	if len(s) < len(suffix) {
		return false
	}

	return slices.Equal(s[len(s)-len(suffix):], suffix)
}

// HasSuffixFold reports whether the slice of runes s ends with suffix,
// using Unicode case-folding to compare runes. It returns true if s ends
// with suffix, and false otherwise.
func HasSuffixFold(s []rune, suffix []rune) bool {
	sl := len(s)
	rl := len(suffix)

	if rl == 0 {
		return true
	}

	if sl < rl {
		return false
	}

	return EqualFold(s[sl-rl:], suffix)
}

// IndexRune returns the index of the first occurrence of the rune r in the slice s.
// It returns -1 if r is not present in s.
func IndexRune(s []rune, r rune) int {
	for i, c := range s {
		if c == r {
			return i
		}
	}

	// Not found
	return -1
}

// IndexRuneFold returns the index of the first occurrence of the rune r in the slice s,
// using Unicode case-folding to compare runes. It returns -1 if r is not present in s.
func IndexRuneFold(s []rune, r rune) int {
	if len(s) == 0 {
		return -1
	}

	for i, c := range s {
		if c == r {
			return i
		}

		if unicode.IsLetter(c) {
			if equalFoldRune(c, r) {
				return i
			}
		}
	}

	return -1
}

// Index returns the index of the first occurrence of the slice of runes r in the slice of runes s.
// It returns -1 if r is not present in s.
func Index(s []rune, r []rune) int {
	sl := len(s)
	rl := len(r)
	if rl == 0 {
		return 0
	}

	if sl < rl {
		return -1
	}

	for i := 0; i < sl; i++ {
		if i+rl > sl {
			return -1
		}

		for j, y := range r {
			x := s[i+j]
			if x == y {
				if j == rl-1 {
					return i
				}

				continue
			}

			break
		}
	}

	return -1
}

// IndexFold returns the index of the first occurrence of the slice of runes r in the slice of runes s,
// using Unicode case-folding to compare runes. It returns -1 if r is not present in s.
func IndexFold(s []rune, r []rune) int {
	sl := len(s)
	rl := len(r)
	if rl == 0 {
		return 0
	}

	for i := 0; i < sl; i++ {
		if i+(rl) > sl {
			return -1
		}

		for j, y := range r {
			x := s[i+j]
			if x == y {
				if j == rl-1 {
					return i
				}

				continue
			}

			if unicode.IsLetter(x) {
				if equalFoldRune(x, y) {
					if j == rl-1 {
						return i
					}

					continue
				}
			}

			break
		}
	}

	return -1
}

// Trim returns a slice of the runes in s with all leading and trailing
// Unicode code points contained in cutset removed. It calls TrimLeft
// and TrimRight to perform the trimming.
func Trim(s []rune, cutset []rune) []rune {
	return TrimLeft(TrimRight(s, cutset), cutset)
}

// TrimLeft removes all leading Unicode code points contained in cutset from the slice s.
// It returns a new slice with the leading cutset code points removed.
//
// Parameters:
//   - s: A slice of runes from which leading cutset runes will be removed.
//   - cutset: A slice of runes that specifies the set of leading runes to be removed.
//
// Returns:
//
//	A new slice of runes with the leading cutset runes removed.
//
// Example:
//
//	s := []rune("!!!Hello, World!!!")
//	cutset := []rune("!")
//	result := TrimLeft(s, cutset) // result will be []rune("Hello, World!!!")
func TrimLeft(s []rune, cutset []rune) []rune {
	if len(s) == 0 {
		return s
	}

	if len(cutset) == 0 {
		return s
	}

	start := 0

	for i := 0; i < len(s); i++ {
		if slices.Contains(cutset, s[i]) {
			start++
			continue
		}

		break
	}

	if start == 0 {
		return s
	}

	return s[start:]
}

// TrimRight removes all trailing Unicode code points contained in cutset from the slice s.
// It returns a new slice with the trailing cutset runes removed.
//
// Parameters:
//   - s: A slice of runes from which trailing runes will be trimmed.
//   - cutset: A slice of runes to be trimmed from the end of s.
//
// Returns:
//
//	A new slice of runes with the trailing runes from cutset removed.
//
// Example:
//
//	s := []rune("Hello, World!!!")
//	cutset := []rune("!")
//	result := TrimRight(s, cutset) // result will be []rune("Hello, World")
func TrimRight(s []rune, cutset []rune) []rune {
	if len(s) == 0 {
		return s
	}

	if len(cutset) == 0 {
		return s
	}

	end := len(s) - 1

	for i := len(s) - 1; i >= 0; i-- {
		if slices.Contains(cutset, s[i]) {
			end = i
			continue
		}

		break
	}

	if end == (len(s) - 1) {
		return s
	}

	return s[:end]
}

// IsSpace checks if all runes in the given slice are whitespace characters.
// It returns true if all runes are whitespace, and false otherwise.
func IsSpace(s []rune) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}

	return true
}

// IsEmptySpace checks if the given slice of runes is empty or consists entirely of space characters.
// It returns true if the slice is empty or if all runes in the slice are space characters.
func IsEmptySpace(s []rune) bool {
	if len(s) == 0 {
		return true
	}

	return IsSpace(s)
}

func equalFoldRune(x, y rune) bool {
	xx := unicode.SimpleFold(x)
	if xx == y {
		return true
	}
	yy := unicode.SimpleFold(y)
	return yy == x
}
