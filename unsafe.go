package chrono

import (
	_ "unsafe" // for go:linkname
)

//go:linkname monotime runtime.nanotime
func monotime() int64

//go:linkname walltime runtime.walltime
func walltime() (secs int64, nsec int32)

//go:linkname zoneSources time.zoneSources
var zoneSources []string

//go:linkname embeddedTzData tzdata.zipdata
var embeddedTzData string

//go:linkname initLocal time.initLocal
func initLocal()
