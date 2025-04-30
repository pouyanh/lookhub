package main

import (
	"context"
	"fmt"

	"github.com/janstoon/toolbox/kareless"
	"github.com/janstoon/toolbox/kareless/std"
	"github.com/janstoon/toolbox/tricks"

	"github.com/pouyanh/lookhub/adapters"
	"github.com/pouyanh/lookhub/config"
	"github.com/pouyanh/lookhub/drivers"
	"github.com/pouyanh/lookhub/settings"
)

func main() {
	k := kareless.Compile(config.BundlesEssentials()...).
		Feed(std.LocalEarlyLoadedSettingSource("conf", "/etc/lookhub")).
		Feed(settings.Default).
		Equip(adapters.Adapters...).
		Connect(drivers.Drivers...).
		AfterStart(showWelcome)
	if err := k.Run(context.Background()); err != nil {
		panic(err)
	}
}

func showWelcome(
	_ context.Context, ss *kareless.Settings, _ *kareless.InstrumentBank, apps []kareless.Application,
) error {
	fmt.Printf("Operation Mode: %s\n", settings.OperationMode(ss))
	fmt.Printf("Active apps: %v\n", tricks.Map(apps, func(src kareless.Application) string {
		return fmt.Sprintf("%T", src)
	}))
	config.PrintVars()

	return nil
}
