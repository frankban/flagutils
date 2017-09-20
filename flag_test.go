// Licensed under the MIT license, see LICENCE file for details.

package flagutils_test

import (
	"flag"
	"testing"

	qt "github.com/frankban/quicktest"

	"github.com/frankban/flagutils"
)

// These assignments are used to ensure that flag.Value is implemented.
var _ flag.Value = (*flagutils.StringSlice)(nil)
var _ flag.Value = (*flagutils.StringMap)(nil)

var sliceTests = []struct {
	about               string
	name                string
	value               string
	defaultValue        []string
	expectedValue       flagutils.StringSlice
	expectedStringValue string
	expectedError       string
}{{
	about:               "single string",
	name:                "single",
	value:               "exterminate",
	expectedValue:       flagutils.StringSlice{"exterminate"},
	expectedStringValue: "exterminate",
}, {
	about:               "multiple strings",
	name:                "multiple",
	value:               "these,are,the,voyages",
	expectedValue:       flagutils.StringSlice{"these", "are", "the", "voyages"},
	expectedStringValue: "these,are,the,voyages",
}, {
	about:               "weird formatting",
	name:                "weird",
	value:               "  these , are,the,  voyages ",
	expectedValue:       flagutils.StringSlice{"these", "are", "the", "voyages"},
	expectedStringValue: "these,are,the,voyages",
}, {
	about:         "default value: with value",
	name:          "def1",
	value:         "exterminate",
	defaultValue:  []string{"default", "not", "used"},
	expectedValue: flagutils.StringSlice{"exterminate"},
}, {
	about:         "default value: without value",
	name:          "def2",
	defaultValue:  []string{"default", "used"},
	expectedValue: flagutils.StringSlice{"default", "used"},
}, {
	about:         "error: empty string",
	name:          "err",
	expectedError: "cannot include empty strings in the list",
}, {
	about:         "error: multiple values with empty string",
	name:          "err",
	value:         ",bad,wolf",
	expectedError: "cannot include empty strings in the list",
}}

func TestSlice(t *testing.T) {
	for _, test := range sliceTests {
		runIsolated(t, test.about, func(c *qt.C) {
			v := flagutils.Slice(test.name, test.defaultValue, "slice usage")
			if test.value != "" || test.defaultValue == nil {
				err := flag.Set(test.name, test.value)
				if test.expectedError == "" {
					c.Assert(err, qt.Equals, nil)
				} else {
					c.Assert(err, qt.ErrorMatches, test.expectedError)
				}
			}
			c.Assert(*v, qt.DeepEquals, test.expectedValue)
		})
	}
}

func TestSliceVar(t *testing.T) {
	for _, test := range sliceTests {
		runIsolated(t, test.about, func(c *qt.C) {
			var v flagutils.StringSlice
			flagutils.SliceVar(&v, test.name, test.defaultValue, "slice usage")
			if test.value != "" || test.defaultValue == nil {
				err := flag.Set(test.name, test.value)
				if test.expectedError == "" {
					c.Assert(err, qt.Equals, nil)
				} else {
					c.Assert(err, qt.ErrorMatches, test.expectedError)
				}
			}
			c.Assert(v, qt.DeepEquals, test.expectedValue)
		})
	}
}

func TestStringSliceSet(t *testing.T) {
	for _, test := range sliceTests {
		runIsolated(t, test.about, func(c *qt.C) {
			if test.defaultValue != nil {
				return
			}
			var v flagutils.StringSlice
			err := v.Set(test.value)
			if test.expectedError == "" {
				c.Assert(err, qt.Equals, nil)
			} else {
				c.Assert(err, qt.ErrorMatches, test.expectedError)
			}
			c.Assert(v, qt.DeepEquals, test.expectedValue)
		})
	}
}

func TestStringSliceString(t *testing.T) {
	for _, test := range sliceTests {
		runIsolated(t, test.about, func(c *qt.C) {
			if test.defaultValue != nil {
				return
			}
			var v flagutils.StringSlice
			v.Set(test.value)
			c.Assert(v.String(), qt.Equals, test.expectedStringValue)
		})
	}
}

