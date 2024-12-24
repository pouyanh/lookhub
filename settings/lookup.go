package settings

import (
	"time"

	"github.com/janstoon/toolbox/kareless"
)

const (
	lookupServer = "lookup.server"
	lookupTTL    = "lookup.ttl"
)

func LookupServer(ss *kareless.Settings) string {
	return ss.GetString(lookupServer)
}

func LookupTTL(ss *kareless.Settings) time.Duration {
	return ss.GetDuration(lookupTTL)
}

func init() {
	Default[lookupServer] = "8.8.8.8"
	Default[lookupServer] = "5m"
}
