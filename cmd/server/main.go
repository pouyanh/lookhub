package main

import (
	"context"
	"fmt"

	"github.com/janstoon/toolbox/kareless"
	"github.com/janstoon/toolbox/kareless/std"
	"github.com/janstoon/toolbox/tricks"

	"gitlab.snapp.ir/pouyanh/sdig/config"
	"gitlab.snapp.ir/pouyanh/sdig/settings"
)

func main() {
	k := kareless.Compile().
		Feed(std.LocalEarlyLoadedSettingSource("conf", "/etc/sdig")).
		Feed(settings.Default).
		AfterStart(showWelcome)
	if err := k.Run(context.Background()); err != nil {
		panic(err)
	}
}

func showWelcome(
	ctx context.Context, ss *kareless.Settings, ib *kareless.InstrumentBank, apps []kareless.Application,
) error {
	fmt.Printf("Operation Mode: %s\n", settings.OperationMode(ss))
	fmt.Printf("Active apps: %v\n", tricks.Map(apps, func(src kareless.Application) string {
		return fmt.Sprintf("%T", src)
	}))
	config.PrintVars()

	return nil
}
