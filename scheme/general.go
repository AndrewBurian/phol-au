package scheme

import (
)

type Name struct {
	Value string `xml:",chardata"`
}

type IDString struct {
	Value string `xml:",chardata"`
}

type TextorXHTML struct {
}

type URLorEmpty struct {
}

type Paragraph struct {
}

type Dollars struct {
}
type DollarsCents struct {
	Value string `xml:",chardata"`
}

type Percent struct {
	Value string `xml:",chardata"`
}

type Days struct {
	Value string `xml:",chardata"`
}

type DollarsOrPercent struct {
	Type string `xml:",attr"`
}