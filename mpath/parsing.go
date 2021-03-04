package mpath

import (
	"strconv"
	"strings"
)

// ParseRange parses a range from a set of runes and a separator index.
//
// The given separator index is assumed to be valid; use ParseIndex instead if
// the separator index is unknown.
func ParseRange(runes []rune, sep int) (Element, error) {
	eStart := sep + 1
	sVal := strings.TrimSpace(string(runes[:sep]))
	eVal := strings.TrimSpace(string(runes[eStart:]))

	s, e, t := 0, 0, 0
	var err error
	if len(sVal) > 0 {
		s, err = strconv.Atoi(sVal)
		if err != nil {
			return nil, BadRangeStartError{Value: sVal}
		}
	} else {
		t |= RangeTagNoStart
	}

	if len(eVal) > 0 {
		e, err = strconv.Atoi(eVal)
		if err != nil {
			return nil, BadRangeEndError{Value: eVal}
		}
	} else {
		t |= RangeTagNoEnd
	}

	return Range{Start: s, End: e, Tag: t}, nil
}

// ParseIndex parses an index type from a set of runes.
func ParseIndex(runes []rune) (Element, error) {
	sep := -1
	for i, r := range runes {
		if r == ':' {
			if sep >= 0 {
				return nil, ErrBadRange
			}
			sep = i
		}
	}
	return ParseIndexExt(runes, sep)
}

// ParseIndexExt parses an index type from a set of runes.
//
// The separator index is assumed to be valid; use ParseIndex instead if the
// separator index is unknown.
func ParseIndexExt(runes []rune, sep int) (Element, error) {
	if sep >= 0 {
		return ParseRange(runes, sep)
	}

	val := strings.TrimSpace(string(runes))
	if val == "-" {
		return ArrayEnd{}, nil
	}

	idx, err := strconv.Atoi(val)
	if err != nil {
		return nil, ErrBadIndex
	}
	return Index(idx), nil
}
