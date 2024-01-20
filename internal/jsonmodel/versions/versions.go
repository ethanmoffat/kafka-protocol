package versions

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	All  = &Range{0, int(math.MaxInt16)}
	None = &Range{0, -1}
)

const (
	NoneString = "none"
)

type Range struct {
	lowest  int
	highest int
}

func Default() (*Range, error) {
	return &Range{lowest: 0, highest: -1}, nil
}

func New(lowest, highest int) (*Range, error) {
	if lowest < 0 || highest < 0 {
		return nil, fmt.Errorf("invalid version range: %d to %d", lowest, highest)
	}
	return &Range{lowest: lowest, highest: highest}, nil
}

func Parse(s string, defaultVersions *Range) (*Range, error) {
	if len(s) == 0 {
		return defaultVersions, nil
	}

	trimmed := strings.TrimSpace(s)
	if len(trimmed) == 0 {
		return defaultVersions, nil
	} else if trimmed == None.String() {
		return None, nil
	} else if strings.HasSuffix(trimmed, "+") {
		if v, err := strconv.ParseInt(string(trimmed[0:len(trimmed)-1]), 10, 16); err != nil {
			return nil, err
		} else {
			return New(int(v), math.MaxInt16)
		}
	} else {
		dashIndex := strings.Index(trimmed, "-")
		if dashIndex < 0 {
			if v, err := strconv.ParseInt(trimmed, 10, 16); err != nil {
				return nil, err
			} else {
				return New(int(v), math.MaxInt16)
			}
		}

		low, err_l := strconv.ParseInt(string(trimmed[:dashIndex]), 10, 16)
		if err_l != nil {
			return nil, err_l
		}

		high, err_h := strconv.ParseInt(string(trimmed[dashIndex+1:]), 10, 16)
		if err_h != nil {
			return nil, err_h
		}

		return New(int(low), int(high))
	}
}

func (v Range) Lowest() int {
	return v.lowest
}

func (v Range) Highest() int {
	return v.highest
}

func (v Range) Empty() bool {
	return v.highest < v.lowest
}

func (v Range) String() string {
	if v.Empty() {
		return NoneString
	} else if v.lowest == v.highest {
		return fmt.Sprintf("%v", v.lowest)
	} else if v.highest == math.MaxInt16 {
		return fmt.Sprintf("%v+", v.lowest)
	} else {
		return fmt.Sprintf("%v-%v", v.lowest, v.highest)
	}
}

func (v *Range) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	parsed, err := Parse(s, &Range{0, -1})
	if err == nil {
		v.lowest = parsed.lowest
		v.highest = parsed.highest
	}
	return err
}
