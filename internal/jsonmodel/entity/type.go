package entity

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Type int

const (
	Unknown Type = iota
	TransactionalId
	ProducerId
	GroupId
	TopicName
	BrokerId
)

func (t Type) String() (s string) {
	switch t {
	case Unknown:
		s = "Unknown"
	case TransactionalId:
		s = "TransactionalId"
	case ProducerId:
		s = "ProducerId"
	case GroupId:
		s = "GroupId"
	case TopicName:
		s = "TopicName"
	case BrokerId:
		s = "BrokerId"
	}
	return
}

func (t *Type) From(s string) error {
	switch strings.ToLower(s) {
	case "":
		fallthrough
	case "unknown":
		*t = Unknown
	case "transactionalid":
		*t = TransactionalId
	case "producerid":
		*t = ProducerId
	case "groupid":
		*t = GroupId
	case "topicname":
		*t = TopicName
	case "brokerid":
		*t = BrokerId
	default:
		return fmt.Errorf("s is an unrecognized EntityType: %s", s)
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
