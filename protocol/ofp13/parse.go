package ofp13

import (
	"errors"

	"github.com/maufl/openflow/protocol/util"
)

func Parse(b []byte) (message util.Message, err error) {
	switch b[1] {
	default:
		err = errors.New("An unknown v1.3 packet type was received. Parse function will discard data.")
	}
	return
}
