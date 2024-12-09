package xrunes

import "unicode"

// HyphenMinusParams defines the parameters for transforming text with hyphen-minus characters.
// PreserveCase indicates whether the original case of the text should be preserved.
// Screaming specifies whether the text should be transformed to uppercase.
type HyphenMinusParams struct {
	// PreserveCase indicates whether the original case of the text should be preserved.
	PreserveCase bool
	// Screaming specifies whether the text should be transformed to uppercase.
	Screaming bool
}

// HyphenMinusOption is a function type that modifies the options for HyphenMinusParams.
// It allows for functional options to be passed to configure the behavior of HyphenMinusParams.
type HyphenMinusOption func(options *HyphenMinusParams)

// Screaming sets the Screaming field of the given HyphenMinusParams to true.
// This function is used to enable the "screaming" transformation.
//
// params: A pointer to a HyphenMinusParams struct.
func Screaming(params *HyphenMinusParams) {
	params.Screaming = true
}

// PreserveCase sets the PreserveCase field of the given HyphenMinusParams
// to true, ensuring that the case of characters is preserved during
// transformations.
//
// Parameters:
//   - params: A pointer to a HyphenMinusParams struct that will be modified.
//
// - options: A variadic list of HyphenMinusOption to customize the transformation.
//
// Returns:
// - A new slice of runes with the transformations applied.
func Underscore(runes []rune, options ...HyphenMinusOption) []rune {
	if len(runes) == 0 {
		return runes
	}

	sb := make([]rune, 0)
	last := rune(0)
	params := &HyphenMinusParams{}
	for _, option := range options {
		option(params)
	}

	for _, r := range runes {
		if unicode.IsLetter(r) {
			if unicode.IsUpper(r) {
				if unicode.IsLetter(last) && unicode.IsLower(last) {
					sb = append(sb, '_')
					if params.PreserveCase || params.Screaming {
						sb = append(sb, r)
						last = r
						continue
					}

					sb = append(sb, unicode.ToLower(r))
					last = r
					continue
				}

				if params.PreserveCase || params.Screaming {
					sb = append(sb, r)
					last = r
					continue
				}

				sb = append(sb, unicode.ToLower(r))
				last = r
				continue
			}

			if params.Screaming {
				sb = append(sb, unicode.ToUpper(r))
			} else if params.PreserveCase {
				sb = append(sb, r)
			} else {
				sb = append(sb, unicode.ToLower(r))
			}

			last = r
			continue
		}

		if unicode.IsNumber(r) {
			sb = append(sb, r)
			last = r
			continue
		}

		if r == '_' || r == '-' || unicode.IsSpace(r) {
			if len(sb) == 0 {
				continue
			}

			if last == '_' {
				continue
			}

			last = '_'
			sb = append(sb, last)
			continue
		}

	}

	if len(sb) > 0 && sb[len(sb)-1] == '_' {
		sb = sb[:len(sb)-1]
	}

	return sb
}

// Dasherize converts a slice of runes into a dasherized format.
// It inserts hyphens between words and converts letters to lowercase by default.
// Options can be provided to preserve the case or convert all letters to uppercase.
//
// Parameters:
//
//	runes: A slice of runes to be transformed.
//	options: Variadic HyphenMinusOption to customize the transformation.
//
// Returns:
//
//	A new slice of runes in dasherized format.
func Dasherize(runes []rune, options ...HyphenMinusOption) []rune {
	if len(runes) == 0 {
		return runes
	}

	sb := make([]rune, 0)
	last := rune(0)
	params := &HyphenMinusParams{}
	for _, option := range options {
		option(params)
	}

	for _, r := range runes {
		if unicode.IsLetter(r) {
			if unicode.IsUpper(r) {
				if unicode.IsLetter(last) && unicode.IsLower(last) {
					sb = append(sb, '-')
					if params.PreserveCase || params.Screaming {
						sb = append(sb, r)
						last = r
						continue
					}

					sb = append(sb, unicode.ToLower(r))
					last = r
					continue
				}

				if params.PreserveCase || params.Screaming {
					sb = append(sb, r)
					last = r
					continue
				}

				sb = append(sb, unicode.ToLower(r))
				last = r
				continue
			}

			if params.Screaming {
				sb = append(sb, unicode.ToUpper(r))
			} else if params.PreserveCase {
				sb = append(sb, r)
			} else {
				sb = append(sb, unicode.ToLower(r))
			}

			last = r
			continue
		}

		if unicode.IsNumber(r) {
			sb = append(sb, r)
			last = r
			continue
		}

		if r == '_' || r == '-' || unicode.IsSpace(r) {
			if len(sb) == 0 {
				continue
			}

			if last == '-' {
				continue
			}

			last = '-'
			sb = append(sb, last)
			continue
		}

	}

	if len(sb) > 0 && sb[len(sb)-1] == '-' {
		sb = sb[:len(sb)-1]
	}

	return sb
}

// CamelCase converts a slice of runes into camel case format.
// It processes the input runes and ensures that the first letter is lowercase,
// letters following an underscore are uppercase, and all other letters are
// in their original case unless they are part of a sequence of uppercase letters,
// in which case they are converted to lowercase. Numbers are preserved as is.
// Non-letter and non-number characters are treated as underscores.
//
// Example:
// Input:  []rune("hello_world_example")
// Output: []rune("helloWorldExample")
func CamelCase(runes []rune) []rune {
	if len(runes) == 0 {
		return runes
	}

	sb := make([]rune, 0)
	last := rune(0)
	for i, r := range runes {
		if unicode.IsLetter(r) {
			if i == 0 {
				sb = append(sb, unicode.ToLower(r))
				last = r
				continue
			}

			if last == '_' {
				sb = append(sb, unicode.ToUpper(r))
				last = r
				continue
			}

			if unicode.IsUpper(r) {
				if unicode.IsUpper(last) {
					sb = append(sb, unicode.ToLower(r))
					last = r
					continue
				}

				sb = append(sb, r)
			}

			sb = append(sb, r)

			continue
		}

		if unicode.IsNumber(r) {
			sb = append(sb, r)
			last = r
			continue
		}

		last = '_'
	}

	return sb
}

// PascalCase converts a slice of runes to PascalCase format.
// It capitalizes the first letter and any letter that follows an underscore,
// while converting other letters to lowercase if they follow an uppercase letter.
// Non-letter characters are treated as underscores and are not included in the result.
//
// Example:
// Input:  []rune("hello_world")
// Output: []rune("HelloWorld")
func PascalCase(runes []rune) []rune {
	if len(runes) == 0 {
		return runes
	}

	sb := make([]rune, 0)
	last := rune(0)
	for i, r := range runes {
		if unicode.IsLetter(r) {
			if i == 0 {
				sb = append(sb, unicode.ToUpper(r))
				last = r
				continue
			}

			if last == '_' {
				sb = append(sb, unicode.ToUpper(r))
				last = r
				continue
			}

			if unicode.IsUpper(r) {
				if unicode.IsUpper(last) {
					sb = append(sb, unicode.ToLower(r))
					last = r
					continue
				}

				sb = append(sb, r)
			}

			sb = append(sb, r)

			continue
		}

		if unicode.IsNumber(r) {
			sb = append(sb, r)
			last = r
			continue
		}

		last = '_'
	}

	return sb
}
