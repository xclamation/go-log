package arg

import (
	"flag"
	"github/xclamation/go-log/loglevel"
)

var v, vv, vvv, vvvv, vvvvv, vvvvvv bool
var initLevel uint8

func init() {
	initFlags()
}

func initFlags() {
	flag.BoolVar(&v, "v", false, "Set logging level to LEVEL 1.")
	flag.BoolVar(&vv, "vv", false, "Set logging level to LEVEL 2.")
	flag.BoolVar(&vvv, "vvv", false, "Set logging level to LEVEL 3.")
	flag.BoolVar(&vvvv, "vvvv", false, "Set logging level to LEVEL 4.")
	flag.BoolVar(&vvvvv, "vvvvv", false, "Set logging level to LEVEL 5.")
	flag.BoolVar(&vvvvvv, "vvvvvv", false, "Set logging level to LEVEL 6.")
	flag.Parse()
	SetInitLevel()
}

func SetInitLevel() {
	switch {
	case vvvvvv:
		initLevel = loglevel.LEVEL_6
	case vvvvv:
		initLevel = loglevel.LEVEL_5
	case vvvv:
		initLevel = loglevel.LEVEL_4
	case vvv:
		initLevel = loglevel.LEVEL_3
	case vv:
		initLevel = loglevel.LEVEL_2
	case v:
		initLevel = loglevel.LEVEL_1
	default:
		initLevel = loglevel.LEVEL_0
	}
}

func GetInitLevel() uint8 {
	return initLevel
}
