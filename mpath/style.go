package mpath

import (
	"strings"
)

// PathStyle is a formatter and parser interface for paths.
type PathStyle interface {
	// Strict indicates if this path style uses strong types.
	//
	// A path style which doesn't use strong types may return all values as
	// keys, leaving it up to the consumer of the path to decide how to handle
	// the key based on the type they are operating on.
	Strict() bool

	// Format converts a set of elements to a string in this style.
	Format([]Element) string

	// Parse parses a set of elements out of the given string.
	Parse(string) ([]Element, error)
}

// DotNotation is a simple dot and brackets style path notation.
type DotNotation struct{}

// Strict indicates if this path style uses strong types.
//
// For DotNotation style paths, Strict() always returns true.
func (dn DotNotation) Strict() bool {
	return true
}

// Format converts a set of elements to a string in this style.
func (dn DotNotation) Format(elements []Element) string {
	if len(elements) == 0 {
		return ""
	}

	b := &strings.Builder{}
	if elements[0].Type() == IndexType {
		b.WriteRune('[')
		b.WriteString(elements[0].String())
		b.WriteRune(']')
	} else {
		b.WriteString(elements[0].String())
	}
	for _, e := range elements[1:] {
		if e.Type() == IndexType {
			b.WriteRune('[')
			b.WriteString(e.String())
			b.WriteRune(']')
		} else {
			b.WriteRune('.')
			b.WriteString(e.String())
		}
	}
	return b.String()
}

// Parse parses a string into a dot-notation path.
func (dn DotNotation) Parse(value string) ([]Element, error) {
	if value == "" {
		return nil, nil
	}

	var elements []Element
	runes := []rune(value)

	i, first := 0, true
	for i < len(runes) {
		switch runes[i] {
		default:
			// Error, missing separator
			if !first {
				return nil, ErrMissingSep
			}
			i--
			fallthrough
		case '.':
			elem, n, err := dn.parseKey(runes, i+1)
			if err != nil {
				return nil, err
			}
			elements = append(elements, elem)
			i = n
		case '[':
			elem, n, err := dn.parseIndex(runes, i+1)
			if err != nil {
				return nil, err
			}
			elements = append(elements, elem)
			i = n
		case ']':
			// Error - mismatched closing bracket
			return nil, ErrUnmatchedCloseBracket
		}
		first = false
	}

	return elements, nil
}

func (dn DotNotation) parseKey(runes []rune, i int) (Element, int, error) {
	builder := &strings.Builder{}
	for ; i < len(runes); i++ {
		switch runes[i] {
		case '.', '[', ']':
			return Key(builder.String()), i, nil
		case '\\':
			i++
			if i >= len(runes) {
				return nil, i, ErrInvalidEscape
			}
			builder.WriteRune(runes[i])
		default:
			builder.WriteRune(runes[i])
		}
	}
	return Key(builder.String()), i, nil
}

func (dn DotNotation) parseIndex(runes []rune, i int) (Element, int, error) {
	start := i
	for ; i < len(runes); i++ {
		if runes[i] == ']' {
			elem, err := ParseIndex(runes[start:i])
			return elem, i + 1, err
		}
	}
	return nil, i, ErrUnmatchedOpenBracket
}
