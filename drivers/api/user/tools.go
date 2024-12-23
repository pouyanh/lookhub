//go:build tools
// +build tools

package user

import (
	_ "github.com/go-swagger/go-swagger"
	_ "github.com/go-swagger/go-swagger/cmd/swagger/commands"
	_ "github.com/go-swagger/go-swagger/codescan"
	_ "github.com/go-swagger/go-swagger/generator"
)
