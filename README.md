# flagutils

This Go package provides helpers for parsing command line flags to string
slices and maps.

For instance, consider a program declaring the following flags:
```go
config := flagutils.Map("config", nil, "the program configuration as a JSON encoded string")
things := flagutils.Slice("things", nil, "a comma separated list of things")
```
At this point it is possible to call the program passing a JSON encoded value
and a list, for instance:
```
myprogram -config '{"key": "value", "nested": {"key2": ["value2"]}}' -things a,list,of,things
```
The two `config` and `things` flags are parsed into a *StringMap*
(*map[string]interface{}*) and a *StringSlice* (*[]string*) respectively.

See the [go documentation](https://godoc.org/github.com/frankban/flagutils) for
this library.
