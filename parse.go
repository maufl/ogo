package openflow

import (
	"github.com/maufl/openflow/openflow10"
	"github.com/maufl/openflow/openflow13"
	"github.com/maufl/openflow/openflowxx"
)

func Parse(b []byte) (message openflowxx.Message, err error) {
	switch b[0] {
	case 1:
		message, err = openflow10.Parse(b)
	case 4:
		message, err = openflow13.Parse(b)
	}
	return
}
