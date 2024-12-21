package groper

import "github.com/janstoon/toolbox/kareless"

type Application struct {
	domains domainRepository
}

func NewApp(ss *kareless.Settings, ib *kareless.InstrumentBank) *Application {
	return &Application{
		domains: kareless.ResolveInstrumentByType[domainRepository](ib, "repo/dnslv/domain"),
	}
}
