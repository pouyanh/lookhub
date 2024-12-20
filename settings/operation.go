package settings

import (
	"strings"

	"github.com/janstoon/toolbox/kareless"
)

const mode = "mode"

type Mode string

const (
	Standard   Mode = "STANDARD"
	Privileged Mode = "PRIVILEGED"
	Testing    Mode = "TESTING"
)

func OperationMode(ss *kareless.Settings) Mode {
	return Mode(strings.ToUpper(ss.GetString(mode)))
}