var mapTests = []struct {
	about               string
	name                string
	value               string
	defaultValue        map[string]interface{}
	expectedValue       flagutils.StringMap
	expectedStringValue string
	expectedError       string
}{{
	about: "single pair",
	name:  "single",
	value: `{"gisf": true}`,
	expectedValue: flagutils.StringMap{
		"gisf": true,
	},
	expectedStringValue: `{"gisf":true}`,
}, {
	about: "multiple pairs",
	name:  "multiple",
	value: `{"gisf": true, "url": "https://1.2.3.4"}`,
	expectedValue: flagutils.StringMap{
		"gisf": true,
		"url":  "https://1.2.3.4",
	},
	expectedStringValue: `{"gisf":true,"url":"https://1.2.3.4"}`,
}, {
	about: "nested map",
	name:  "nested",
	value: `{"gisf": true, "flags": {"profile": true, "status": true}}`,
	expectedValue: flagutils.StringMap{
		"gisf": true,
		"flags": map[string]interface{}{
			"profile": true,
			"status":  true,
		},
	},
	expectedStringValue: `{"flags":{"profile":true,"status":true},"gisf":true}`,
}, {
	about: "weird formatting",
	name:  "weird",
	value: `  {  "gisf" :  true } `,
	expectedValue: flagutils.StringMap{
		"gisf": true,
	},
	expectedStringValue: `{"gisf":true}`,
}, {
	about:               "empty object",
	name:                "empty",
	value:               `{}`,
	expectedValue:       flagutils.StringMap{},
	expectedStringValue: "{}",
}, {
	about: "no braces: single pair",
	name:  "single",
	value: `"gisf": true`,
	expectedValue: map[string]interface{}{
		"gisf": true,
	},
	expectedStringValue: `{"gisf":true}`,
}, {
	about: "no braces: multiple pairs",
	name:  "multiple",
	value: `"gisf": true, "url": "https://1.2.3.4"`,
	expectedValue: flagutils.StringMap{
		"gisf": true,
		"url":  "https://1.2.3.4",
	},
	expectedStringValue: `{"gisf":true,"url":"https://1.2.3.4"}`,
}, {
	about: "no braces: nested map",
	name:  "nested",
	value: `"gisf": true, "flags": {"profile": true, "status": true}`,
	expectedValue: flagutils.StringMap{
		"gisf": true,
		"flags": map[string]interface{}{
			"profile": true,
			"status":  true,
		},
	},
	expectedStringValue: `{"flags":{"profile":true,"status":true},"gisf":true}`,
}, {
	about: "no braces: weird formatting",
	name:  "weird",
	value: `    "gisf" :  true  `,
	expectedValue: flagutils.StringMap{
		"gisf": true,
	},
	expectedStringValue: `{"gisf":true}`,
}, {
	about: "default value: with value",
	name:  "single",
	value: `{"gisf": true}`,
	defaultValue: map[string]interface{}{
		"answer": 42,
	},
	expectedValue: flagutils.StringMap{
		"gisf": true,
	},
}, {
	about: "default value: without value",
	name:  "single",
	defaultValue: map[string]interface{}{
		"answer": 42,
	},
	expectedValue: flagutils.StringMap{
		"answer": 42,
	},
}, {
	about:               "empty string",
	name:                "empty",
	expectedValue:       flagutils.StringMap{},
	expectedStringValue: "{}",
}, {
	about:               "error: not a map",
	name:                "err",
	value:               "42",
	expectedStringValue: "null",
	expectedError:       "cannot unmarshal JSON: invalid character .*",
}, {
	about:               "error: invalid JSON",
	name:                "err",
	value:               "!",
	expectedStringValue: "null",
	expectedError:       "cannot unmarshal JSON: invalid character .*",
}}

func TestMap(t *testing.T) {
	for _, test := range mapTests {
		runIsolated(t, test.about, func(c *qt.C) {
			v := flagutils.Map(test.name, test.defaultValue, "map usage")
			if test.value != "" || test.defaultValue == nil {
				err := flag.Set(test.name, test.value)
				if test.expectedError == "" {
					c.Assert(err, qt.Equals, nil)
				} else {
					c.Assert(err, qt.ErrorMatches, test.expectedError)
				}
			}
			c.Assert(*v, qt.DeepEquals, test.expectedValue)
		})
	}
}

func TestMapVar(t *testing.T) {
	for _, test := range mapTests {
		runIsolated(t, test.about, func(c *qt.C) {
			var v flagutils.StringMap
			flagutils.MapVar(&v, test.name, test.defaultValue, "map usage")
			if test.value != "" || test.defaultValue == nil {
				err := flag.Set(test.name, test.value)
				if test.expectedError == "" {
					c.Assert(err, qt.Equals, nil)
				} else {
					c.Assert(err, qt.ErrorMatches, test.expectedError)
				}
			}
			c.Assert(v, qt.DeepEquals, test.expectedValue)
		})
	}
}

func TestStringMapSet(t *testing.T) {
	for _, test := range mapTests {
		runIsolated(t, test.about, func(c *qt.C) {
			if test.defaultValue != nil {
				return
			}
			var v flagutils.StringMap
			err := v.Set(test.value)
			if test.expectedError == "" {
				c.Assert(err, qt.Equals, nil)
			} else {
				c.Assert(err, qt.ErrorMatches, test.expectedError)
			}
			c.Assert(v, qt.DeepEquals, test.expectedValue)
		})
	}
}

func TestStringMapString(t *testing.T) {
	for _, test := range mapTests {
		runIsolated(t, test.about, func(c *qt.C) {
			if test.defaultValue != nil {
				return
			}
			var v flagutils.StringMap
			v.Set(test.value)
			c.Assert(v.String(), qt.Equals, test.expectedStringValue)
		})
	}
}

// runIsolated runs the given test function without clobbering global flags.
func runIsolated(t *testing.T, name string, f func(c *qt.C)) {
	restore := resetForTesting()
	defer restore()
	qt.New(t).Run(name, f)
}

// resetForTesting creates a new flag set for the global command line and
// returns a function that restores the original global command line.
func resetForTesting() (restore func()) {
	original := flag.CommandLine
	flag.CommandLine = flag.NewFlagSet("", flag.ContinueOnError)
	return func() {
		flag.CommandLine = original
	}
}
