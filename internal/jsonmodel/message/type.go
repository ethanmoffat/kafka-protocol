package message

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Type int

const (
	Request Type = iota
	Response
	Header
	Metadata
	Data
)

func (t Type) String() (s string) {
	switch t {
	case Request:
		s = "Request"
	case Response:
		s = "Response"
	case Header:
		s = "Header"
	case Metadata:
		s = "Metadata"
	case Data:
		s = "Data"
	}
	return
}

func (t *Type) From(s string) error {
	switch strings.ToLower(s) {
	case "request":
		*t = Request
	case "response":
		*t = Response
	case "header":
		*t = Header
	case "metadata":
		*t = Metadata
	case "data":
		*t = Data
	default:
		return fmt.Errorf("s is an unrecognized Type: %s", s)
	}
	return nil
}

func (t *Type) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return t.From(s)
}
