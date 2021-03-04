package mpath

import (
	"fmt"
	"strconv"
)

// ElementType encodes the type of path element.
type ElementType int

const (
	// KeyType should be returned by values which operate on map keys.
	KeyType ElementType = 0

	// IndexType should be returned by values which operate on arrays.
	IndexType ElementType = 1
)

// Range tags are used internally by the Range element to mark the start or end
// of the range as not specified.
const (
	RangeTagFull    = 0
	RangeTagNoStart = 1 << 0
	RangeTagNoEnd   = 1 << 1
	RangeTagEmpty   = RangeTagNoStart | RangeTagNoEnd
	RangeTagMask    = 0x3
)

// Element is the interface for elements of a path.
type Element interface {
	Type() ElementType
	String() string
	Copy() Element
}

// Key is a key type which denotes a single element of a map.
type Key string

// Type returns the type of this path element.
func (k Key) Type() ElementType {
	return KeyType
}

// String returns the universal string representation of this element.
func (k Key) String() string {
	return string(k)
}

// Copy returns a copy of this Element.
//
// Because strings are immutable, we can just return the current element.
func (k Key) Copy() Element {
	return k
}

// Index is an index type which denotes a single element of an array.
type Index int

// Type returns the type of this path element.
func (i Index) Type() ElementType {
	return IndexType
}

// String returns the universal string representation of this element.
func (i Index) String() string {
	return strconv.Itoa(int(i))
}

// Copy returns a copy of this Element.
//
// Because integers are immutable, we can just return the current element.
func (i Index) Copy() Element {
	return i
}

// Range is an index type which denotes a set of elements.
type Range struct {
	Start int
	End   int
	Tag   int
}

// RangeFull returns a new Range with both start and end values.
func RangeFull(start, end int) Range {
	return Range{
		Start: start,
		End:   end,
		Tag:   RangeTagFull,
	}
}

// RangeStart returns a new Range with just the start value specified.
func RangeStart(start int) Range {
	return Range{
		Start: start,
		End:   0,
		Tag:   RangeTagNoEnd,
	}
}

// RangeEnd returns a new Range with just the end value specified.
func RangeEnd(end int) Range {
	return Range{
		Start: 0,
		End:   end,
		Tag:   RangeTagNoStart,
	}
}

// RangeEmpty returns a new Range with no values specified.
func RangeEmpty() Range {
	return Range{
		Start: 0,
		End:   0,
		Tag:   RangeTagEmpty,
	}
}

// Type returns the type of this path element.
func (r Range) Type() ElementType {
	return IndexType
}

// String returns the universal string representation of this element.
func (r Range) String() string {
	switch r.Tag & RangeTagMask {
	case RangeTagFull:
		return fmt.Sprintf("%d:%d", r.Start, r.End)
	case RangeTagNoEnd:
		return fmt.Sprintf("%d:", r.Start)
	case RangeTagNoStart:
		return fmt.Sprintf(":%d", r.End)
	}
	return ":"
}

// Copy returns a copy of this Element.
func (r Range) Copy() Element {
	return Range{
		Start: r.Start,
		End:   r.End,
		Tag:   r.Tag,
	}
}

// ArrayEnd is an index type which denotes the element 1 past the end of the
// array.
//
// This type is used to 'insert' an element instead of setting one.
type ArrayEnd struct{}

// Type returns the type of this path element.
func (a ArrayEnd) Type() ElementType {
	return IndexType
}

// String returns the universal string representation of this element.
func (a ArrayEnd) String() string {
	return "-"
}

// Copy returns a copy of this Element.
func (a ArrayEnd) Copy() Element {
	return ArrayEnd{}
}
