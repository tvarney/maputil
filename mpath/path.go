package mpath

// Path is a collection of path elements and an output style.
type Path struct {
	Filename string
	Elements []Element
	Style    PathStyle
}

// New returns a new path with the given style and elements.
func New(style PathStyle, elements ...Element) *Path {
	return &Path{
		Elements: elements,
		Style:    style,
	}
}

// Parse parses a path from the given string.
func Parse(style PathStyle, value string) (*Path, error) {
	elems, err := style.Parse(value)
	if err != nil {
		return nil, err
	}
	return &Path{Style: style, Elements: elems}, nil
}

// Add appends the given element to this path.
func (p *Path) Add(elem Element) *Path {
	p.Elements = append(p.Elements, elem)
	return p
}

// Pop removes the top-most element from this path.
func (p *Path) Pop() *Path {
	if len(p.Elements) == 0 {
		return p
	}
	p.Elements = p.Elements[:len(p.Elements)-1]
	return p
}

// PopN removes the top N elements from this path.
func (p *Path) PopN(n int) *Path {
	if len(p.Elements) < n {
		p.Elements = p.Elements[:0]
		return p
	}
	p.Elements = p.Elements[:len(p.Elements)-n]
	return p
}

// Clear removes all elements from this path.
func (p *Path) Clear() *Path {
	p.Elements = p.Elements[:0]
	return p
}

// Copy returns a copy of the path.
func (p *Path) Copy() *Path {
	if len(p.Elements) == 0 {
		return &Path{
			Style:    p.Style,
			Elements: nil,
		}
	}
	elements := make([]Element, 0, len(p.Elements))
	for _, e := range p.Elements {
		elements = append(elements, e.Copy())
	}
	return &Path{
		Filename: p.Filename,
		Elements: elements,
		Style:    p.Style,
	}
}

// String returns the string representation of this path.
func (p *Path) String() string {
	if p.Filename != "" {
		return p.Filename + ": " + p.Style.Format(p.Elements)
	}
	return p.Style.Format(p.Elements)
}
