// Licensed under the MIT license, see LICENCE file for details.

package flagutils

import (
	"encoding/json"
	"flag"
	"fmt"
	"strings"
)

// Slice defines a string slice flag with specified name, default value, and
// usage string. The return value is the address of a StringSlice variable that
// stores the value of the flag.
func Slice(name string, value []string, usage string) *StringSlice {
	var s StringSlice
	SliceVar(&s, name, value, usage)
	return &s
}

// SliceVar defines a string slice flag with specified name, default value, and
// usage string. The argument p points to a StringSlice variable in which to
// store the value of the flag.
func SliceVar(p *StringSlice, name string, value []string, usage string) {
	*p = value
	flag.Var(p, name, usage)
}

// StringSlice holds a slice of strings that can be provided via the command
// line as a comma separated list of values.
type StringSlice []string

// String implements flag.Value by returning the slice as a string.
func (s *StringSlice) String() string {
	return strings.Join(*s, ",")
}

// Set implements flag.Value by populating the slice from the given comma
// separated value.
func (s *StringSlice) Set(value string) error {
	*s = nil
	for _, v := range strings.Split(value, ",") {
		v = strings.TrimSpace(v)
		if v == "" {
			return fmt.Errorf("cannot include empty strings in the list")
		}
		*s = append(*s, v)
	}
	return nil
}

// Map defines a flag containing a map of strings with specified name, default
// value, and usage string. The return value is the address of a StringMap
// variable that stores the value of the flag.
func Map(name string, value map[string]interface{}, usage string) *StringMap {
	var s StringMap
	MapVar(&s, name, value, usage)
	return &s
}

// MapVar defines a flag containing a map of strings with specified name,
// default value, and usage string. The argument p points to a StringMap
// variable in which to store the value of the flag.
func MapVar(p *StringMap, name string, value map[string]interface{}, usage string) {
	*p = value
	flag.Var(p, name, usage)
}

// StringMap holds a map strings to empty interfaces that can be provided via
// the command line as a JSON encoded string.
type StringMap map[string]interface{}

// String implements flag.Value by returning the map as a string.
func (s *StringMap) String() string {
	b, err := json.Marshal(*s)
	if err != nil {
		// This should never happen.
		panic(err)
	}
	return string(b)
}

// Set implements flag.Value by unmarshaling the JSON encoded value into the
// string map. The JSON enclosing braces can be omitted.
func (s *StringMap) Set(value string) error {
	*s = nil
	value = strings.TrimSpace(value)
	if !strings.HasPrefix(value, "{") {
		value = "{" + value + "}"
	}
	if err := json.Unmarshal([]byte(value), s); err != nil {
		return fmt.Errorf("cannot unmarshal JSON: %v", err)
	}
	return nil
}
