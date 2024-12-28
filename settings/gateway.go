package settings

import (
	"strings"

	"github.com/janstoon/toolbox/kareless"
	"github.com/janstoon/toolbox/tricks"
)

const (
	gateways       = "gateways"
	gatewayAddress = "address"
	gatewayProto   = "proto"

	apis          = "apis"
	apiCorsOrigin = "cors.origin"
	apiGateways   = "gateways"
)

type Gateway struct {
	Name    string
	Address string
	Proto   string
}

func GatewayByName(ss *kareless.Settings, name string) Gateway {
	prefix := strings.Join([]string{gateways, name}, separator)

	return Gateway{
		Name:    name,
		Address: ss.GetString(strings.Join([]string{prefix, gatewayAddress}, separator)),
		Proto:   tricks.Coalesce(ss.GetString(strings.Join([]string{prefix, gatewayProto}, separator)), "tcp"),
	}
}

func Gateways(ss *kareless.Settings) []Gateway {
	return tricks.Map(ss.Children(gateways), func(src string) Gateway {
		return GatewayByName(ss, src)
	})
}

type API struct {
	Name     string
	Cors     Cors
	Gateways []Gateway
}

type Cors struct {
	OriginWhitelist []string
}

func APIByName(ss *kareless.Settings, name string) API {
	prefix := strings.Join([]string{apis, name}, separator)

	return API{
		Name: name,
		Cors: Cors{
			OriginWhitelist: ss.GetStringSlice(strings.Join([]string{prefix, apiCorsOrigin}, separator)),
		},
		Gateways: tricks.Map(ss.GetStringSlice(strings.Join([]string{prefix, apiGateways}, separator)),
			func(src string) Gateway {
				return GatewayByName(ss, src)
			}),
	}
}

func APIs(ss *kareless.Settings) []API {
	return tricks.Map(ss.Children(apis), func(src string) API {
		return APIByName(ss, src)
	})
}

func init() {
	Default[gateways] = map[string]any{
		"user": map[string]string{
			gatewayAddress: ":30080",
		},
		"metrics": map[string]string{
			gatewayAddress: ":30089",
		},
	}

	Default[apis] = map[string]any{
		"user": map[string]any{
			apiCorsOrigin: "localhost *",
			apiGateways:   []string{"user"},
		},
		"metrics": map[string]any{
			apiCorsOrigin: "localhost *",
			apiGateways:   []string{"metrics"},
		},
	}
}
