package settings

import "github.com/janstoon/toolbox/kareless"

const (
	lookupServer = "lookup.server"
)

func LookupServer(ss *kareless.Settings) string {
	return ss.GetString(lookupServer)
}

func init() {
	Default[lookupServer] = "8.8.8.8"
}
