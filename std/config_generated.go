package goyoga

import yoga "github.com/husnulhamidiah/goyoga/gen"

func NewConfig() *Config {
	return wrapConfig(yoga.ConfigNew())
}

func (c *Config) Free() {
	if c == nil || c.ref == nil {
		return
	}
	yoga.ConfigFree(c.ref)
	c.ref = nil
	c.constRef = nil
}

func (c *Config) Destroy() {
	c.Free()
}

func DefaultConfig() *Config {
	return wrapConstConfig(yoga.ConfigGetDefault())
}

func (c *Config) SetUseWebDefaults(enabled bool) {
	if c == nil {
		return
	}
	yoga.ConfigSetUseWebDefaults(c.yogaRef(), enabled)
}

func (c *Config) UseWebDefaults() bool {
	if c == nil {
		return false
	}
	return yoga.ConfigGetUseWebDefaults(c.yogaConstRef())
}

func (c *Config) SetPointScaleFactor(pixelsInPoint float32) {
	if c == nil {
		return
	}
	yoga.ConfigSetPointScaleFactor(c.yogaRef(), pixelsInPoint)
}

func (c *Config) PointScaleFactor() float32 {
	if c == nil {
		return 0
	}
	return yoga.ConfigGetPointScaleFactor(c.yogaConstRef())
}

func (c *Config) SetErrata(errata Errata) {
	if c == nil {
		return
	}
	yoga.ConfigSetErrata(c.yogaRef(), yoga.Errata(errata))
}

func (c *Config) Errata() Errata {
	if c == nil {
		return Errata(0)
	}
	return Errata(yoga.ConfigGetErrata(c.yogaConstRef()))
}

func (c *Config) SetLogger(logger Logger) {
	if c == nil {
		return
	}
	yoga.ConfigSetLogger(c.yogaRef(), logger)
}

func (c *Config) SetExperimentalFeatureEnabled(feature ExperimentalFeature, enabled bool) {
	if c == nil {
		return
	}
	yoga.ConfigSetExperimentalFeatureEnabled(c.yogaRef(), yoga.ExperimentalFeature(feature), enabled)
}

func (c *Config) IsExperimentalFeatureEnabled(feature ExperimentalFeature) bool {
	if c == nil {
		return false
	}
	return yoga.ConfigIsExperimentalFeatureEnabled(c.yogaConstRef(), yoga.ExperimentalFeature(feature))
}

func (c *Config) SetCloneNodeFunc(callback CloneNodeFunc) {
	if c == nil {
		return
	}
	yoga.ConfigSetCloneNodeFunc(c.yogaRef(), callback)
}
