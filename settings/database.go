package settings

import (
	"strings"

	"github.com/janstoon/toolbox/kareless"
	"github.com/janstoon/toolbox/tricks"
)

const (
	databases        = "databases"
	databaseAdapter  = "adapter"
	databaseHost     = "host"
	databasePort     = "port"
	databaseName     = "name"
	databaseUsername = "username"
	databasePassword = "password"
)

type Database struct {
	Name     string
	Adapter  string
	Host     string
	Port     int
	DbName   string
	Username string
	Password string
}

func DatabaseByName(ss *kareless.Settings, name string) Database {
	prefix := strings.Join([]string{databases, name}, separator)

	return Database{
		Name:     name,
		Adapter:  ss.GetString(strings.Join([]string{prefix, databaseAdapter}, separator)),
		Host:     ss.GetString(strings.Join([]string{prefix, databaseHost}, separator)),
		Port:     ss.GetInt(strings.Join([]string{prefix, databasePort}, separator)),
		DbName:   ss.GetString(strings.Join([]string{prefix, databaseName}, separator)),
		Username: ss.GetString(strings.Join([]string{prefix, databaseUsername}, separator)),
		Password: ss.GetString(strings.Join([]string{prefix, databasePassword}, separator)),
	}
}

func Databases(ss *kareless.Settings) []Database {
	return tricks.Map(ss.Children(databases), func(src string) Database {
		return DatabaseByName(ss, src)
	})
}

func init() {
	Default[databases] = map[string]any{}
}
