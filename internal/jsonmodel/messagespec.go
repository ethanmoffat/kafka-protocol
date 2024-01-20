package jsonmodel

import (
	"github.com/ethanmoffat/kafka-protocol/internal/jsonmodel/entity"
	"github.com/ethanmoffat/kafka-protocol/internal/jsonmodel/listener"
	"github.com/ethanmoffat/kafka-protocol/internal/jsonmodel/message"
	"github.com/ethanmoffat/kafka-protocol/internal/jsonmodel/versions"
)

type MessageSpec struct {
	Name                  string          `json:"name"`
	ValidVersions         versions.Range  `json:"validVersions"`
	DeprecatedVersions    versions.Range  `json:"deprecatedVersions"`
	Fields                []FieldSpec     `json:"fields"`
	ApiKey                *int            `json:"apiKey"`
	Type                  message.Type    `json:"type"`
	CommonStructs         []StructSpec    `json:"commonStructs"`
	FlexibleVersions      versions.Range  `json:"flexibleVersions"`
	Listeners             []listener.Type `json:"listeners"`
	LatestVersionUnstable bool            `json:"latestVersionUnstable"`
}

type FieldSpec struct {
	Name             string
	Versions         versions.Range
	Fields           []FieldSpec
	Type             string
	MapKey           bool
	NullableVersions versions.Range
	Default          any
	Ignorable        bool
	EntityType       entity.Type
	About            string
	TaggedVersions   versions.Range
	FlexibleVersions *versions.Range
	Tag              *int
	ZeroCopy         bool
}

type StructSpec struct {
	Name               string
	Versions           versions.Range
	DeprecatedVersions versions.Range
	Fields             []FieldSpec
}
