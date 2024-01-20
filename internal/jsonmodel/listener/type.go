package listener

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Type int

const (
	ZookeeperBroker Type = iota
	Broker
	Controller
)

func (t Type) String() (s string) {
	switch t {
	case ZookeeperBroker:
		s = "ZookeeperBroker"
	case Broker:
		s = "Broker"
	case Controller:
		s = "Controller"
	}
	return
}

func (t *Type) From(s string) error {
	switch strings.ToLower(s) {
	case "zkbroker":
		*t = ZookeeperBroker
	case "broker":
		*t = Broker
	case "controller":
		*t = Controller
	default:
		return fmt.Errorf("s is an unrecognized RequestListenerType: %s", s)
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
