package arg

import (
	"flag"
)

var Visible bool

func init() {
	flag.BoolVar(&Visible, "v", false, "Specify wheter it is necessary to output logs. Default: false.")
	flag.Parse()
}
