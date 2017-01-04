package sensor

import (
	"github.com/hsson/open-home-controller/hardware"
)

type sensor struct {
	hardware.Module
	Values []string `json:"values"`
}
