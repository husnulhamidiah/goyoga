package goyoga

import yoga "github.com/husnulhamidiah/goyoga/gen"

type Config struct {
	ref      yoga.ConfigRef
	constRef yoga.ConfigConstRef
}

func wrapConfig(ref yoga.ConfigRef) *Config {
	if ref == nil {
		return nil
	}
	return &Config{ref: ref, constRef: yoga.ConfigConstRef(ref)}
}

func wrapConstConfig(ref yoga.ConfigConstRef) *Config {
	if ref == nil {
		return nil
	}
	return &Config{constRef: ref}
}

func (c *Config) yogaRef() yoga.ConfigRef {
	if c == nil {
		return nil
	}
	return c.ref
}

func (c *Config) yogaConstRef() yoga.ConfigConstRef {
	if c == nil {
		return nil
	}
	if c.constRef != nil {
		return c.constRef
	}
	return yoga.ConfigConstRef(c.ref)
}

type Size = yoga.Size
type Value yoga.Value

var ValueAuto = Value(yoga.ValueAuto)
var ValueUndefined = Value(yoga.ValueUndefined)
var ValueZero = Value(yoga.ValueZero)

func (v Value) Equal(other Value) bool {
	if v.IsUndefined() && other.IsUndefined() {
		return true
	}
	return yoga.Value(v) == yoga.Value(other)
}

func (v Value) IsAuto() bool {
	return v.Equal(ValueAuto)
}

func (v Value) IsUndefined() bool {
	raw := yoga.Value(v)
	return raw != raw
}

func (v Value) NotEqual(other Value) bool {
	return !v.Equal(other)
}

type Layout struct {
	Left        float32
	Top         float32
	Right       float32
	Bottom      float32
	Width       float32
	Height      float32
	Direction   Direction
	HadOverflow bool
}

type Logger = yoga.Logger
type CloneNodeFunc = yoga.CloneNodeFunc
type DirtiedFunc = yoga.DirtiedFunc
type MeasureFunc = yoga.MeasureFunc
type MinContentMeasureFunc = yoga.MinContentMeasureFunc
type BaselineFunc = yoga.BaselineFunc
