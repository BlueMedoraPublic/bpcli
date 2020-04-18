package sdk

import (
//"testing"
)

func testPopulatedLogSourceConfig() LogSourceConfig {
	x := newLogSourceConfig()
	x.ID = "abc"
	x.Name = "abc"
	x.Source.ID = "abc"
	x.Source.Name = "abc"
	x.Source.Version = "1.0"
	x.Configuration["ReadFromHead"] = true
	x.Configuration["ReadInterval"] = 2
	x.Configuration["MaxReads"] = "2"
	return x
}
