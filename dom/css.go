package dom

type Stylesheet struct {
	rules []Rule
}

type Rule struct {
	selectors    []Selector
	declarations []Declaration
}

// ----------------------------------------------------------------------------
type Selector interface {
	selector()
}

type TagSelector struct {
	tag_name string
}

type IdSelector struct {
	id string
}

type ClassSelector struct {
	class string
}

func (*TagSelector) selector()
func (*IdSelector) selector()
func (*ClassSelector) selector()

// ----------------------------------------------------------------------------

type Declaration struct {
	key   string
	value Value
}

type Value interface {
	Value()
}

type Keyword struct {
	keyword string
}

type Color struct {
	r uint8
	g uint8
	b uint8
	a uint8
}

type Size struct {
	length float64
	unit   string
}

func (*Keyword) Value() {}
func (*Color) Value()   {}
func (*Size) Value()    {}

// ----------------------------------------------------------------------------
