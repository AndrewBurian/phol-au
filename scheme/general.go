package scheme

import (
	"encoding/xml"
	"fmt"
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

type BoolValue bool

func (b *BoolValue) UnmarshalXML(d *xml.Decoder, e xml.StartElement) error {
	var v struct {
		Data string `xml:",chardata"`
	}
	if err := d.DecodeElement(&v, &e); err != nil {
		return fmt.Errorf("unable to decode value for boolean element: %w", err)
	}

	bVal, err := stringToBool(v.Data)
	if err != nil {
		return fmt.Errorf("error converting element text to boolean: %w", err)
	}
	if bVal {
		*b = true
		return nil
	}
	*b = false
	return nil

}

func stringToBool(s string) (bool, error) {
	switch s {
	case "true", "1":
		return true, nil
	case "false", "0":
		return false, nil
	}
	return false, fmt.Errorf("illegal boolean value: %s", s)
}
